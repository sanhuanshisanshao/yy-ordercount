package user

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
	"yy-ordercount/auto-buy/baseinfo"
	"yy-ordercount/auto-buy/client"
	"yy-ordercount/util"
)

var paraFormat = "{\"msgtype\": \"text\",\"text\": {\"content\": \"%s\"}, \"at\": {\"atMobiles\": [\"%s\"], \"isAtAll\": false}}"

type user struct {
	Url       string
	Phone     string
	Cookie    string
	IsDeleted bool
}

type Users struct {
	sync.Mutex
	Users []user
}

var UniqueUsers *Users //唯一用户实例

func NewUsers(cookie []string, url, phone string) {
	UniqueUsers = &Users{
		Users: make([]user, 0),
	}

	if len(cookie) > 0 {
		for _, v := range cookie {
			UniqueUsers.Users = append(UniqueUsers.Users, user{
				Cookie:    v,
				IsDeleted: false,
				Url:       url,
				Phone:     phone,
			})
		}
	}

	go UniqueUsers.AutoBuy()
}

func (u *Users) Add(cookie string) {
	u.Lock()
	defer u.Unlock()
	for _, v := range u.Users {
		if v.Cookie == cookie {
			v.IsDeleted = false
			return
		}
	}
	u.Users = append(u.Users, user{Cookie: cookie, IsDeleted: false})
}

func (u *Users) Delete(cookie string) {
	u.Lock()
	defer u.Unlock()
	for i := 0; i < len(u.Users); i++ {
		if u.Users[i].Cookie == cookie {
			logrus.Infof("delete cookie success")
			u.Users[i].IsDeleted = true
		}
	}
}

func (u *Users) AutoBuy() {
	for {
		t := 0
		for _, v := range u.Users {
			if v.IsDeleted == false {
				t++
				logrus.Infof("start to Auto buy .... ")
				if s := v.autoBuy(); s == "over" {
					u.Delete(v.Cookie)
				}
			}
		}
		if t == 0 {
			logrus.Warnf("no account to buy then stop server ...")
			os.Exit(0)
		}

		<-time.After(2 * time.Minute)
	}
}

func (u *user) autoBuy() (s string) {
	URL := "http://www.uuplush.com/user/buyorder"
	para := struct {
		OrderNum int    `json:"ordernum"`
		Gcid     int    `json:"gcid"`
		Gpid     int    `json:"gpid"`
		FieldNum string `json:"fieldnum"`
		Price    int    `json:"buyprice"`
	}{}

	if u.IsDeleted {
		logrus.Warnf("cookie %v is deleted", u.Cookie)
		return
	}
	account, err := u.getAccount() //获取账户余额
	if err != nil {
		logrus.Errorf("auto buy get account failed：%v", err)
		return
	}
	if (int(account) / 100) < 2 { //余额低于200不能下单
		logrus.Warnf("auto buy get account ¥%.2f less than ¥200", account)
		return
	}

	gcID, gpID, area, field, err := u.getOrder(account) //查询最近可下单期号
	if err != nil {
		logrus.Errorf("auto buy try to buy failed：%v", err)
		return
	}

	para.OrderNum = 49040000000000000 + int(time.Now().Unix()*1000) + rand.Intn(999-100) + 100
	para.Gcid = gcID
	para.Gpid = gpID
	para.FieldNum = field
	para.Price = (int(account) / 100) * 100
	//para.Price = 100

	ref := fmt.Sprintf("http://www.uuplush.com/buyorder?gcid=%d&gpid=%d&fieldnum=%s", gcID, gpID, field)
	b, _ := json.Marshal(&para)
	resp, err := client.HttpPost(URL, string(b), u.Cookie, ref) //下单
	if err != nil {
		logrus.Errorf("auto buy order failed：%v", err)
		return
	}

	if strings.Contains(string(resp), "每日限20单") {
		u.SendDDMsg(fmt.Sprintf("下单 ¥%v 失败: %v", para.Price, "每日限20单"), u.Phone)
		s = "over"
		return
	}

	logrus.Infof("buy ¥%v response:%v", para.Price, string(resp))

	msg := make(map[string]string)
	err = json.Unmarshal(resp, &msg)
	if err != nil {
		logrus.Errorf("json unmarshal error：%v", err)
		return
	}

	err = u.SendDDMsg(fmt.Sprintf("下单 %v %v ¥%v: %v", area, field, para.Price, msg["message"]), u.Phone)
	if err != nil {
		logrus.Errorf("send DD msg error：%v", err)
		return
	}
	return
}

//getAccount 获取账户余额
func (u *user) getAccount() (float64, error) {
	url := "http://www.uuplush.com/user/refyue"
	referer := "http://www.uuplush.com/v2/index?gcid=11"
	resp, err := client.HttpPost(url, "", u.Cookie, referer)
	if err != nil {
		logrus.Errorf("get account error: %v", err)
		return 0, err
	}
	f, err := strconv.ParseFloat(string(resp[1:len(resp)-1]), 32)
	if err != nil {
		logrus.Errorf("get account parse float error: %v", err)
		return 0, err
	}
	return f, nil
}

func (u *user) getOrder(price float64) (gcid int, gpid int, area string, fieldNum string, err error) {
	url := "http://www.uuplush.com/user/fieldlist"
	var p = struct {
		Gcid int `json:"gcid"`
		Gpid int `json:"gpid"`
	}{}
	candidates := []map[string]interface{}{}

	for _, val := range baseinfo.UniqueAreaIds {
		candicate := make(map[string]interface{})
		p.Gcid = val.Gcid
		p.Gpid = val.Gpid

		b, _ := json.Marshal(&p)
		resp, err := client.HttpPost(url, string(b), u.Cookie, "")
		if err != nil {
			return gcid, gpid, area, fieldNum, err
		}
		if len(resp) < 10 {
			logrus.Errorf("http post %v error", string(resp))
			return gcid, gpid, area, fieldNum, fmt.Errorf("http post %v error", string(resp))
		}
		m := util.ConvertResponse(resp)

		i := 0
		for _, v := range m {
			s := strings.Replace(v["syprice"].(string), "\"", "", -1) //去除 ""
			f, err := strconv.ParseFloat(s, 32)
			if err != nil {
				logrus.Errorf("get order parse float error: %v", err)
				return gcid, gpid, area, fieldNum, err
			}
			if f <= price {
				i++
				if i == 4 { //每个地区取前三个有效期号，无剩余交易量则取下一个地区
					logrus.Warnf("gcid %v order is full", val.Gcid)
					break
				}
				continue
			}
			//gcid = val.Gcid
			//gpid = val.Gpid
			//fieldNum = strings.Replace(v["fieldnum"].(string), "\"", "", -1)

			candicate["gcid"] = val.Gcid
			candicate["gpid"] = val.Gpid
			candicate["fieldNum"] = strings.Replace(v["fieldnum"].(string), "\"", "", -1)
			t, err := time.Parse("2006-01-02 15:04", v["kjtime"].(string))
			if err != nil {
				logrus.Errorf("time parse %v error: %v", v["kjtime"].(string), err)
				return 0, 0, "", "", fmt.Errorf("parse time error: %v", err)
			}
			candicate["kjTime"] = t
			candicate["area"] = val.Area
			candidates = append(candidates, candicate)
			break
		}
	}

	min := make(map[string]interface{})
	for i := 0; i < len(candidates); i++ {
		if i == 0 {
			min = candidates[0]
		} else if candidates[i]["kjTime"].(time.Time).Unix() < min["kjTime"].(time.Time).Unix() {
			min = candidates[i]
		}
	}

	if len(candidates) > 0 {
		return min["gcid"].(int), min["gpid"].(int), min["area"].(string), min["fieldNum"].(string), nil
	}
	return 0, 0, "", "", fmt.Errorf("get no condicate")
}

func (u *user) SendDDMsg(content, phone string) error {
	para := fmt.Sprintf(paraFormat, content, phone)
	_, err := client.HttpPost(u.Url, para, "", "")
	if err != nil {
		return err
	}
	return nil
}

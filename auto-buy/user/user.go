package user

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
	"yy-ordercount/auto-buy/baseinfo"
	"yy-ordercount/auto-buy/client"
	"yy-ordercount/util"
)

type user struct {
	Cookie    string
	IsDeleted bool
}

type Users struct {
	sync.Mutex
	Users []user
}

var UniqueUsers *Users //唯一用户实例

func NewUsers(cookie string) {
	UniqueUsers = &Users{
		Users: make([]user, 0),
	}
	if len(cookie) > 0 {
		UniqueUsers.Users = append(UniqueUsers.Users, user{
			Cookie:    cookie,
			IsDeleted: false,
		})
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
	for _, v := range u.Users {
		if v.Cookie == cookie {
			v.IsDeleted = true
		}
	}
}

func (u *Users) AutoBuy() {
	for {
		logrus.Infof("start to Auto buy .... at %v", time.Now())
		for _, v := range u.Users {
			v.autoBuy()
		}
		<-time.After(5 * time.Minute)
	}
}

func (u *user) autoBuy() {
	URL := "http://www.uuplush.com/user/buyorder"
	para := struct {
		OrderNum int    `json:"ordernum"`
		Gcid     int    `json:"gcid"`
		Gpid     int    `json:"gpid"`
		FieldNum string `json:"fieldnum"`
		Price    int    `json:"buyprice"`
	}{}

	if u.IsDeleted {
		logrus.Warnf("cookie %v is deleted")
		return
	}
	account, err := u.getAccount() //获取账户余额
	if err != nil {
		logrus.Errorf("auto buy get account failed：%v", err)
		return
	}
	if (int(account) / 100) < 1 { //余额低于100不能下单
		logrus.Warnf("auto buy get account %.2f less than 100", account)
		return
	}

	gcID, gpID, field, err := u.getOrder(account) //查询最近可下单期号
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

	logrus.Infof("buy response: %v at %v", string(resp), time.Now())
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

func (u *user) getOrder(price float64) (gcid int, gpid int, fieldNum string, err error) {
	url := "http://www.uuplush.com/user/fieldlist"
	var p = struct {
		Gcid int `json:"gcid"`
		Gpid int `json:"gpid"`
	}{}

LABEL:
	for _, val := range baseinfo.UniqueAreaIds {
		p.Gcid = val.Gcid
		p.Gpid = val.Gpid

		b, _ := json.Marshal(&p)
		resp, err := client.HttpPost(url, string(b), u.Cookie, "")
		if err != nil {
			return gcid, gpid, fieldNum, err
		}
		if len(resp) < 10 {
			logrus.Errorf("http post %v error", string(resp))
			return gcid, gpid, fieldNum, fmt.Errorf("http post %v error", string(resp))
		}
		m := util.ConvertResponse(resp)

		i := 0
		for _, v := range m {
			s := strings.Replace(v["syprice"].(string), "\"", "", -1) //去除 ""
			f, err := strconv.ParseFloat(s, 32)
			if err != nil {
				logrus.Errorf("get order parse float error: %v", err)
				return gcid, gpid, fieldNum, err
			}
			if f <= price+20000 { //20000并发容错
				i++
				if i == 4 { //每个地区取前三个有效期号，无剩余交易量则取下一个地区
					logrus.Warnf("gcid %v order is full", val.Gcid)
					goto LABEL
				}
				continue
			}
			gcid = val.Gcid
			gpid = val.Gpid
			fieldNum = strings.Replace(v["fieldnum"].(string), "\"", "", -1)
			return gcid, gpid, fieldNum, err
		}
	}
	return
}

package server

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
	"sync"
	"time"
	"yy-ordercount/client"
	"yy-ordercount/util"
)

const (
	MAX_GCID    = 8
	REDIS_ADDR  = "0.0.0.0:6379"
	KEY_FORMATE = "%v_%v_%v_%v"
	URL         = "http://www.uuplush.com/user/fieldlist"
	COOKIE      = "safedog-flow-item=635BE5DFF566A6414BC4F418238FBA41; userid=4904; username=4904; ASP.NET_SessionId=avhcnzfz1510lv00izjp0apy"
)

type para struct {
	Gcid int `json:"gcid"`
	Gpid int `json:"gpid"`
}

type Server struct {
	sync.Mutex
	Cookie    string
	StartTime time.Time
}

func NewServer(cookie string) *Server {
	return &Server{
		Cookie:    cookie,
		StartTime: time.Now(),
	}
}

func (s *Server) Ping(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("OK"))
}

func (s *Server) SetCookie(resp http.ResponseWriter, req *http.Request) {
	//todo:
}

func (s *Server) GetServerInfo(resp http.ResponseWriter, req *http.Request) {
	b, _ := json.Marshal(s)
	resp.Write(b)
}

func (s *Server) CountOrderInOneDayByGcID(id string) {
	maxOrder := 90
	gcids := []string{"1", "2", "3", "5", "8"}
	hashKey := "hash_5_" + util.GetToday()

	for _, gcid := range gcids {
		for i := 1; i <= maxOrder; i++ {
			key := "5_1_" + gcid + "_" + util.GetToday() + "-" + strconv.Itoa(i)
			if i < 10 {
				key = "5_1_" + gcid + "_" + util.GetToday() + "-0" + strconv.Itoa(i)
			}

			result, err := client.RedisClient.HGet(key, "end_time").Result()
			if err != nil {
				continue
			}
			if len(result) == 0 {
				continue
			}

			t, err := time.Parse("2006-01-02 15:04", result)
			if err != nil {
				log.Errorf("time parse %v error: %v", result, err)
				break
			}

			if t.Before(time.Now().Add(8 * time.Hour).UTC()) {
				remain, _ := client.RedisClient.HGet(key, "remain").Result()
				total, _ := client.RedisClient.HGet(key, "total").Result()
				num, _ := client.RedisClient.HGet(key, "num").Result()

				count := util.Subtraction(total, remain)
				if count > 0 {
					_, err = client.RedisClient.HSet(hashKey, fmt.Sprintf("%v_%v", gcid, num), count).Result()
					if err != nil {
						log.Errorf("hset %v error: %v", hashKey, err)
					}
					log.Infof("remain:%v  total:%v  num:%v  count:%v  end_time:%v", remain, total, num, count, t)
				}
			}
		}
	}
}

func (s *Server) GetTodayOrderCount(resp http.ResponseWriter, req *http.Request) {
	count, _ := s.getTodayOrderCount()
	log.Infof("GetTodayOrderCount total: %v", count)
	resp.Write([]byte(strconv.FormatInt(count, 10)))
}

func (s *Server) getTodayOrderCount() (int64, error) {
	key := fmt.Sprintf("%v_%v_%v", "hash", "5", util.GetToday())
	result, err := client.RedisClient.HGetAll(key).Result()
	if err != nil {
		log.Errorf("getTodayOrderCount hgetall %v error %v", key, err)
		return 0, err
	}
	var totalCount int64
	for _, val := range result {
		data, _ := strconv.ParseInt(val, 10, 0)
		totalCount += data
	}
	return totalCount, nil
}

func (s *Server) GetYYOrderInfo() {
	//每分钟定时执行，余额误差保持在 1 min 之内
	for {
		log.Infof("start to get gcid data ...")
		go func() {
			if time.Now().Hour() < 8 {
				return
			}
			for i := 1; i <= MAX_GCID; i++ {
				para := para{}
				para.Gcid = i
				para.Gpid = 1
				b, _ := json.Marshal(&para)

				resp, err := client.HttpPost(URL, string(b), s.Cookie)
				if err != nil {
					log.Errorf("http get YY error: %v", err)
				}

				if len(resp) <= 10 {
					//gcid =[4,6,7]
					continue
				}

				m := util.ConvertResponse(resp)

				for j := 0; j < len(m)-1; j++ {
					v := m[j]
					key := fmt.Sprintf(KEY_FORMATE, "5", v["gpid"], i, v["fieldnum"])
					client.RedisClient.HSet(key, "remain", v["syprice"])
					client.RedisClient.HSet(key, "total", m[len(m)-2]["syprice"])
					client.RedisClient.HSet(key, "end_time", v["kjtime"])
					client.RedisClient.HSet(key, "num", v["fieldnum"])

					client.RedisClient.Expire(key, 7*24*time.Hour) //设置过期时间
					log.Infof("store %v to redis success", key)
				}

				go s.CountOrderInOneDayByGcID(strconv.Itoa(i))
			}
		}()

		<-time.After(50 * time.Second)
	}
}

func Run() {
	err := client.NewRedis(REDIS_ADDR, "")
	if err != nil {
		log.Fatalf("connect to redis server failed: %v", err)
	}
	log.Infof("connect to redis success")

	server := NewServer(COOKIE)
	go server.GetYYOrderInfo()

	http.HandleFunc("/ping", server.Ping)
	http.HandleFunc("/info", server.GetServerInfo)
	http.HandleFunc("/setcookie", server.SetCookie)
	http.HandleFunc("/count", server.GetTodayOrderCount)

	log.Fatal(http.ListenAndServe("0.0.0.0:9988", nil))
}

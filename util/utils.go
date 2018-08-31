package util

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

//转化YY返回值为[]map[string]interface{}
//map[gpid:1 fieldnum:180830-40 syprice:-15100.00 kjtime:2018-08-30 15:39]
func ConvertResponse(resp []byte) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	resp = resp[1 : len(resp)-1] //去除双引号 ""
	respStr := string(resp)
	respStr = strings.Replace(respStr, `\`, "", -1) //去除转译符 \
	err := json.Unmarshal([]byte(respStr), &result)
	if err != nil {
		log.Error("ConvertResponse error: %v", err)
	}
	return result
}

//GetToday 获取当前时间，格式为：180929
func GetToday() string {
	s := time.Now().Format("2006/01/02")
	s = strings.Replace(s, "/", "", -1)
	return string([]rune(s)[2:])
}

//Subtraction return a - b
func Subtraction(a, b string) float64 {
	i, _ := strconv.ParseFloat(a, 0)
	j, _ := strconv.ParseFloat(b, 0)
	return i - j
}

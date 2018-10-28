package client

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpPost(url, param, cookie string) ([]byte, error) {
	return httpPost(url, param, cookie)
}

func httpPost(urlStr string, reqBody string, cookie string) (respBytes []byte, err error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", cookie)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	req.Header.Add("Referer", "https://uu.98765u.cn/buyorder?gcid=12&gpid=31&fieldnum=180910-74") //v2

	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

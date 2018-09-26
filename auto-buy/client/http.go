package client

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpPost(url, param, cookie, referer string) ([]byte, error) {
	return httpPost(url, param, cookie, referer)
}

func httpPost(urlStr string, reqBody string, cookie, referer string) (respBytes []byte, err error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	if len(cookie) > 0 {
		req.Header.Add("Cookie", cookie)

	}
	if len(referer) > 0 {
		req.Header.Add("Referer", referer)
	}

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

func HttpGet(url, cookie string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Cookie", cookie)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	r, _ := ioutil.ReadAll(resp.Body)
	return string(r), nil
}

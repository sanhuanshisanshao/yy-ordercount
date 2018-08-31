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

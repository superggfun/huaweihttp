package myProxy

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
)

func (m *APIG) myPost() (string, map[string]string, int, error) {
	url := m.getUrl()
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	req, err := http.NewRequest("POST", url, strings.NewReader(m.Body))
	if err != nil {
		return "", nil, 0, err
	}

	m.delHeader() // 删除华为云的请求头
	for key, value := range m.Headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", nil, 0, err
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", nil, 0, err
	}

	text, err := doBody(&bodyText, url)

	if err != nil {
		return "", nil, 0, err
	}

	h := make(map[string]string, 50)
	for key, value := range resp.Header {
		h[key] = value[0]
	}

	return text, h, resp.StatusCode, nil

}

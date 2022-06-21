package myProxy

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func doBody(bodyText *[]byte, url string) (string, error) {
	if strings.Contains(string(*bodyText), `charset=gbk`) {
		b, _ := gbkToUtf8(bodyText)
		bodyText = &b
	}

	reg := regexp.MustCompile(`"/(.*?).css"`)
	reg2 := regexp.MustCompile(`"/(.*?).js"`)
	reg3 := regexp.MustCompile(`href="(.*?)"`)
	reg4 := regexp.MustCompile(`src="(.*?)"`)
	reg5 := regexp.MustCompile(`(http|ftp|https)://(.*?)/`)

	url2 := reg5.FindAllStringSubmatch(url, -1)[0][0]
	text := string(*bodyText)
	if strings.Contains(text, `<head>`) {
		arr := strings.Split(text, `<head>`)
		text = strings.Join([]string{arr[0], `<head><meta name="referrer" content="no-referrer"/>`, strings.Join(arr[1:], "<head>")}, "")
	}

	for _, v := range reg.FindAllStringSubmatch(text, -1) {
		if strings.Contains(v[0], `//`) {
			continue
		}
		text = strings.Replace(text, v[0], strings.Join([]string{`"`, url2, `/`, v[1], `.css"`}, ""), -1)

	}

	for _, v := range reg2.FindAllStringSubmatch(text, -1) {
		if strings.Contains(v[0], `//`) {
			continue
		}
		text = strings.Replace(text, v[0], strings.Join([]string{`"`, url2, `/`, v[1], `.js"`}, ""), -1)
	}

	for _, v := range reg3.FindAllStringSubmatch(text, -1) {
		if strings.Contains(v[0], `//`) {
			continue
		}
		text = strings.Replace(text, v[0], strings.Join([]string{`href="`, url2, `/`, v[1]}, ""), -1)
	}

	for _, v := range reg4.FindAllStringSubmatch(text, -1) {
		if strings.Contains(v[0], `//`) {
			continue
		}
		text = strings.Replace(text, v[0], strings.Join([]string{`src="`, url2, `/`, v[1]}, ""), -1)

	}
	return text, nil
}

func gbkToUtf8(str *[]byte) (b []byte, err error) { // GBK 转 utf8编码
	b, err = ioutil.ReadAll(transform.NewReader(bytes.NewReader(*str), simplifiedchinese.GBK.NewDecoder()))
	if err != nil {
		return
	}
	return
}

package request_test

import (
	"github.com/gopkg-dev/go-common/ext"
	"github.com/gopkg-dev/go-common/request"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"testing"


)

func TestGet(t *testing.T) {
	url := "http://httpbin.org/get"
	resp, err := request.Get(url, nil)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(resp, "httpbin.org") {
		t.Error()
	}
}

func TestGetByte(t *testing.T) {
	url := "https://www.iqiyipic.com/common/fix/128-128-logo.png"
	resp, err := request.GetByte(url, nil)
	if err != nil {
		t.Error(err)
	}
	if len(resp) != 5638 {
		t.Error()
	}
}

func TestContentType(t *testing.T) {
	url := "https://www.iqiyipic.com/common/fix/128-128-logo.png"
	contentType, err := request.ContentType(url, nil)
	if err != nil {
		t.Error(err)
	}
	if contentType != "image/png" {
		t.Error()
	}
}

func TestHeaders(t *testing.T) {
	url := "http://log.mmstat.com/eg.js"
	headers, err := request.Headers(url, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(headers)
	if len(headers) == 0 {
		t.Error()
	}
}

func TestSize(t *testing.T) {
	url := "https://www.iqiyipic.com/common/fix/128-128-logo.png"
	size, err := request.Size(url, nil)
	if err != nil {
		t.Error(err)
	}
	if size != 5638 {
		t.Error()
	}
}

func TestRequest(t *testing.T) {

	postUrl := "https://httpbin.org/post"

	jar, _ := cookiejar.New(nil)

	ext.SetCookie(jar, postUrl, "test", "123456")

	config := request.Config{
		Headers: map[string]string{
			"User-Agent": request.RandomUserAgent(),
		},
		Proxy:    nil,
		Jar:      jar,
		Redirect: false,
	}

	postBody := "test=test&haha=11111111111"
	res, err := request.Request(http.MethodPost, postUrl, strings.NewReader(postBody), &config)
	if err != nil {
		t.Error(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil && err != io.EOF {
		panic(err)
	}

	if !strings.Contains(string(body), postBody) {
		t.Error()
	}

	rawCookies, err := ext.JarToRawCookies(jar, postUrl)
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(rawCookies, "123456") {
		t.Error()
	}

	test, err := ext.GetCookie(jar, postUrl, "test")
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(test, "123456") {
		t.Error()
	}

}

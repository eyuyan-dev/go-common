package request_test

import (
	"strings"
	"testing"

	"github.com/eyuyan-dev/go-common/request"
)

func TestNewHttpProxy(t *testing.T) {

	getUrl := "http://httpbin.org/get"

	config := request.Config{
		Proxy: request.NewHttpProxy("http://127.0.0.1:1080"),
	}

	resp, _ := request.Get(getUrl, &config)

	if !strings.Contains(resp, "httpbin.org") {
		t.Error("http代理可能失效了")
	}
}

func TestNewSocks5Proxy(t *testing.T) {
	getUrl := "http://httpbin.org/get"

	config := request.Config{
		Proxy: request.NewSocks5Proxy("127.0.0.1:1080"),
	}

	resp, _ := request.Get(getUrl, &config)

	if !strings.Contains(resp, "httpbin.org") {
		t.Error()
	}
}

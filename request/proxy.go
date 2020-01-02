package request

import (
	"net"
	"net/http"
	netURL "net/url"
	"time"

	"golang.org/x/net/proxy"
)

type HttpProxy http.Transport

func NewHttpProxy(httpProxy string) *HttpProxy {
	var _proxy, err = netURL.Parse(httpProxy)
	if err != nil {
		return nil
	}
	HttpProxy := &HttpProxy{
		Proxy: http.ProxyURL(_proxy),
	}
	return HttpProxy
}

func NewSocks5Proxy(socks5Proxy string) *HttpProxy {
	dialer, err := proxy.SOCKS5(
		"tcp",
		socks5Proxy,
		nil,
		&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		},
	)
	if err != nil {
		return nil
	}
	HttpProxy := &HttpProxy{
		Dial: dialer.Dial,
	}
	return HttpProxy
}

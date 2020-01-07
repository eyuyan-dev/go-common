package request

import (
	"compress/flate"
	"compress/gzip"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"
	"time"
)

var (

	// FakeHeaders fake http headers
	FakeHeaders = map[string]string{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.81 Safari/537.36",
	}

	// RetryTimes how many times to retry when the download failed
	RetryTimes int
)

type Config struct {
	Headers  map[string]string
	Proxy    *HttpProxy
	Jar      *cookiejar.Jar
	Redirect bool
}

// Request base request
func Request(method, url string, body io.Reader, config *Config) (*http.Response, error) {

	transport := &http.Transport{}

	if config != nil && config.Proxy != nil {
		transport = (*http.Transport)(config.Proxy)
	}

	transport.DisableCompression = true
	transport.TLSHandshakeTimeout = 10 * time.Second
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	transport.DisableKeepAlives = false

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	//set default http request headers
	for k, v := range FakeHeaders {
		req.Header.Set(k, v)
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   5 * time.Minute,
	}

	if config != nil {

		for k, v := range config.Headers {
			req.Header.Set(k, v)
		}

		if config.Jar != nil {
			client.Jar = config.Jar
		}

		if config.Redirect {
			client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}
		}
	}

	var (
		res          *http.Response
		requestError error
	)

	for i := 0; ; i++ {
		res, requestError = client.Do(req)
		if requestError == nil && res.StatusCode < 400 {
			break
		} else if i+1 >= RetryTimes {
			var err error
			if requestError != nil {
				err = fmt.Errorf("request error: %v", requestError)
			} else {
				err = fmt.Errorf("%s request error: HTTP %d", url, res.StatusCode)
			}
			return nil, err
		}
		time.Sleep(1 * time.Second)
	}

	return res, nil
}

// Get get request
func Get(url string, config *Config) (string, error) {
	body, err := GetByte(url, config)
	return string(body), err
}

// GetByte get request
func GetByte(url string, config *Config) ([]byte, error) {
	res, err := Request(http.MethodGet, url, nil, config)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var reader io.ReadCloser
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(res.Body)
	case "deflate":
		reader = flate.NewReader(res.Body)
	default:
		reader = res.Body
	}
	defer reader.Close()

	body, err := ioutil.ReadAll(reader)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return body, nil
}

// Headers return the HTTP Headers of the url
func Headers(url string, config *Config) (http.Header, error) {
	res, err := Request(http.MethodGet, url, nil, config)
	if err != nil {
		return nil, err
	}
	return res.Header, nil
}

// Size get size of the url
func Size(url string, config *Config) (int64, error) {
	h, err := Headers(url, config)
	if err != nil {
		return 0, err
	}
	s := h.Get("Content-Length")
	if s == "" {
		return 0, errors.New("Content-Length is not present")
	}
	size, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return size, nil
}

// ContentType get Content-Type of the url
func ContentType(url string, config *Config) (string, error) {
	h, err := Headers(url, config)
	if err != nil {
		return "", err
	}
	s := h.Get("Content-Type")
	// handle Content-Type like this: "text/html; charset=utf-8"
	return strings.Split(s, ";")[0], nil
}

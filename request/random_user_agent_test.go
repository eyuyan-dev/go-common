package request_test

import (
	"strings"
	"testing"

	"github.com/gopkg-dev/go-common/request"
)

func TestRandomUserAgent(t *testing.T) {
	getUrl := "https://httpbin.org/get"
	userAgent := request.RandomUserAgent()
	config := request.Config{
		Headers: map[string]string{
			"User-Agent": userAgent,
		},
	}
	resp, _ := request.Get(getUrl, &config)
	if !strings.Contains(resp, userAgent) {
		t.Errorf("UserAgent Error : %s \n %s", resp, userAgent)
	}
}

package ext

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	netURL "net/url"
	"time"
)

// GetCookie
func GetCookie(jar *cookiejar.Jar, url, name string) (string, error) {
	uri, err := netURL.Parse(url)
	if err != nil {
		return "", err
	}
	var cookies []*http.Cookie
	cookies = jar.Cookies(uri)
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie.Value, nil
		}
	}
	return "", nil
}

// SetCookie
func SetCookie(jar *cookiejar.Jar, url, name string, value string) {
	uri, err := netURL.Parse(url)
	if err != nil {
		return
	}
	jar.SetCookies(uri, []*http.Cookie{
		{
			Name:     name,
			Value:    value,
			Path:     "/",
			Domain:   uri.Host,
			Expires:  time.Unix(time.Now().Unix(), 0).AddDate(0, 1, 0),
			Secure:   false,
			HttpOnly: false,
		},
	})
}

// RawCookiesToJar
func RawCookiesToJar(cookie, url string) (*cookiejar.Jar, error) {

	uri, err := netURL.Parse(url)
	if err != nil {
		return nil, err
	}

	result := MatchAll(cookie, `([^;=]+)=([^;]+);?\s*`)
	if len(result) == 0 {
		return nil, errors.New("cookies error")
	}

	var cookies []*http.Cookie
	for _, v := range result {
		if len(v) == 3 {
			cookies = append(cookies, &http.Cookie{
				Name:     v[1],
				Value:    v[2],
				Path:     "/",
				Domain:   uri.Host,
				Expires:  time.Unix(time.Now().Unix(), 0).AddDate(0, 1, 0),
				Secure:   false,
				HttpOnly: false,
			})
		}
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	jar.SetCookies(uri, cookies)

	return jar, nil
}

// JarToRawCookies
func JarToRawCookies(jar *cookiejar.Jar, url string) (string, error) {
	uri, err := netURL.Parse(url)
	if err != nil {
		return "", err
	}
	cookies := jar.Cookies(uri)
	rawCookies := ""
	for i, v := range cookies {
		rawCookies += v.String()
		if i != len(cookies)-1 {
			rawCookies += "; "
		}
	}
	return rawCookies, nil
}

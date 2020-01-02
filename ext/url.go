package ext

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"

	"github.com/eyuyan-dev/go-common/request"
)

func MapStringToUrlParams(val map[string]string) string {
	values := url.Values{}
	for k, v := range val {
		values.Add(k, v)
	}
	return values.Encode()
}

func StructToUrlParams(i interface{}) string {

	values := url.Values{}

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段
			value := fmt.Sprintf("%v", v.Field(i).Interface())
			if len(t.Field(i).Tag) == 0 {
				values.Add(t.Field(i).Name, value)
			} else if name := t.Field(i).Tag.Get("json"); name != "" {
				values.Add(name, value)
				continue
			} else {
				fmt.Printf("Ignore field %s %s = %v -tag:%s \n",
					t.Field(i).Name,
					t.Field(i).Type,
					v.Field(i).Interface(),
					t.Field(i).Tag)
			}
		}
	}
	return values.Encode()
}

// Domain get the domain of given URL
func Domain(url string) string {
	domainPattern := `([a-z0-9][-a-z0-9]{0,62})\.` +
		`(com\.cn|com\.hk|` +
		`cn|com|net|edu|gov|biz|org|info|pro|name|xxx|xyz|be|` +
		`me|top|cc|tv|tt)`
	domain := MatchOneOf(url, domainPattern)
	if domain != nil {
		return domain[1]
	}
	return "Universal"
}

// GetNameAndExt return the name and ext of the URL
// https://img9.bcyimg.com/drawer/15294/post/1799t/1f5a87801a0711e898b12b640777720f.jpg ->
// 1f5a87801a0711e898b12b640777720f, jpg
func GetNameAndExt(uri string) (string, string, error) {
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return "", "", err
	}
	s := strings.Split(u.Path, "/")
	filename := strings.Split(s[len(s)-1], ".")
	if len(filename) > 1 {
		return filename[0], filename[1], nil
	}
	// Image url like this
	// https://img9.bcyimg.com/drawer/15294/post/1799t/1f5a87801a0711e898b12b640777720f.jpg/w650
	// has no suffix
	contentType, err := request.ContentType(uri, nil)
	if err != nil {
		return "", "", err
	}
	return filename[0], strings.Split(contentType, "/")[1], nil
}

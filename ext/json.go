package ext

import (
	"bytes"
	"encoding/json"

	"github.com/tidwall/gjson"
)

// GetStringFromJson get the string value from json path
func GetStringFromJson(json, path string) string {
	return gjson.Get(json, path).String()
}

// EncodePrettyJson encode pretty json
func EncodePrettyJson(v interface{}) string {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", "\t")
	jsonEncoder.Encode(v)
	return bf.String()
}

package utils

import (
	"fmt"
	"testing"
)

type Test struct {
	URL    string
	Origin string
}

func TestJSONEncode(t *testing.T) {
	// JSONDecode()

}
func TestJSONDecode(t *testing.T) {
	var data = make(map[string]string)

	data["test"] = "123"
	response, err := HTTPPost("https://httpbin.org/post", data)
	if err != nil {
		t.Errorf("HTTPGet Error: %s", err)
	}

	jsonData := response.Body()

	var s interface{}

	// json.Unmarshal(jsonData, &s)

	d := JSONDecode(jsonData, s)
	// fmt.Println(11)
	// fmt.Println(d)
	m := d.(map[string]interface{})
	for k, v := range m {
		fmt.Println(k, v)
	}
}

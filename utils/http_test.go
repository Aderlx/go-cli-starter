package utils

import (
	"fmt"
	"testing"
)

func TestHTTPGet(t *testing.T) {
	_, err := HTTPGet("http://www.baidu.com")
	if err != nil {
		t.Errorf("HTTPGet Error: %s", err)
	}

}

func TestHTTPPost(t *testing.T) {

	var data = make(map[string]string)

	data["test"] = "123"

	response, err := HTTPPost("https://httpbin.org/post", data)

	fmt.Println(string(response.Body()))
	if err != nil {
		t.Errorf("HTTPGet Error: %s", err)
	}
}

package utils

import (
	"go-cli-starter/log"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

// HTTPGet 简单封装下 fasthttp Get 请求
func HTTPGet(url string) (*fasthttp.Response, error) {

	var request = &fasthttp.Request{}

	// 设置 url
	request.SetRequestURI(url)

	// 设置 请求方式
	request.Header.SetMethod(fasthttp.MethodGet)

	var response = &fasthttp.Response{}

	var client = &fasthttp.Client{}

	if err := client.Do(request, response); err != nil {
		log.Logger.Error(err.Error())
		return nil, err
	}
	log.Logger.Info("Request URL: " + url)
	return response, nil

}

// HTTPPost fasthttp Post 请求 application/json
func HTTPPost(url string, data map[string]string) (*fasthttp.Response, error) {

	var request = &fasthttp.Request{}

	// map 转化 成 json 字符串
	jsonData := JSONEncode(data)

	requestBody := []byte(jsonData)

	request.SetBody(requestBody)

	request.Header.SetContentType("application/json")

	// 设置 url
	request.SetRequestURI(url)

	// 设置 请求方式
	request.Header.SetMethod(fasthttp.MethodPost)

	var response = &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(request, response); err != nil {
		log.Logger.Error(err.Error(), zap.String("service", "http"), zap.String("httpMethod", "POST"))
		return nil, err
	}

	log.Logger.Info("Request URL: "+url, zap.String("service", "http"), zap.String("httpMethod", "POST"))
	return response, nil

}

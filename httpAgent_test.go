package common

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func BenchmarkHttpAgent(t *testing.B) {
	for i := 0; i < 20; i++ {
		request := New()
		request = request.SetHeader("AHost", "commonServer")
		request = request.Post("https://s-api.xyhj.io/v1/w/zh/user/loginByAccount")

		data := map[string]string{
			"orgId":    "99",
			"account":  "songliang1573629825",
			"password": "123456",
		}

		_, body, err := request.ContentType(TypeFormUrlencoded).SendForm(data).End()
		if err != nil {
			println(err.Error())
			return
		}
		println(string(body))
	}
}

func TestHttpAgent_Timeout(t *testing.T) {
	request := New()
	request = request.SetHeader("AHost", "commonServer").Timeout(time.Millisecond * 500)
	request = request.Post("https://s-api.xyhj.io/v1/w/zh/user/loginByAccount")

	data := map[string]string{
		"orgId":    "99",
		"account":  "songliang1573629825",
		"password": "123456",
	}

	_, body, err := request.ContentType(TypeFormUrlencoded).SendForm(data).End()
	if err != nil {
		println(err.Error())
		return
	}
	println(string(body))
}

func TestHttpAgent(t *testing.T) {
	// 通过向基础公共服务进行登录注册上传文件，验证httpAgent是否有效

	// 1. POST登录
	request := New()
	request = request.SetHeader("AHost", "commonServer")
	request = request.Post("https://s-api.xyhj.io/v1/w/zh/user/loginByAccount")

	data := map[string]string{
		"orgId":    "99",
		"account":  "songliang1573629825",
		"password": "123456",
	}

	_, body, err := request.ContentType(TypeFormUrlencoded).SendForm(data).End()
	if err != nil {
		println(err.Error())
		return
	}
	println(string(body))

	// GET查询用户信息
	_, body, err = request.Get("https://s-api.xyhj.io/v1/w/zh/user/getUserInfo").End()
	if err != nil {
		println(err.Error())
		return
	}
	println(string(body))

	// 上传文件
	file1, err := os.OpenFile("httpUtil.go", os.O_RDONLY, 0755)
	if err != nil {
		println(err.Error())
		return
	}
	defer file1.Close()
	file2, err := os.OpenFile("logger.go", os.O_RDONLY, 0755)
	if err != nil {
		println(err.Error())
		return
	}
	defer file2.Close()
	data1, _ := ioutil.ReadAll(file1)
	data2, _ := ioutil.ReadAll(file2)
	request = request.SendFile(File{FileName: "httpUtil.go", FieldName: "a", Data: data1})
	request = request.SendFile(File{FileName: "logger.go", FieldName: "b", Data: data2})
	_, body, err = request.Post("https://s-api.xyhj.io/v1/w/zh/octopus/uploadResource").ContentType(TypeMultipartFormData).End()
	if err != nil {
		println(err.Error())
		return
	}
	println(string(body))
}

func TestHttpAgent_CurlCommand(t *testing.T) {
	request := New()
	request = request.SetHeader("AHost", "commonServer").Timeout(time.Millisecond * 500)
	request = request.Post("https://s-api.xyhj.io/v1/w/zh/user/loginByAccount")

	data := map[string]string{
		"orgId":    "99",
		"account":  "songliang1573629825",
		"password": "123456",
	}

	cmd, err := request.ContentType(TypeFormUrlencoded).SendForm(data).CurlCommand()
	if err != nil {
		println(err.Error())
		return
	}
	println(cmd)
}

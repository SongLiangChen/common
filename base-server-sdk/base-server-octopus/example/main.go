package main

import (
	"fmt"
	"github.com/becent/golang-common/base-server-sdk"
	"github.com/becent/golang-common/base-server-sdk/base-server-octopus"
	"time"
)

func main() {
	base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		AppId:           "10002",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			//OctopusServerHost: "http://127.0.0.1:5051",
			OctopusServerHost: "127.0.0.1:15051",
		},
		GRpcOnly: true,
	})
	defer base_server_sdk.ReleaseBaseServerSdk()

	now := time.Now()
	defer func(now time.Time) {
		println(time.Since(now).String())
	}(now)

	// email
	//err := base_server_octopus.SendEmailCode(5, base_server_octopus.BusinessLogin, "xxx@qq.com", "zh")
	//ret, err := base_server_octopus.VerifyEmailCode(5,base_server_octopus.BusinessLogin, "xxx@qq.com", "1235")
	//ret, err := base_server_octopus.CheckLastEmailVerifyResult(5,base_server_octopus.BusinessLogin,"email")
	//if err != nil {
	//	println(err.String())
	//}

	// sim
	//err := base_server_octopus.SendSimCode(5,base_server_octopus.BusinessLogin, "130xxxx1234", "zh")
	//ret, err := base_server_octopus.VerifySimCode(5, base_server_octopus.BusinessLogin, "130xxxx1234", "54321")
	//res, err := base_server_octopus.CheckLastSimVerifyResult(5,base_server_octopus.BusinessLogin, "130xxxx1234")

	// gt
	//res, err := base_server_octopus.InitGT(5, base_server_octopus.BusinessLogin,"130xxxx1234", "284770")
	//base_server_octopus.VerifyGt(1, base_server_octopus.BusinessLogin, "130xxxx1234", "ip", "challenge", "validate", "seccode")

	// ga
	//res, err := base_server_octopus.GenerateGa(5, base_server_octopus.BusinessLogin, "130xxxx1234")
	//ret, err := base_server_octopus.VerifyGa(5, base_server_octopus.BusinessLogin, "130xxxx1234", "secret", "code")

	// idCard
	//res, err := base_server_octopus.AuthRealName(5, "张三", "010203201909201234")

	// resource
	formFile := make(map[string]string)
	formFile["file1"] = "test.log"
	formFile["file2"] = "test.2.log"
	res, err := base_server_octopus.Upload(1, 1, formFile)

	if err != nil {
		println(err.String())
	}
	fmt.Printf("res:%+v \n", res)
}

package main

import (
	"fmt"
	"github.com/becent/golang-common/base-server-sdk"
	"github.com/becent/golang-common/base-server-sdk/base-server-statistic"
	"time"
)

func main() {
	base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		AppId:           "10002",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			StatisticServerHost: "127.0.0.1:18081",
		},
		GRpcOnly: true,
	})
	defer base_server_sdk.ReleaseBaseServerSdk()

	now := time.Now()
	defer func(now time.Time) {
		println(time.Since(now).String())
	}(now)

	response, err := base_server_statistic.ListStatisticItems(0, "", "", "", 1548731008, 1574996608, 20, 0)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("获取统计数据成功：[%v]\n", response)
	}

}

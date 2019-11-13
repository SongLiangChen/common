package main

import (
	"fmt"
	"github.com/becent/golang-common/base-server-sdk"
	"github.com/becent/golang-common/base-server-sdk/base-server-interact"
	"time"
)

func main() {
	base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		AppId:           "10002",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			InteractServerHost: "http://127.0.0.1:8082",
		},
		GRpcOnly: false,
	})
	defer base_server_sdk.ReleaseBaseServerSdk()

	// 添加一条评论
	r, err := base_server_interact.AddInteract(&base_server_interact.InteractRecord{
		OrgId:      3,
		UserId:     10000,
		NickName:   "song",
		Type:       1,
		MediaId:    10,
		MediaType:  "news",
		MediaOwner: 10000,
		Title:      "hello",
		Content:    "world",
		Imgs:       []string{"http.123.png"},
	}, 0)
	if err != nil {
		println(err.String())
		return
	}
	fmt.Printf("%v\n", *r)

	// 添加一条点赞
	r, err = base_server_interact.AddInteract(&base_server_interact.InteractRecord{
		OrgId:      3,
		UserId:     10000,
		NickName:   "song",
		Type:       3,
		MediaId:    10,
		MediaType:  "news",
		MediaOwner: 10000,
	}, 1)
	if err != nil {
		println(err.String())
		return
	}
	fmt.Printf("%v\n", *r)

	// 删除一条点赞
	err = base_server_interact.DelInteract(3, 10000, r.InteractId)
	if err != nil {
		println(err.String())
		return
	}

	// 列表
	list, err := base_server_interact.InteractList(3, 0, 0, 0, 0, 0, 0, 0, 0)
	if err != nil {
		println(err.String())
		return
	}
	for _, r := range list {
		fmt.Printf("%v\n", *r)
	}

	// 详情
	detail, err := base_server_interact.InteractDetail(3, 273274329515426180)
	if err != nil {
		println(err.String())
		return
	}
	fmt.Printf("%v\n", *detail)

	// 计数
	counting, err := base_server_interact.InteractCounting(3, 10000, "news", []int64{10}, []int{1, 3})
	if err != nil {
		println(err.String())
		return
	}
	for ttype, c := range counting {
		println(ttype)
		fmt.Printf("%v\n", c)
	}

	// 得分
	score, err := base_server_interact.InteractScore(3, 10, "news")
	fmt.Printf("%0.2f\n", score)
}

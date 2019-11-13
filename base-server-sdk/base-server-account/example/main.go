package main

import (
	"base-server-account/model"
	"encoding/json"
	"github.com/becent/golang-common/base-server-sdk"
	base_server_account "github.com/becent/golang-common/base-server-sdk/base-server-account"
	"time"
)

func main() {
	base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		AppId:           "10008",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			AccountServerHost: "http://127.0.0.1:5050",
		},
		GRpcOnly: false,
	})
	defer base_server_sdk.ReleaseBaseServerSdk()

	now := time.Now()
	defer func(now time.Time) {
		println(time.Since(now).String())
	}(now)

	//// 创建账户
	//account, err := base_server_account.CreateAccount(8, 100000, []string{"CC", "USD"})
	//
	//if err != nil {
	//	println(err.String())
	//} else {
	//	fmt.Printf("创建成功: {%v}\n", account)
	//}

	//// 账户信息
	//accounts, err := base_server_account.AccountInfo(8, 100000, "CC")
	//if err != nil {
	//	println(err.String())
	//} else {
	//	fmt.Printf("账户信息: {%v}\n", accounts)
	//}
	//
	// 账户信息
	//accounts, err := base_server_account.AccountsInfo(8, "100000", "")
	//if err != nil {
	//	println(err.String())
	//} else {
	//	fmt.Printf("账户信息: {%v}\n", accounts)
	//}

	//accountList, err := base_server_account.AccountList(8, 0, "", 0, 0, 0, 0, 10)
	//if err != nil {
	//	println(err.String())
	//} else {
	//	fmt.Printf("账户列表: {%v}\n", accountList)
	//}

	//// 账户状态更新
	//err = base_server_account.UpdateStatus(8, 9, base_server_account.ACCOUNT_STATUS_FREEZE)
	//if err != nil {
	//	println(err.String())
	//} else {
	//	println("状态更新成功")
	//}
	//
	// 金额操作
	callback := &model.TaskCallBack{
		CallBackUrl: "baidu.com",
		Data:        nil,
	}
	callbackStr, _ := json.Marshal(callback)
	err := base_server_account.OperateAmount(8, 1, base_server_account.OP_TYPE_AVAIL_ADD, 1, 0, "100", "custom json string", "custom json string", string(callbackStr))
	if err != nil {
		println(err.String())
	} else {
		println("操作成功")
	}
	//
	//// 账户日志列表
	//logList, err := base_server_account.AccountLogList(8, 1, base_server_account.OP_TYPE_AVAIL_SUB, 1, "", 1568776001, 1569228191, 1, 10)
	//if err != nil {
	//	println(err.String())
	//} else {
	//	fmt.Printf("账户列表： {%v}", logList)
	//}
	//
	////账户日志总计
	//sumAmount, err := base_server_account.SumLog(8, 100000, base_server_account.OP_TYPE_AVAIL_ADD, 1, "CC", 0, 0)
	//if err != nil {
	//	println(err.String())
	//} else {
	//	fmt.Printf("总和： {%v}", sumAmount)
	//}
	//
	//// 批量账户操作
	//var taskDetails []*base_server_account.TaskDetail
	//taskDetails = append(taskDetails, &base_server_account.TaskDetail{
	//	OpType:        1,
	//	BsType:        1,
	//	AccountId:     1,
	//	UserId:        100000,
	//	Currency:      "CC",
	//	AllowNegative: 0,
	//	Amount:        "100",
	//	Detail:        "sss",
	//	Ext:           "ddd",
	//}, &base_server_account.TaskDetail{
	//	OpType:        1,
	//	BsType:        1,
	//	AccountId:     1,
	//	UserId:        100001,
	//	Currency:      "CC",
	//	AllowNegative: 0,
	//	Amount:        "200",
	//	Detail:        "SS",
	//	Ext:           "DD",
	//})
	//callback := &base_server_account.TaskCallBack{
	//	CallBackUrl: "xyz.com/callback",
	//	Data: map[string]string{
	//		"field1": "val1",
	//		"field2": "val2",
	//	},
	//}
	//
	//err = base_server_account.BatchOperateAmount(8, 1, taskDetails, callback)
	//if err != nil {
	//	println(err.String())
	//} else {
	//	fmt.Printf("操作成功")
	//}

	//账户账转
	//err = base_server_account.Transfer(8, 1, 2, "1000")
	//if err != nil {
	//	println(err.String())
	//} else {
	//	println("操作成功")
	//}
}

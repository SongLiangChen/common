package main

import (
	"fmt"
	"github.com/becent/golang-common/base-server-sdk"
	"github.com/becent/golang-common/base-server-sdk/base-server-pay"
	"strconv"
	"time"
)

func main() {
	base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		AppId:           "10002",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			AiPayServerHost: "https://t-openapi.aipaybox.com" +
				"",
		},
		GRpcOnly: false,
	})
	defer base_server_sdk.ReleaseBaseServerSdk()

	now := time.Now()
	defer func(now time.Time) {
		println(time.Since(now).String())
	}(now)

	mchId := "1"
	signKey := "cea67d27e295e352c09e4b580503ae42"

	response, err := base_server_pay.SelectPayMethods(mchId, base_server_pay.CURRENCY_CC, "", signKey)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("查询支付方式：[%v]\n", response)
	}

	response1, err1 := base_server_pay.SelectPayChannels(mchId, "wechat_qr_code_cny",
		"100", base_server_pay.CURRENCY_CC, base_server_pay.CURRENCY_CC, signKey)
	if err1 != nil {
		println(err1.String())
	} else {
		fmt.Printf("查询支付通道：[%v]\n", response1)
	}

	response2, err2 := base_server_pay.SubmitPayOrder(mchId, "wechat_qr_code_cny", "4", base_server_pay.RandString(10),
		"100", base_server_pay.CURRENCY_CC, base_server_pay.CURRENCY_CC, "rrr", "ee", "http://www.baidu.com", signKey)
	if err2 != nil {
		println(err2.String())
	} else {
		fmt.Printf("支付下单：[%v]\n", response2)
	}

	response3, err3 := base_server_pay.GenerateUnionPayUrl(mchId, base_server_pay.CURRENCY_CC, "222", strconv.FormatInt(time.Now().Unix(), 10),
		"100", "http://www.baidu.com", "http://www.baidu.com", signKey)
	if err3 != nil {
		println(err3.String())
	} else {
		fmt.Printf("聚合支付链接：[%v]\n", response3)
	}

	response4, err4 := base_server_pay.SelectWithdrawMethods(mchId, base_server_pay.CURRENCY_CC, signKey)
	if err4 != nil {
		println(err4.String())
	} else {
		fmt.Printf("提现方式：[%v]\n", response4)
	}

	response5, err5 := base_server_pay.SelectWithdrawChannels(mchId, "wdw_bank_card_pay_cny", "1", base_server_pay.CURRENCY_CC, base_server_pay.CURRENCY_CC, signKey)
	if err5 != nil {
		println(err5.String())
	} else {
		fmt.Printf("提现通道：[%v]\n", response5)
	}

	response6, err6 := base_server_pay.SubmitWithdrawOrder(mchId, "wdw_bank_card_pay_cny", "2", base_server_pay.RandString(10),
		"10", base_server_pay.CURRENCY_CC, base_server_pay.CURRENCY_CC, "", "", "http://www.baidu.com", signKey, &base_server_pay.BankAccount{
			BankCity:       "深圳",
			BankProvince:   "广东",
			BankCardNo:     "44444444444444444444444444444444444444444444444444444444",
			BankName:       "银行名称",
			BankBranchName: "支行名称",
			BankUserName:   "姓名",
			BankUserPhone:  "",
		})
	if err6 != nil {
		println(err6.String())
	} else {
		fmt.Printf("提现订单：[%v]\n", response6)
	}

	response7, err7 := base_server_pay.QueryWithdrawOrder(mchId, "TYCGWSBNKY", "210011571217172904", signKey)
	if err7 != nil {
		println(err7.String())
	} else {
		fmt.Printf("代付查询：[%v]\n", response7)
	}
}

package base_server_octopus

import (
	"errors"
)

type BusinessId int

const (
	BusinessRegister        BusinessId = 1000 // 注册
	BusinessLogin           BusinessId = 1001 // 登录
	BusinessBindPhone       BusinessId = 1002 // 绑定手机
	BusinessUnBindPhone     BusinessId = 1003 // 解绑手机
	BusinessBindEmail       BusinessId = 1004 // 绑定邮箱
	BusinessUnBindEmail     BusinessId = 1005 // 解绑邮箱
	BusinessGetBackLoginPwd BusinessId = 1006 // 找回密码
	BusinessGetBackTransPwd BusinessId = 1007 // 找回支付密码
)

var businessIdMap = map[int]BusinessId{
	1000: BusinessRegister,
	1001: BusinessLogin,
	1002: BusinessBindPhone,
	1003: BusinessUnBindPhone,
	1004: BusinessBindEmail,
	1005: BusinessUnBindEmail,
	1006: BusinessGetBackLoginPwd,
	1007: BusinessGetBackTransPwd,
}

func GetBusinessId(id int) (BusinessId, error) {
	if _, ok := businessIdMap[id]; !ok {
		return 0, errors.New("invalid business type")
	}

	return businessIdMap[id], nil
}

package base_server_octopus

import (
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
)

// 发送邮件验证码
// --- businessId ---
//  1000 // 注册
//  1001 // 登录
//  1002 // 更新手机
//  1003 // 绑定手机
//  1004 // 更新邮箱
//  1005 // 绑定邮箱
//  1006 // 找回密码

func SendSimCode(orgId int, businessId BusinessId, countryCode, phone, lang string) *base_server_sdk.Error {
	if orgId == 0 || businessId == 0 || phone == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["businessId"] = strconv.Itoa(int(businessId))
	params["countryCode"] = countryCode
	params["phone"] = phone
	params["lang"] = lang

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.OctopusServerHost, "sim", "sendSimCode", params)
	if err != nil {
		return err
	}

	return nil
}

func SendSimMsg(orgId int, businessId BusinessId, countryCode, phone, lang, message string) *base_server_sdk.Error {
	if orgId == 0 || businessId == 0 || phone == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["businessId"] = strconv.Itoa(int(businessId))
	params["countryCode"] = countryCode
	params["phone"] = phone
	params["lang"] = lang
	params["message"] = message

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.OctopusServerHost, "sim", "sendSimMsg", params)
	if err != nil {
		return err
	}

	return nil
}

func VerifySimCode(orgId int, businessId BusinessId, countryCode, phone, code string) (bool, *base_server_sdk.Error) {
	if orgId == 0 || businessId == 0 || phone == "" {
		return false, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["businessId"] = strconv.Itoa(int(businessId))
	params["countryCode"] = countryCode
	params["phone"] = phone
	params["code"] = code

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.OctopusServerHost, "sim", "verifyCode", params)
	if err != nil {
		return false, err
	}

	return true, nil
}

func CheckLastSimVerifyResult(orgId int, businessId BusinessId, countryCode, phone string) (bool, *base_server_sdk.Error) {
	if orgId == 0 || businessId == 0 || phone == "" {
		return false, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["businessId"] = strconv.Itoa(int(businessId))
	params["countryCode"] = countryCode
	params["phone"] = phone

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.OctopusServerHost, "sim", "checkLastVerifyResult", params)
	if err != nil {
		return false, err
	}

	return true, nil
}

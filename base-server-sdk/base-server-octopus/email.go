package base_server_octopus

import (
	"encoding/json"
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
)

// 发送邮件验证码
// --- businessId ---
//  1000 注册
//  1001 登录
//  1002 更新手机
//  1003 绑定手机
//  1004 更新邮箱
//  1005 绑定邮箱
//  1006 找回密码
func SendEmailCode(orgId int, businessId BusinessId, email, lang string) *base_server_sdk.Error {
	if orgId == 0 || businessId == 0 || email == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["businessId"] = strconv.Itoa(int(businessId))
	params["email"] = email
	params["lang"] = lang

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.OctopusServerHost, "email", "sendEmailCode", params)
	if err != nil {
		return err
	}
	return nil
}

func SendEmailMsg(orgId int, businessId BusinessId, email, lang string, message map[string]interface{}) *base_server_sdk.Error {
	if orgId == 0 || businessId == 0 || email == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["businessId"] = strconv.Itoa(int(businessId))
	params["email"] = email
	params["lang"] = lang
	bytes, _ := json.Marshal(message)
	params["message"] = string(bytes)

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.OctopusServerHost, "email", "sendEmailMsg", params)
	if err != nil {
		return err
	}
	return nil
}

func VerifyEmailCode(orgId int, businessId BusinessId, email, code string) (bool, *base_server_sdk.Error) {
	if orgId == 0 || businessId == 0 || email == "" {
		return false, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["businessId"] = strconv.Itoa(int(businessId))
	params["email"] = email
	params["code"] = code

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.OctopusServerHost, "email", "verifyCode", params)
	if err != nil {
		return false, err
	}
	return true, nil
}

func CheckLastEmailVerifyResult(orgId int, businessId BusinessId, email string) (bool, *base_server_sdk.Error) {
	if orgId == 0 || businessId == 0 || email == "" {
		return false, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["businessId"] = strconv.Itoa(int(businessId))
	params["email"] = email

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.OctopusServerHost, "email", "checkLastVerifyResult", params)
	if err != nil {
		return false, err
	}
	return true, nil
}

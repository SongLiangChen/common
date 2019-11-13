package base_server_octopus

import (
	"encoding/json"
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
)

type InitCaptchaResponse struct {
	Success   bool   `json:"success"`
	CaptchaId string `json:"captchaId"`
	Image     string `json:"image"`
}

func InitCaptcha(orgId int, businessId BusinessId, length, width, height int) (*InitCaptchaResponse, *base_server_sdk.Error) {
	if orgId == 0 || businessId == 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["businessId"] = strconv.Itoa(int(businessId))
	params["length"] = strconv.Itoa(length)
	params["width"] = strconv.Itoa(width)
	params["height"] = strconv.Itoa(height)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.OctopusServerHost, "captcha", "initCaptcha", params)
	if err != nil {
		return nil, err
	}
	var resp InitCaptchaResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, base_server_sdk.ErrServiceBusy
	}
	return &resp, nil
}

func ReloadCaptcha(orgId int, businessId BusinessId, captchaId string, width, height int) (*InitCaptchaResponse, *base_server_sdk.Error) {
	if orgId == 0 || businessId == 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}
	if captchaId == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["businessId"] = strconv.Itoa(int(businessId))
	params["captchaId"] = captchaId
	params["width"] = strconv.Itoa(width)
	params["height"] = strconv.Itoa(height)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.OctopusServerHost, "captcha", "reloadCaptcha", params)
	if err != nil {
		return nil, err
	}
	var resp InitCaptchaResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, base_server_sdk.ErrServiceBusy
	}
	return &resp, nil
}

func VerifyCaptcha(orgId int, businessId BusinessId, captchaId string, digits string) (bool, *base_server_sdk.Error) {
	if orgId == 0 || businessId == 0 {
		return false, base_server_sdk.ErrInvalidParams
	}
	if captchaId == "" || digits == "" {
		return false, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["businessId"] = strconv.Itoa(int(businessId))
	params["captchaId"] = captchaId
	params["digits"] = digits

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.OctopusServerHost, "captcha", "verifyCaptcha", params)
	if err != nil {
		return false, err
	}

	return true, nil
}

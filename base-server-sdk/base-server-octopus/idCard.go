package base_server_octopus

import (
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
)

func AuthRealName(orgId int, name string, cardNo string) (bool, *base_server_sdk.Error) {
	if orgId == 0 || name == "" || cardNo == "" {
		return false, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["name"] = name
	params["cardNo"] = cardNo

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.OctopusServerHost, "idcard", "authRealName", params)
	if err != nil {
		return false, err
	}

	return true, nil
}

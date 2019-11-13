package base_server_goods

import (
	"encoding/json"
	common "github.com/becent/golang-common"
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
)

//添加商品类别
func AddCategory(orgId int, mchId int64, parentId int, name string, sortOrder int) (*Category, *base_server_sdk.Error) {

	if orgId <= 0 || mchId <= 0 || parentId < 0 || sortOrder <= 0 || name == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["mchId"] = strconv.FormatInt(mchId, 10)
	params["parentId"] = strconv.Itoa(parentId)
	params["sortOrder"] = strconv.Itoa(sortOrder)
	params["name"] = name

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.GoodsServerHost, "goods", "addCategory", params)
	if err != nil {
		return nil, err
	}
	c := &Category{}
	if err := json.Unmarshal(data, c); err != nil {
		common.ErrorLog("baseServerSdk_AddCategory", params, "unmarshal category fail"+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}
	return c, nil
}

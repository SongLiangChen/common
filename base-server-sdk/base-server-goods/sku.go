package base_server_goods

import (
	"base-server-goods/model"
	"encoding/json"
	common "github.com/becent/golang-common"
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
)

//添加属性
func AddAttribute(orgId int, mchId int64, categoryId int, attributeName string, sortOrder int) (*AttributeKey, *base_server_sdk.Error) {
	if orgId <= 0 || mchId <= 0 || categoryId <= 0 || sortOrder <= 0 || attributeName == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["mchId"] = strconv.FormatInt(mchId, 10)
	params["categoryId"] = strconv.Itoa(categoryId)
	params["sortOrder"] = strconv.Itoa(sortOrder)
	params["attributeName"] = attributeName

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.GoodsServerHost, "goods", "addAttribute", params)
	if err != nil {
		return nil, err
	}
	attrK := &AttributeKey{}
	if err := json.Unmarshal(data, attrK); err != nil {
		common.ErrorLog("baseServerSdk", params, "unmarshal AttributeKey fail"+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}
	return attrK, nil
}

//添加属性值
func AddAttributeValue(orgId int, mchId, attributeId int64, attributeValue string, sortOrder int) (*AttributeValue, *base_server_sdk.Error) {
	if orgId <= 0 || mchId <= 0 || attributeId <= 0 || sortOrder <= 0 || attributeValue == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["mchId"] = strconv.FormatInt(mchId, 10)
	params["attributeId"] = strconv.FormatInt(attributeId, 10)
	params["sortOrder"] = strconv.Itoa(sortOrder)
	params["attributeValue"] = attributeValue

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.GoodsServerHost, "goods", "addAttributeValue", params)
	if err != nil {
		return nil, err
	}
	attrV := &AttributeValue{}
	if err := json.Unmarshal(data, attrV); err != nil {
		common.ErrorLog("baseServerSdk", params, "unmarshal AttributeValue fail"+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}
	return attrV, nil
}

//添加sku
func AddSku(orgId int, mchId, productId int64, specs map[string]string, sortOrder int, price string, stock string) (*Sku, *base_server_sdk.Error) {
	if orgId <= 0 || mchId <= 0 || productId <= 0 || len(specs) == 0 || sortOrder <= 0 || price == "" || stock == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["mchId"] = strconv.FormatInt(mchId, 10)
	params["productId"] = strconv.FormatInt(productId, 10)
	params["sortOrder"] = strconv.Itoa(sortOrder)
	byteV, e := json.Marshal(specs)
	if e != nil {
		return nil, base_server_sdk.ErrInvalidParams
	}
	params["specs"] = string(byteV)
	params["price"] = price
	params["stock"] = stock

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.GoodsServerHost, "goods", "addSku", params)
	if err != nil {
		return nil, err
	}
	sku := &Sku{}
	if err := json.Unmarshal(data, sku); err != nil {
		common.ErrorLog("baseServerSdk", params, "unmarshal sku fail"+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}
	return sku, nil
}

// 操作库存
func OperateStock(orgId int, mchId, productId, skuId int64, qty string, opType model.StockOpType) *base_server_sdk.Error {
	if orgId <= 0 || productId <= 0 || skuId <= 0 || qty == "" || qty == "0" || opType <= 0 {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["mchId"] = strconv.FormatInt(mchId, 10)
	params["productId"] = strconv.FormatInt(productId, 10)
	params["skuId"] = strconv.FormatInt(skuId, 10)
	params["qty"] = qty
	params["opType"] = strconv.Itoa(int(opType))

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.GoodsServerHost, "goods", "operateStock", params)
	if err != nil {
		return err
	}
	return nil
}

// 批量操作库存
func BatchOperateStock(orgId int, batchData []*TaskBatchOperateStock, isAsync int) *base_server_sdk.Error {
	if orgId <= 0 || len(batchData) == 0 {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	data, err := json.Marshal(batchData)
	if err != nil {
		return base_server_sdk.ErrInvalidParams
	}
	params["batchData"] = string(data)
	params["isAsync"] = strconv.Itoa(isAsync)

	client := base_server_sdk.Instance
	if _, err := client.DoRequest(client.Hosts.GoodsServerHost, "goods", "batchOperateStock", params); err != nil {
		return err
	}
	return nil
}


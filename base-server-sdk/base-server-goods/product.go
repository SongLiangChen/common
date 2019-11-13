package base_server_goods

import (
	"encoding/json"
	common "github.com/becent/golang-common"
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
)

//添加商品
func AddProduct(orgId int, mchId int64, categoryId int, title, subTitle, detail, mainImg, subImg, video string,
	status ProductStatus, isTiming int, execTime int64, execStatus ProductStatus) (*Product, *base_server_sdk.Error) {
	if orgId <= 0 || mchId <= 0 || categoryId <= 0 || title == "" || subTitle == "" || detail == "" ||
		mainImg == "" || subImg == "" || video == "" || status <= 0 || (isTiming > 0 && (execTime <= 0 || execStatus <= 0)) {
		return nil, base_server_sdk.ErrInvalidParams
	}
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["mchId"] = strconv.FormatInt(mchId, 10)
	params["categoryId"] = strconv.Itoa(categoryId)
	params["title"] = title
	params["subTitle"] = subTitle
	params["detail"] = detail
	params["mainImg"] = mainImg
	params["subImg"] = subImg
	params["video"] = video
	params["title"] = title
	params["status"] = strconv.Itoa(int(status))
	params["isTiming"] = strconv.Itoa(isTiming)
	params["execTime"] = strconv.FormatInt(execTime, 10)
	params["execStatus"] = strconv.Itoa(int(execStatus))

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.GoodsServerHost, "goods", "addProduct", params)
	if err != nil {
		return nil, err
	}
	pro := &Product{}
	if err := json.Unmarshal(data, pro); err != nil {
		common.ErrorLog("baseServerSdk", params, "unmarshal product fail"+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}
	return pro, nil
}

// 编辑商品
func EditProduct(orgId int, mchId, productId int64, categoryId int, title, subTitle, detail, mainImg, subImg, video string,
	status ProductStatus, isTiming int, execTime int64, execStatus ProductStatus) (*Product, *base_server_sdk.Error) {
	if orgId <= 0 || mchId <= 0 || productId <= 0 || categoryId <= 0 || title == "" || subTitle == "" || detail == "" ||
		mainImg == "" || subImg == "" || video == "" || status <= 0 || (isTiming > 0 && (execTime <= 0 || execStatus <= 0)) {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["mchId"] = strconv.FormatInt(mchId, 10)
	params["categoryId"] = strconv.Itoa(categoryId)
	params["productId"] = strconv.FormatInt(productId, 10)
	params["title"] = title
	params["subTitle"] = subTitle
	params["detail"] = detail
	params["mainImg"] = mainImg
	params["subImg"] = subImg
	params["video"] = video
	params["title"] = title
	params["status"] = strconv.Itoa(int(status))
	params["isTiming"] = strconv.Itoa(isTiming)
	params["execTime"] = strconv.FormatInt(execTime, 10)
	params["execStatus"] = strconv.Itoa(int(execStatus))

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.GoodsServerHost, "goods", "editProduct", params)
	if err != nil {
		return nil, err
	}
	pro := &Product{}
	if err := json.Unmarshal(data, pro); err != nil {
		common.ErrorLog("baseServerSdk", params, "unmarshal product fail"+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}
	return pro, nil
}

// 商品详情
func GetProduct(orgId int, mchId, productId int64) (*ProductDetail, *base_server_sdk.Error) {
	if orgId <= 0 || mchId <= 0 || productId <= 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["mchId"] = strconv.FormatInt(mchId, 10)
	params["productId"] = strconv.FormatInt(productId, 10)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.GoodsServerHost, "goods", "getProduct", params)
	if err != nil {
		return nil, err
	}
	proDetail := &ProductDetail{}
	if err := json.Unmarshal(data, proDetail); err != nil {
		common.ErrorLog("baseServerSdk", params, "unmarshal productDetail fail"+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}
	return proDetail, nil
}

// 产品列表
func FindProduct(orgId int, mchId, productId int64, categoryId int, title, subTitle string, status, withSku int) []*ProductDetail {
	if orgId <= 0 || mchId <= 0 {
		return nil
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["mchId"] = strconv.FormatInt(mchId, 10)
	params["categoryId"] = strconv.Itoa(categoryId)
	params["productId"] = strconv.FormatInt(productId, 10)
	params["title"] = title
	params["subTitle"] = subTitle
	params["title"] = title
	params["status"] = strconv.Itoa(status)
	params["withSku"] = strconv.Itoa(withSku)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.GoodsServerHost, "goods", "findProduct", params)
	if err != nil {
		return nil
	}
	var proList []*ProductDetail
	if err := json.Unmarshal(data, &proList); err != nil {
		common.ErrorLog("baseServerSdk", params, "unmarshal productList fail"+string(data))
		return nil
	}
	return proList
}

//商品删除
func DelProduct(orgId int, mchId, productId int64) *base_server_sdk.Error {
	if orgId <= 0 || mchId <= 0 || productId <= 0 {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["mchId"] = strconv.FormatInt(mchId, 10)
	params["productId"] = strconv.FormatInt(productId, 10)

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.GoodsServerHost, "goods", "delProduct", params)
	if err != nil {
		return err
	}
	return nil
}

//商品上下架
func OnOffShelves(orgId int, mchId, productId int64, status ProductStatus, isTiming int, execTime int64) *base_server_sdk.Error {
	if orgId <= 0 || mchId <= 0 || productId <= 0 || status <= 0 || (isTiming > 0 && execTime <= 0) {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["mchId"] = strconv.FormatInt(mchId, 10)
	params["productId"] = strconv.FormatInt(productId, 10)
	params["status"] = strconv.Itoa(int(status))
	params["isTiming"] = strconv.Itoa(isTiming)
	params["execTime"] = strconv.FormatInt(execTime, 10)

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.GoodsServerHost, "goods", "onOffShelves", params)
	if err != nil {
		return err
	}
	return nil
}

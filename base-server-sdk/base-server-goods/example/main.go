package main

import (
	"encoding/json"
	"fmt"
	"github.com/becent/golang-common/base-server-sdk"
	base_server_goods "github.com/becent/golang-common/base-server-sdk/base-server-goods"
	"time"
)

func main() {
	base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		AppId:           "10008",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			GoodsServerHost: "http://127.0.0.1:5055",
		},
		GRpcOnly: false,
	})
	defer base_server_sdk.ReleaseBaseServerSdk()

	now := time.Now()
	defer func(now time.Time) {
		println(time.Since(now).String())
	}(now)

	//添加类别
	if category, err := base_server_goods.AddCategory(8, 100000, 0, "水果", 1); err != nil {
		println(err.String())
	} else {
		fmt.Printf("添加成功: {%v}\n", category)
	}

	//添加属性
	if attrK, err := base_server_goods.AddAttribute(8, 100000, 11, "新鲜度", 1); err != nil {
		println(err.String())
	} else {
		fmt.Printf("添加成功: {%v}", attrK)
	}

	//添加属性值
	if attrV, err := base_server_goods.AddAttributeValue(8, 100000, 10, "1kg", 1); err != nil {
		println(err.String())
	} else {
		fmt.Printf("添加成功: {%v}", attrV)
	}

	//添加商品
	if product, err := base_server_goods.AddProduct(8, 100000, 11, "新鲜水果", "买一箱送一箱", "水果详情", "/uri/main.jpg", "/uri/subImg1.jpg,/uri/subImg2.jpg", "/uri/video.mp4", base_server_goods.PRODUCT_STATUS_ON_SHELVES, 0, 0, 0); err != nil {
		println(err.String())
	} else {
		fmt.Printf("添加成功: {%v}", product)
	}

	//编辑商品
	if product, err := base_server_goods.EditProduct(8, 100000, 3, 11, "新鲜水果", "买一箱送两箱", "水果详情", "/uri/main.jpg", "/uri/subImg1.jpg,/uri/subImg2.jpg", "/uri/video.mp4", base_server_goods.PRODUCT_STATUS_ON_SHELVES, 0, 0, 0); err != nil {
		println(err.String())
	} else {
		fmt.Printf("编辑成功: {%v}", product)
	}

	//添加sku
	skuSpecs := map[string]string{
		"重量":  "1kg",
		"新鲜度": "100%",
	}
	if sku, err := base_server_goods.AddSku(8, 100000, 3, skuSpecs, 1, "100", "10"); err != nil {
		println(err.String())
	} else {
		fmt.Printf("添加成功: {%v}", sku)
	}

	//操作sku库存
	if err := base_server_goods.OperateStock(8, 100000, 3, 4, "100", 1); err != nil {
		println(err.String())
	} else {
		fmt.Print("操作成功")
	}

	//批量操作sku库存
	var batchData []*base_server_goods.TaskBatchOperateStock
	batchData = append(batchData, &base_server_goods.TaskBatchOperateStock{
		MchId:     100000,
		SkuId:     1,
		ProductId: 1,
		Qty:       "100",
		OpType:    base_server_goods.ADD_STOCK,
	})
	batchData = append(batchData, &base_server_goods.TaskBatchOperateStock{
		MchId:     100000,
		SkuId:     2,
		ProductId: 1,
		Qty:       "10",
		OpType:    base_server_goods.SUB_STOCK,
	})
	if err := base_server_goods.BatchOperateStock(8, batchData, 0); err != nil {
		println(err.String())
	} else {
		fmt.Print("操作成功")
	}

	//商品详情
	if product, err := base_server_goods.GetProduct(8, 100000, 3); err != nil {
		println(err.String())
	} else {
		byteV, _ := json.Marshal(product)
		fmt.Printf("%v", string(byteV))
	}

	//商品列表
	products := base_server_goods.FindProduct(8, 100000, 0, 0, "", "", 0, 1)
	byteV, _ := json.Marshal(products)
	fmt.Printf("%v", string(byteV))

	//删除商品
	if err := base_server_goods.DelProduct(8, 100000, 3); err != nil {
		println(err.String())
	} else {
		println("删除成功")
	}

	//上下架商品
	if err := base_server_goods.OnOffShelves(8, 100000, 3, 1, 0, 0); err != nil {
		println(err.String())
	} else {
		println("操作成功")
	}
}

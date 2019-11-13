# base-server-goods

商品基础服务

## 相关model
```go

type CategoryStatus int

const (
	//商品类别状态
	STATUS_ENABLE  CategoryStatus = 1 //启用
	STATUS_DISABLE CategoryStatus = 2 //禁用
)

type Category struct {
	CategoryId int            `json:"categoryId"` // 类别id
	ParentId   int            `json:"parentId"`   // 父类id
	Name       string         `json:"name"`       // 类别名称
	Status     CategoryStatus `json:"status"`     // 状态 1:可用 2:禁用
	SortOrder  int            `json:"sortOrder"`  // 排列次序
	CreateTime int64          `json:"createTime"` // 创建时间
	UpdateTime int64          `json:"updateTime"` // 更新时间
}

type AttributeKey struct {
	AttributeId   int    `json:"attributeId"`   // 属性id
	CategoryId    int    `json:"categoryId"`    // 商品类别id
	AttributeName string `json:"attributeName"` // 属性名称
	SortOrder     int    `json:"sortOrder"`     // 属性排列
	CreateTime    int64  `json:"createTime"`    // 创建时间
	UpdateTime    int64  `json:"updateTime"`    // 更新时间
}

type AttributeValue struct {
	ValueId        int64  `json:"valueId"`        // 属性值id
	AttributeId    int64  `json:"attributeId"`    // 属性id
	AttributeValue string `json:"attributeValue"` // 属性值
	SortOrder      int    `json:"sortOrder"`      // 排序次序
	CreateTime     int64  `json:"createTime"`     // 创建时间
	UpdateTime     int    `json:"updateTime"`     // 更新时间
}

type Product struct {
	ProductId  int64  `json:"productId"`  // 商品id
	OrgId      int    `json:"orgId"`      // 所属项目id
	MchId      int64  `json:"mchId"`      // 所属商家id
	CategoryId int    `json:"categoryId"` // 类别id
	Title      string `json:"title"`      // 商品名称
	SubTitle   string `json:"subTitle"`   // 商品子名称
	Detail     string `json:"detail"`     // 商品详情
	MainImg    string `json:"mainImg"`    // 主图
	SubImg     string `json:"subImg"`     // 副图集
	Video      string `json:"video"`      // 视频
	Status     int    `json:"status"`     // 商品状态 1:上架 2:下架
	CreateTime int64  `json:"createTime"` // 创建时间
	UpdateTime int64  `json:"updateTime"` // 更新时间
}

type ProductDetail struct {
	*Product
	SkuList []*Sku
}

type StockOpType int

const (
	ADD_STOCK        StockOpType = 1 //加库存
	SUB_STOCK        StockOpType = 2 //减库存
	FREEZE           StockOpType = 3 //冻结库存
	SUB_FREEZE_STOCK StockOpType = 4 //减冻结库存
	UN_FREEZE        StockOpType = 5 //解冻库存
)

type Sku struct {
	SkuId       int64  `json:"skuId"`       // skuId
	ProductId   int64  `json:"productId"`   // 商品id
	Specs       string `json:"specs"`       // 规格列表
	SortOrder   int    `json:"sortOrder"`   // 规格序列
	Price       string `json:"price"`       // 价格
	Stock       int    `json:"stock"`       // 库存
	FreezeStock int    `json:"freezeStock"` // 冻结库存
	CreateTime  int64  `json:"createTime"`  // 创建时间
	UpdateTime  int64  `json:"updateTime"`  // 更新时间
}

const (
	TASK_STATUS_SUCCESS TaskStatus = 1 //任务待处理状态
	TASK_STATUS_PROCESS TaskStatus = 2 //任务完成状态

	TASK_TYPE_GOODS_CHANGE_STATUS TaskType = 1 //上下架任务类型

	PRODUCT_STATUS_ON_SHELVES  ProductStatus = 1 //上架
	PRODUCT_STATUS_OFF_SHELVES ProductStatus = 2 //下架
)

type TaskStatus int
type TaskType int
type ProductStatus int

type Task struct {
	TaskId       int64      `json:"taskId"`
	TaskType     TaskType   `json:"taskType"`
	Status       TaskStatus `json:"status"`
	OrgId        int        `json:"orgId"`
	Detail       string     `json:"detail"`
	ParentTaskId int64      `json:"parentTaskId"`
	ExecTime     int64      `json:"execTime"`
	CreateTime   int64      `json:"createTime"`
	UpdateTime   int64      `json:"updateTime"`
}

type TaskGoodsChangeStatus struct {
	ProductId int64         `json:"productId"`
	Status    ProductStatus `json:"status"`
}

type TaskCallBack struct {
	CallBackUrl string            `json:"callBackUrl"`
	Data        map[string]string `json:"data"`
}

```

# 接口文档

### 添加商品类别

`post` `/goods/addCategory`

| 参数名 | 值 | 必填 | 说明 |
| :------ | :------ | :------ | :------ |
| orgId | int | 是 | 项目id |
| mchId | int64 | 是 | 商家id |
| parentId | int | 是 | 父类id |
| name | string | 是 | 类别名称 |

```go
返回值:
{
    "gateway-success": true,
    "gateway-orgId": "0",
    "gateway-userId": "0",
    "success": true,
    "payload": {
        "categoryId": 10,
        "parentId": 8,
        "name": "iphone11 pro",
        "status": 1,
        "sortOrder": 1,
        "createTime": 1572934082,
        "updateTime": 0
    }
}

异常错误:
1001 参数错误
2001 添加商品类别失败
```


### 添加商品属性

`post` `/goods/addAttribute`

| 参数名 | 值 | 必填 | 说明 |
| :------ | :------ | :------ | :------ |
| orgId | int | 是 | 项目id |
| mchId | int64 | 是 | 商家id |
| categoryId | int | 是 | 类别id |
| attributeName | string | 是 | 属性名称 |
| sortOrder | number | 是 | 排序 |

```go
返回值:
{
    "gateway-success": true,
    "gateway-orgId": "0",
    "gateway-userId": "0",
    "success": true,
    "payload": {
        "attributeId": 8,
        "categoryId": 8,
        "attributeName": "内存",
        "sortOrder": 1,
        "createTime": 1572936848,
        "updateTime": 0
    }
}

异常错误:
1001 参数错误
2002 添加商品属性失败
```

### 添加商品属性值

`post` `/goods/addAttributeValue`

| 参数名 | 值 | 必填 | 说明 |
| :------ | :------ | :------ | :------ |
| orgId | int | 是 | 项目id |
| mchId | int64 | 是 | 商家id |
| attributeId | int64 | 是 | 属性id |
| attributeValue | string | 是 | 属性值 |
| sortOrder | int | 是 | 排序 |

```go
返回值:
{
    "gateway-success": true,
    "gateway-orgId": "0",
    "gateway-userId": "0",
    "success": true,
    "payload": {
        "valueId": 6,
        "attributeId": 8,
        "attributeValue": "256G",
        "sortOrder": 3,
        "createTime": 1572937058,
        "updateTime": 0
    }
}

异常错误:
1001 参数错误
2003 添加商品属性值失败
```

### 添加商品

`post` `/goods/addProduct`

| 参数名 | 值 | 必填 | 说明 |
| :------ | :------ | :------ | :------ |
| orgId | int | 是 | 项目id |
| mchId | int64 | 是 | 商家id |
| mchId | int64 | 是 | 商家id |
| categoryId | int | 是 | 商品类别id |
| title | string | 是 | 标题 |
| subTitle | string | 是 | 副标题 |
| detail | string | 是 | 商品详情 |
| mainImg | string | 是 | 主图地址 |
| subImg | string | 是 | 副图地址,逗分 |
| video | string | 是 | 视频地址 |
| status | ProductStatus | 是 | 产品状态 |
| isTiming | int | 否 | 是否定时上下架 |
| execTime | int | 否 | 定时时间 |
| execStatus | int | 否 | 定时修改的状态 |


```go
type ProductStatus int
PRODUCT_STATUS_ON_SHELVES  ProductStatus = 1 //上架
PRODUCT_STATUS_OFF_SHELVES ProductStatus = 2 //下架

返回值:
{
    "gateway-success": true,
    "gateway-orgId": "0",
    "gateway-userId": "0",
    "success": true,
    "payload": {
        "productId": 2,
        "orgId": 8,
        "mchId": 100000,
        "categoryId": 8,
        "title": "iphone xs max 金色 64G",
        "subTitle": "双十一 双一送一",
        "detail": "{\"a\":1,\"b\",2}",
        "mainImg": "img/main.jpg",
        "subImg": "img/1.jpg,img/2.jpg,img/3.jpg",
        "video": "video.mp4",
        "status": 2,
        "createTime": 1573193615,
        "updateTime": 0
    }
}

异常错误:
1001 参数错误
2004 添加商品失败
4001 添加任务失败
```


### 编辑商品

`post` `/goods/editProduct`

| 参数名 | 值 | 必填 | 说明 |
| :------ | :------ | :------ | :------ |
| orgId | int | 是 | 项目id |
| mchId | int64 | 是 | 商家id |
| mchId | int64 | 是 | 商家id |
| productId | int64 | 是 | 商品id |
| categoryId | int | 是 | 商品类别id |
| title | string | 是 | 标题 |
| subTitle | string | 是 | 副标题 |
| detail | string | 是 | 商品详情 |
| mainImg | string | 是 | 主图地址 |
| subImg | string | 是 | 副图地址,逗分 |
| video | string | 是 | 视频地址 |
| status | ProductStatus | 是 | 产品状态 |
| isTiming | int | 否 | 是否定时上下架 |
| execTime | int | 否 | 定时时间 |
| execStatus | int | 否 | 定时修改的状态 |


```go
type ProductStatus int
PRODUCT_STATUS_ON_SHELVES  ProductStatus = 1 //上架
PRODUCT_STATUS_OFF_SHELVES ProductStatus = 2 //下架

返回值:
{
    "gateway-success": true,
    "gateway-orgId": "0",
    "gateway-userId": "0",
    "success": true,
    "payload": {
        "productId": 2,
        "orgId": 8,
        "mchId": 100000,
        "categoryId": 8,
        "title": "iphone xs max 金色 64G",
        "subTitle": "双十一 双一送一",
        "detail": "{\"a\":1,\"b\",2}",
        "mainImg": "img/main.jpg",
        "subImg": "img/1.jpg,img/2.jpg,img/3.jpg",
        "video": "video.mp4",
        "status": 2,
        "createTime": 1573193615,
        "updateTime": 0
    }
}

异常错误:
1001 参数错误
2005 商品不存在
2006 编辑商品失败
4001 添加任务失败
```

### 商品详情

`post` `/goods/getProduct`

| 参数名 | 值 | 必填 | 说明 |
| :------ | :------ | :------ | :------ |
| orgId | int | 是 | 项目id |
| mchId | int64 | 是 | 商家id |
| mchId | int64 | 是 | 商家id |
| productId | int64 | 是 | 产品id |

```go
返回值:
{
    "gateway-success": true,
    "gateway-orgId": "0",
    "gateway-userId": "0",
    "success": true,
    "payload": {
        "productId": 1,
        "orgId": 8,
        "mchId": 100000,
        "categoryId": 8,
        "title": "iphone xs max 金色 256G",
        "subTitle": "双十一 双一送一",
        "detail": "{\"a\":1,\"b\",2}",
        "mainImg": "img/main.jpg",
        "subImg": "img/1.jpg,img/2.jpg,img/3.jpg",
        "video": "video.mp4",
        "price": "10000",
        "stock": 100,
        "status": 1,
        "createTime": 1572940642,
        "updateTime": 0,
        "SkuList": [
            {
                "skuId": 1,
                "productId": 1,
                "specs": "{\"颜色\":\"黑色\", \"版本\":\"64G\"}",
                "sortOrder": 1,
                "price": "5000",
                "stock": 100,
                "createTime": 1572942384,
                "updateTime": 0
            },
            {
                "skuId": 2,
                "productId": 1,
                "specs": "{\"颜色\":\"白色\", \"版本\":\"64G\"}",
                "sortOrder": 1,
                "price": "5000",
                "stock": 100,
                "createTime": 1572942469,
                "updateTime": 0
            },
            {
                "skuId": 3,
                "productId": 1,
                "specs": "{\"颜色\":\"金色\", \"版本\":\"256G\"}",
                "sortOrder": 1,
                "price": "15000",
                "stock": 100,
                "createTime": 1572942482,
                "updateTime": 0
            }
        ]
    }
}

异常错误:
1001 参数错误
2005 商品不存在
```

### 商品列表

`post` `/goods/findProduct`

| 参数名 | 值 | 必填 | 说明 |
| :------ | :------ | :------ | :------ |
| orgId | int | 是 | 项目id |
| mchId | int64 | 是 | 商家id |
| mchId | int64 | 是 | 商家id |
| withSku | int | 否 | 是否带产品sku信息 1:带sku 0:默认不带 |

```go
返回值:
{
    "gateway-success": true,
    "gateway-orgId": "0",
    "gateway-userId": "0",
    "success": true,
    "payload": [
        {
            "productId": 1,
            "orgId": 8,
            "mchId": 100000,
            "categoryId": 8,
            "title": "iphone xs max 金色 256G",
            "subTitle": "双十一 双一送一",
            "detail": "{\"a\":1,\"b\",2}",
            "mainImg": "img/main.jpg",
            "subImg": "img/1.jpg,img/2.jpg,img/3.jpg",
            "video": "video.mp4",
            "price": "10000",
            "stock": 100,
            "status": 1,
            "createTime": 1572940642,
            "updateTime": 0,
            "SkuList": [
                {
                    "skuId": 1,
                    "productId": 1,
                    "specs": "{\"颜色\":\"黑色\", \"版本\":\"64G\"}",
                    "sortOrder": 1,
                    "price": "5000",
                    "stock": 100,
                    "createTime": 1572942384,
                    "updateTime": 0
                },
                {
                    "skuId": 2,
                    "productId": 1,
                    "specs": "{\"颜色\":\"白色\", \"版本\":\"64G\"}",
                    "sortOrder": 1,
                    "price": "5000",
                    "stock": 100,
                    "createTime": 1572942469,
                    "updateTime": 0
                },
                {
                    "skuId": 3,
                    "productId": 1,
                    "specs": "{\"颜色\":\"金色\", \"版本\":\"256G\"}",
                    "sortOrder": 1,
                    "price": "15000",
                    "stock": 100,
                    "createTime": 1572942482,
                    "updateTime": 0
                }
            ]
        }
    ]
}

异常错误:
1001 参数错误
```


### 删除商品

`post` `/goods/delProduct`

| 参数名 | 值 | 必填 | 说明 |
| :------ | :------ | :------ | :------ |
| orgId | int | 是 | 项目id |
| mchId | int64 | 是 | 商家id |
| productId | int64 | 是 | 产品id |

```go
返回值:
{
    "gateway-success": true,
    "gateway-orgId": "0",
    "gateway-userId": "0",
    "success": true,
    "payload": true
}

异常错误:
1001 参数错误
2005 商品不存在
2007 删除商品失败
```

### 商品上下架

`post` `/goods/onOffShelves`

| 参数名 | 值 | 必填 | 说明 |
| :------ | :------ | :------ | :------ |
| orgId | int | 是 | 项目id |
| mchId | int64 | 是 | 商家id |
| productId | int64 | 是 | 产品id |
| status | ProductStatus | 是 | 状态 |
| isTiming | int | 是 | 是否定时扫行 |
| execTime | int64 | 是 | 定时执行时间 |

```go
type ProductStatus int
PRODUCT_STATUS_ON_SHELVES  ProductStatus = 1 //上架
PRODUCT_STATUS_OFF_SHELVES ProductStatus = 2 //下架

返回值:
{
    "gateway-success": true,
    "gateway-orgId": "0",
    "gateway-userId": "0",
    "success": true,
    "payload": true
}

异常错误:
1001 参数错误
2005 商品不存在
4001 添加任务失败
```


### 添加SKU

`post` `/goods/addSku`

| 参数名 | 值 | 必填 | 说明 |
| :------ | :------ | :------ | :------ |
| orgId | int | 是 | 项目id |
| mchId | int64 | 是 | 商家id |
| productId | int64 | 是 | 产品id |
| specs | string | 是 | 具体属性列表 json |
| sortOrder | int | 是 | 排序 |
| price | string | 是 | 价格 |
| stock | int | 是 | 库存 |

```go
返回值:
{
    "gateway-success": true,
    "gateway-orgId": "0",
    "gateway-userId": "0",
    "success": true,
    "payload": {
        "skuId": 3,
        "productId": 1,
        "specs": "{\"颜色\":\"金色\", \"版本\":\"256G\"}",
        "sortOrder": 1,
        "price": "15000",
        "stock": 100,
        "createTime": 1572942482,
        "updateTime": 0
    }
}

异常错误:
1001 参数错误
2003 添加商品属性值失败
```


### 操作sku库存

`post` `/goods/operateStock`

| 参数名 | 值 | 必填 | 说明 |
| :------ | :------ | :------ | :------ |
| orgId | int | 是 | 项目id |
| mchId | int64 | 是 | 商家id |
| productId | int64 | 是 | 产品id |
| skuId | int64 | 是 | skuId |
| qty | int | 是 | 数量 |
| opType | StockOpType | 是 | 操作类型 |

```go
type StockOpType int

AddStock       StockOpType = 1 //加库存
SubStock       StockOpType = 2 //减库存
Freeze         StockOpType = 3 //冻结库存
SubFreezeStock StockOpType = 4 //减冻结库存
UnFreeze       StockOpType = 5 //解冻库存

返回值:
{
    "gateway-success": true,
    "gateway-orgId": "0",
    "gateway-userId": "0",
    "success": true,
    "payload": true
}

异常错误:
1001 参数错误
3002 sku不存在
3003 sku库存不足
3004 添加库存失败
3005 减库存失败
3006 冻结库存失败
3007 减冻结库存失败
3008 解冻库存失败
```
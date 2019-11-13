package base_server_goods

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

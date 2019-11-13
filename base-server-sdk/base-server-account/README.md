# base_server_account 接口说明文档

## 初始化base_server_sdk
```go
base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
    OrgId:           8,
    AppId:           "10008",
    AppSecretKey:    "hiojklsankldlksdnlsdasd",
    RequestTimeout:  5 * time.Second,
    IdleConnTimeout: 10 * time.Minute,
    Hosts: base_server_sdk.Hosts{
        AccountServerHost: "http://127.0.0.1:8081",
    },
    GRpcOnly: false,
})

// ....

defer base_server_sdk.ReleaseBaseServerSdk()
```

## 相关model
```go
type Account struct {
	AccountId    int64  `json:"accountId"`
	OrgId        int    `json:"orgId"`
	UserId       int64  `json:"userId"`
	Currency     string `json:"currency"`
	AvailAmount  string `json:"availAmount"`
	FreezeAmount string `json:"freezeAmount"`
	Status       int    `json:"status"`
	CreateTime   int64  `json:"createTime"`
	UpdateTime   int64  `json:"updateTime"`
}

type LogList struct {
	LogId      int64	`json:"logId"`
	UserId     int64	`json:"userId"`
	Currency   string	`json:"currency"`
	LogType    int		`json:"logType"`
	Amount     string	`json:"amount"`
	CreateTime int64	`json:"createTime"`
}

type TaskDetail struct {
	OpType    int    `json:"opType"`        //操作类型
	BsType    int    `json:"bsType"`        //业务类型
	AccountId int64  `json:"accountId"`     //账户id
	Amount 	  string `json:"amount"`        //金额
    AllowNegative int    `json:"allowNegative"` //是否允许为负数
	UserId    int64  `json:"userId"`        //用户id
	Currency  string `json:"currency"`      //货币类型
	Detail    string `json:"detail"`        //操作详情
	Ext       string `json:"ext"`           //扩展字段
}

type TaskCallBack struct {
	CallBackUrl string            `json:"callBackUrl"`  //回调url
	Data        map[string]string `json:"data"`         //回调时传回数据
}

type base_server_sdk.Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
```

## 相关错误码
```go
1001 参数错误
2001 账户已存在
2002 账户创建失败
2003 账户不存在
2004 更新状态失败
1009 BC操作失败
2005 账户可用增加失败
2007 可用余额不足
2008 解冻失败
2009 账户可用减少失败
2010 账户冻结减少失败
2011 账户日志创建失败
2004 更新状态失败
```

## 接口说明

- 创建账户

func CreateAccount(orgId int, userId int64, currency []string) (*[]Account, *base_server_sdk.Error)

```go
注意:
1. orgId必须大于0

异常错误:
1001 参数错误
2001 账户已存在
2002 账户创建失败
```

- 账户信息

func AccountInfo(orgId int, userId int64, currency string) (*Account, *base_server_sdk.Error)

```go
异常错误:
1001 参数错误
2003 账户不存在
```

- 个人账户信息列表

func AccountsInfo(orgId int, userIds []int64, currency string) (*Account, *base_server_sdk.Error)

```go
异常错误:
1001 参数错误
2003 账户不存在
```

- 所有账户分页列表

func AccountList(orgId int, accountId int64, currency string, beginTime, endTime int64, status, page, limit int) ([]*Account, *base_server_sdk.Error) 

```go
异常错误:
1001 参数错误
```

- 状态变更

func UpdateStatus(orgId int, accountId int64, status int) *base_server_sdk.Error

```go
status状态枚举:
1：正常
2：冻结

异常错误:
1001 参数错误
2003 账户不存在
2004 更新状态失败
```

- 金额操作

func OperateAmount(orgId int, accountId, userId int64, currency string, opType, bsType int, amount, detail, ext, callback *TaskCallBack) *base_server_sdk.Error

```go
注意:
参数accountId与userId,currency两组必须满足一个不为空.同时存在以accountId为主

opType 类型枚举:
1	//可用-加
2	//可用-减
3	//冻结-加
4	//冻结-减
5	//解冻-冻结进可用

bsType 类型为项目特有业务类型


异常错误:
1001 参数错误
2003 账户不存在
1009 BC操作失败
2005 账户可用增加失败
2007 可用余额不足
2008 解冻失败
2009 账户可用减少失败
2010 账户冻结减少失败
2011 账户日志创建失败
```

- 批量金额操作

func BatchOperateAmount(orgId, isAsync int, details []*TaskDetail, callback *TaskCallBack) *base_server_sdk.Error

```go
注意：

isAsync 批量处理方式, 同步/异步 0:同步(默认) 1:异步

opType 类型枚举:
1	//可用-加
2	//可用-减
3	//冻结-加
4	//冻结-减
5	//解冻-冻结进可用

bsType 类型为项目特有业务类型

type TaskDetail struct {
	OpType    int    `json:"opType"`        //操作类型
	BsType    int    `json:"bsType"`        //业务类型
	AccountId int64  `json:"accountId"`     //账户id
	Amount 	  string `json:"amount"`        //金额
	UserId    int64  `json:"userId"`        //用户id
	Currency  string `json:"currency"`      //货币类型
	Detail    string `json:"detail"`        //操作详情
	Ext       string `json:"ext"`           //扩展字段
}

```

- 账户日志列表

func AccountLogList(orgId int, userId int64, opType, bsType int, currency string, beginTime, endTime int, page, limit int) (*[]LogList, *base_server_sdk.Error) 

```go
opType 类型枚举:
1	//可用-加
2	//可用-减
3	//冻结-加
4	//冻结-减
5	//解冻-冻结进可用

bsType 类型为项目特有业务类型


异常错误:
1001 参数错误
```

- 账户日志总和

func SumLog(orgId int, userId int64, opType, bsType int, currency string, beginTime, endTime int) (string, *base_server_sdk.Error) 

```go
opType 类型枚举:
1	//可用-加
2	//可用-减
3	//冻结-加
4	//冻结-减
5	//解冻-冻结进可用

bsType 类型为项目特有业务类型


异常错误:
1001 参数错误
```

- 账户划转
func Transfer(orgId int, fromAccountId, toAccountId int64, amount string) *base_server_sdk.Error 

```go
异常错误:
1001 参数错误
2003 账户不存在
2012 账户币种不同
```
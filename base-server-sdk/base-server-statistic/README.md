# base_server_user 接口说明文档

## 初始化base_server_sdk
```go
base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		AppId:           "10002",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			StatisticServerHost: "127.0.0.1:18081",
		},
		GRpcOnly: true,
	})

// ....

defer base_server_sdk.ReleaseBaseServerSdk()
```

## 相关model
```go
type StatisticItem struct {
	ItemId      int64  `json:"ItemId"`
	OrgId       int    `json:"OrgId"`
	Tag         string `json:"Tag"`
	KeyFields   string `json:"KeyFields"`
	Time        int64  `json:"Time"`
	Granularity string `json:"Granularity"`
	Value       string `json:"Value"`
	CreateTime  int64  `json:"CreateTime"`
	UpdateTime  int64  `json:"UpdateTime"`
}
```

## 相关错误码
```go
1000 服务繁忙
```


## 接口说明

- 查询统计数据

func ListStatisticItems(orgId int, tag string, keyFieldVal string, granularity string, beginTime int, endTime int, limit int, page int) (*[]StatisticItem, *base_server_sdk.Error)

```go
1.orgId 应用ID ,0 查询所有
2.tag 统计标签 ,0 查询所有
3.keyFieldVal 统计关键字值 , 0查询所有
4.granularity 时间粒度 minute,hour,day ,"" 查询所有
5.beginTime 开始时间 ,0 查询所有
6.endTime 结束时间,0 查询所有
7.limit 页大小,默认20
8.page 页码, 0 开始

异常返回:
1000 服务繁忙
```


- 查询今天统计数据

func ListStatisticItems(orgId int, tag string, keyFieldVal string, granularity string, limit int, page int) (*[]StatisticItem, *base_server_sdk.Error)

```go
1.orgId 应用ID ,0 查询所有
2.tag 统计标签 ,0 查询所有
3.keyFieldVal 统计关键字值 , 0查询所有
4.granularity 时间粒度 minute,hour,day ,"" 查询所有
7.limit 页大小,默认20
8.page 页码, 0 开始

异常返回:
1000 服务繁忙
```

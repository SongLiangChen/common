package base_server_statistic

import (
	"encoding/json"
	"github.com/becent/golang-common"
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
	"time"
)

//统计数据
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

//ListStatisticItems 查询统计数据
//orgId 应用ID ,0 查询所有
//tag 统计标签 ,0 查询所有
//keyFieldVal 统计关键字值 , 0查询所有
//granularity 时间粒度 minute,hour,day ,"" 查询所有
//'beginTime 开始时间 ,0 查询所有
//endTime 结束时间,0 查询所有
//limit 页大小,默认20
//page 页码, 0 开始
func ListStatisticItems(orgId int, tag string, keyFieldVal string, granularity string, beginTime int, endTime int, limit int, page int) ([]*StatisticItem, *base_server_sdk.Error) {
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["tag"] = tag
	params["keyFieldVal"] = keyFieldVal
	params["granularity"] = granularity

	params["beginTime"] = strconv.Itoa(beginTime)
	params["endTime"] = strconv.Itoa(endTime)
	params["limit"] = strconv.Itoa(limit)
	params["page"] = strconv.Itoa(page)

	client := base_server_sdk.Instance
	response, err := client.DoRequest(client.Hosts.StatisticServerHost, "statistic", "list", params)
	if err != nil {
		return nil, err
	}

	statisticDatas := &[]*StatisticItem{}
	if err := json.Unmarshal(response, statisticDatas); err != nil {
		common.ErrorLog("baseServerSdk_ListStatisticData", params, "unmarshal fail: "+string(response))
		return nil, base_server_sdk.ErrServiceBusy
	}
	return *statisticDatas, nil
}

//ListStatisticItems 查询统计数据
//orgId 应用ID ,0 查询所有
//tag 统计标签 ,0 查询所有
//keyFieldVal 统计关键字值 , 0查询所有
//granularity 时间粒度 minute,hour,day ,"" 查询所有
//limit 页大小,默认20
//page 页码, 0 开始
func ListTodayStatisticItems(orgId int, tag string, keyFieldVal string, granularity string, limit int, page int) ([]*StatisticItem, *base_server_sdk.Error) {
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["tag"] = tag
	params["keyFieldVal"] = keyFieldVal
	params["granularity"] = granularity

	now := time.Now().Unix()
	params["beginTime"] = strconv.FormatInt(now-(now%86400), 10)
	params["endTime"] = strconv.FormatInt(now-(now%86400)+86400, 10)
	params["limit"] = strconv.Itoa(limit)
	params["page"] = strconv.Itoa(page)

	client := base_server_sdk.Instance
	response, err := client.DoRequest(client.Hosts.StatisticServerHost, "statistic", "list", params)
	if err != nil {
		return nil, err
	}

	statisticDatas := &[]*StatisticItem{}
	if err := json.Unmarshal(response, statisticDatas); err != nil {
		common.ErrorLog("baseServerSdk_ListStatisticData", params, "unmarshal fail: "+string(response))
		return nil, base_server_sdk.ErrServiceBusy
	}
	return *statisticDatas, nil
}

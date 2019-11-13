package base_server_interact

import (
	"encoding/json"
	"github.com/becent/golang-common"
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
	"strings"
)

type InteractRecord struct {
	Id         int64    `json:"id"`
	InteractId int64    `json:"interactId,string"` // 互动id，服务器保证唯一性
	UserId     int64    `json:"userId"`            // 互动的发起人
	NickName   string   `json:"nickName"`          // 当时互动时候的昵称
	Avatar     string   `json:"avatar"`            // 当时互动时候的头像
	OrgId      int      `json:"orgId"`             // orgId
	Type       int      `json:"type"`              // 代表互动的类型 例如 1评论 2点赞 3收藏...有业务方自己决定
	SubType    int      `json:"subType"`           // 子类型，例如评价中的好评、差评
	MediaId    int64    `json:"mediaId"`           // 资源的id，特殊情况下，InteractId也可以成为MediaId，比如对评论进行回复，相同MediaType的MediaId需要保证唯一
	MediaType  string   `json:"mediaType"`         // 资源的类型，可自定义，例如news代表新闻，外部资源必须填，为空代表是评论或者点赞类型的资源
	MediaOwner int64    `json:"mediaOwner"`        // 资源所有者
	CreateTime int64    `json:"createTime"`        // 创建时间
	UpdateTime int64    `json:"updateTime"`        // 更新时间
	Status     int      `json:"status"`            // 状态 1有效 0无效
	Title      string   `json:"title"`             // 标题
	Content    string   `json:"content"`           // 互动内容
	Score      int      `json:"score"`             // 得分，默认100
	Imgs       []string `json:"imgs"`              // 图片集
	At         []string `json:"at"`                // @对象
	Videos     []string `json:"videos"`            // 视频链接
}

type Counting struct {
	Num     map[int64]int  `json:"num"`     // key是mediaId, val是个数
	HasJoin map[int64]bool `json:"hasJoin"` // key是mediaId, val是一个布尔, 代表该用户是否参与过该互动
}

// InteractList 获取互动列表
//
// # 常用业务场景
//
// 现假设有 mediaType=news, comment分别代表新闻和评论资源、type=1代表评论交互类型
//
// - 查询某条新闻的所有评论：mediaType=news、mediaId=xxx(新闻的id)、type=1
//
// - 查询某条评论的所有评论：mediaType=comment、mediaId=xxx(评论的id)、type=1
//
// - 查询某用户对新闻类产品做出的所有评论：mediaType=news、userId=xxx(用户id)、type=1
//
// - 查询某用户对某条新闻的所有评论：mediaType=news、mediaId=xxx(新闻的id)、userId=xxx(用户id)、type=1
//
// - 查询某用户发出的所有评论：userId=xxx(用户id)、type=1
//
// - 查询某人收到的所有评论：mediaOwner=xxx(某人id)、type=1
//
// - 查询某人收到的所有新闻类的评论：mediaOwner=xxx(某人id)、mediaType=news、type=1
func InteractList(orgId int, userId int64, mediaId int64, mediaOwner int64, ttype, subType int, orderType int, page int, limit int) ([]*InteractRecord, *base_server_sdk.Error) {
	if orgId <= 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 20
	}
	orderType = 0 // TODO 目前只支持时间降序

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["mediaId"] = strconv.FormatInt(mediaId, 10)
	params["mediaOwner"] = strconv.FormatInt(mediaOwner, 10)
	params["type"] = strconv.Itoa(ttype)
	params["subType"] = strconv.Itoa(subType)
	params["orderType"] = strconv.Itoa(orderType)
	params["page"] = strconv.Itoa(page)
	params["limit"] = strconv.Itoa(limit)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.InteractServerHost, "interact", "list", params)
	if err != nil {
		return nil, err
	}

	records := []*InteractRecord{}
	if err := json.Unmarshal(data, &records); err != nil {
		common.ErrorLog("baseServerSdk_InteractList", params, "unmarshal records info fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return records, nil
}

// AddInteract 添加一条互动记录
//
// orgId, userId, mediaId, mediaType, mediaOwner, type等参数必填
// score默认为100分
func AddInteract(record *InteractRecord, atomicity int) (*InteractRecord, *base_server_sdk.Error) {
	if record == nil {
		return nil, base_server_sdk.ErrInvalidParams
	}

	if record.OrgId <= 0 || record.UserId <= 0 || record.MediaId <= 0 || record.MediaOwner <= 0 || record.MediaType == "" || record.Type <= 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(record.OrgId)
	params["userId"] = strconv.FormatInt(record.UserId, 10)
	params["mediaId"] = strconv.FormatInt(record.MediaId, 10)
	params["mediaType"] = record.MediaType
	params["mediaOwner"] = strconv.FormatInt(record.MediaOwner, 10)
	params["type"] = strconv.Itoa(record.Type)
	params["subType"] = strconv.Itoa(record.SubType)
	params["nickName"] = record.NickName
	params["avatar"] = record.Avatar
	params["title"] = record.Title
	params["content"] = record.Content
	if record.Score > 0 {
		params["score"] = strconv.Itoa(record.Score)
	}
	if len(record.Imgs) > 0 {
		params["imgs"] = strings.Join(record.Imgs, ",")
	}
	if len(record.At) > 0 {
		params["at"] = strings.Join(record.At, ",")
	}
	if len(record.Videos) > 0 {
		params["videos"] = strings.Join(record.Videos, ",")
	}
	params["atomicity"] = strconv.Itoa(atomicity)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.InteractServerHost, "interact", "add", params)
	if err != nil {
		return nil, err
	}
	println(string(data))

	record = &InteractRecord{}
	if err := json.Unmarshal(data, record); err != nil {
		common.ErrorLog("baseServerSdk_AddInteract", params, "unmarshal record info fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return record, nil
}

// DelInteract 删除一条互动记录
//
// 所有参数必填
func DelInteract(orgId int, userId int64, interactId int64) *base_server_sdk.Error {
	if orgId <= 0 || userId <= 0 || interactId <= 0 {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["interactId"] = strconv.FormatInt(interactId, 10)

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.InteractServerHost, "interact", "del", params)
	if err != nil {
		return err
	}

	return nil
}

// InteractDetail 查看互动记录详情
func InteractDetail(orgId int, interactId int64) (*InteractRecord, *base_server_sdk.Error) {
	if orgId <= 0 || interactId <= 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["interactId"] = strconv.FormatInt(interactId, 10)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.InteractServerHost, "interact", "detail", params)
	if err != nil {
		return nil, err
	}

	record := &InteractRecord{}
	if err := json.Unmarshal(data, record); err != nil {
		common.ErrorLog("baseServerSdk_AddInteract", params, "unmarshal record info fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return record, nil
}

// InteractScore 查看某资源的平均得分
func InteractScore(orgId int, mediaId int64, mediaType string) (float64, *base_server_sdk.Error) {
	if orgId <= 0 || mediaId <= 0 || mediaType == "" {
		return 0, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["mediaId"] = strconv.FormatInt(mediaId, 10)
	params["mediaType"] = mediaType

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.InteractServerHost, "interact", "score", params)
	if err != nil {
		return 0, err
	}

	type Score struct {
		Score float64 `json:"score"`
	}

	s := &Score{}
	if err := json.Unmarshal(data, s); err != nil {
		common.ErrorLog("baseServerSdk_InteractScore", params, "unmarshal score info fail: "+string(data))
		return 0, base_server_sdk.ErrServiceBusy
	}

	return s.Score, nil
}

// InteractCounting 查看资源计数信息
//
// 注意:该接口只支持同类型mediaType进行批量查询,一次查询最多40条资源.查看的互动类型每次不超过10种.
//
// 例如可以同时查看新闻资源1000, 10001, 1002的点赞个数,评论个数: mediaType=news  mediaIds=[]int64{1000, 1001, 1002}  types=[]int{1, 2}
//
// 返回值map[int]*Counting，key代表互动类型，val代表该互动类型的计数详情(详见Counting结构定义)
func InteractCounting(orgId int, userId int64, mediaType string, mediaIds []int64, types []int) (map[int]*Counting, *base_server_sdk.Error) {
	if orgId <= 0 || len(mediaIds) == 0 || len(mediaIds) > 40 || len(types) == 0 || len(types) > 10 || mediaType == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["mediaIds"] = strings.Join(common.Int64SliceToStringSlice(mediaIds), ",")
	params["mediaType"] = mediaType
	params["userId"] = strconv.FormatInt(userId, 10)
	params["types"] = strings.Join(common.IntSliceToStringSlice(types), ",")

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.InteractServerHost, "interact", "counting", params)
	if err != nil {
		return nil, err
	}

	ret := make(map[int]*Counting)
	if err := json.Unmarshal(data, &ret); err != nil {
		common.ErrorLog("baseServerSdk_InteractCounting", params, "unmarshal counting info fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return ret, nil
}

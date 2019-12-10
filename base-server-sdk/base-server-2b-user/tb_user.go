package base_server_2b_user

import (
	"encoding/json"
	"github.com/becent/golang-common"
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
	"strings"
)

type User2B struct {
	Id                    int64  `json:"id"`
	OrgId                 int    `json:"orgId"`                 // 组织项目id
	UserId                int64  `json:"userId"`                // 会员id
	DirectParentUserId    int64  `json:"directParentUserId"`    // 一级邀请用户id
	SecondParentUserId    int64  `json:"secondParentUserId"`    // 二级邀请用户id
	DirectInviteUserNum   int    `json:"directInviteUserNum"`   // 该用户直接邀请的用户人数
	IndirectInviteUserNum int    `json:"indirectInviteUserNum"` // 该用户间接邀请的用户人数
	ParentAdminId         int64  `json:"parentAdminId"`         // 父级会员id
	InviteCode            string `json:"inviteCode"`            // 用户的邀请码
	CreateTime            int64  `json:"createTime"`            // 创建时间

	Lv1AdminId int64 `json:"lv1AdminId"` // 所属1级会员id
	Lv2AdminId int64 `json:"lv2AdminId"` // 所属2级会员id
	Lv3AdminId int64 `json:"lv3AdminId"` // 所属3级会员id
	Lv4AdminId int64 `json:"lv4AdminId"` // 所属4级会员id
	Lv5AdminId int64 `json:"lv5AdminId"` // 所属5级会员id
	Lv6AdminId int64 `json:"lv6AdminId"` // 所属6级会员id
	Lv7AdminId int64 `json:"lv7AdminId"` // 所属7级会员id
	Lv8AdminId int64 `json:"lv8AdminId"` // 所属8级会员id
	Lv9AdminId int64 `json:"lv9AdminId"` // 所属9级会员id

	CountryCode      string `json:"countryCode"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	Account          string `json:"account"`
	ThirdPartyOpenId string `json:"thirdPartyOpenId"`
	ThirdPartyType   int    `json:"thirdPartyType"`
	LoginPwd         string `json:"loginPwd"`
	TransPwd         string `json:"transPwd"`
	NickName         string `json:"nickName"`
	Avatar           string `json:"avatar"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	IdCard           string `json:"idCard"`
	Sex              int    `json:"sex"`
	BirthDay         string `json:"birthDay"`
	Status           int    `json:"status"`
	Ext              string `json:"ext"`
}

type Admin2B struct {
	Id            int64  `json:"id"`
	OrgId         int    `json:"orgId"`         // 组织项目id
	AdminId       int64  `json:"adminId"`       // 会员id
	RoleId        int    `json:"roleId"`        // 角色id
	IsSub         int    `json:"isSub"`         // 是否是子账号
	ParentAdminId int64  `json:"parentAdminId"` // 父级会员id
	InviteCode    string `json:"inviteCode"`    // 会员邀请码
	CreateTime    int64  `json:"createTime"`    // 创建时间

	Lv1AdminId int64 `json:"lv1AdminId"` // 所属1级会员id
	Lv2AdminId int64 `json:"lv2AdminId"` // 所属2级会员id
	Lv3AdminId int64 `json:"lv3AdminId"` // 所属3级会员id
	Lv4AdminId int64 `json:"lv4AdminId"` // 所属4级会员id
	Lv5AdminId int64 `json:"lv5AdminId"` // 所属5级会员id
	Lv6AdminId int64 `json:"lv6AdminId"` // 所属6级会员id
	Lv7AdminId int64 `json:"lv7AdminId"` // 所属7级会员id
	Lv8AdminId int64 `json:"lv8AdminId"` // 所属8级会员id
	Lv9AdminId int64 `json:"lv9AdminId"` // 所属9级会员id

	CountryCode      string `json:"countryCode"`
	Phone            string `json:"phone"`
	Email            string `json:"email"`
	Account          string `json:"account"`
	ThirdPartyOpenId string `json:"thirdPartyOpenId"`
	ThirdPartyType   int    `json:"thirdPartyType"`
	LoginPwd         string `json:"loginPwd"`
	TransPwd         string `json:"transPwd"`
	NickName         string `json:"nickName"`
	Avatar           string `json:"avatar"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	IdCard           string `json:"idCard"`
	Sex              int    `json:"sex"`
	BirthDay         string `json:"birthDay"`
	Status           int    `json:"status"`
	Ext              string `json:"ext"`
}

func RegisterUser(user *User2B, code string, currencyTypes []string) (*User2B, *base_server_sdk.Error) {
	if user.OrgId <= 0 || user.InviteCode == "" || (user.Phone == "" && user.Account == "" && user.Email == "") || (user.Account != "" && user.LoginPwd == "") {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(user.OrgId)
	params["countryCode"] = user.CountryCode
	params["phone"] = user.Phone
	params["email"] = user.Email
	params["account"] = user.Account
	params["loginPwd"] = user.LoginPwd
	params["transPwd"] = user.TransPwd
	params["nickName"] = user.NickName
	params["avatar"] = user.Avatar
	params["sex"] = strconv.Itoa(user.Sex)
	params["birthDay"] = user.BirthDay
	params["ext"] = user.Ext
	params["inviteCode"] = user.InviteCode
	params["code"] = code
	if currencyTypes != nil && len(currencyTypes) > 0 {
		params["currencyTypes"] = strings.Join(currencyTypes, ",")
	}

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.TbUserServerHost, "user", "register", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, user); err != nil {
		common.ErrorLog("baseServerSdk_RegisterUser", params, "unmarshal user fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return user, nil
}

func GetUserInfo(orgId int, userId int64, baseInfo int) (*User2B, *base_server_sdk.Error) {
	if orgId <= 0 || userId <= 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["baseInfo"] = strconv.Itoa(baseInfo)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.TbUserServerHost, "user", "getUserInfo", params)
	if err != nil {
		return nil, err
	}
	user := &User2B{}
	if err := json.Unmarshal(data, user); err != nil {
		common.ErrorLog("baseServerSdk_GetUserInfo", params, "unmarshal user fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return user, nil
}

func GetUsersInfo(orgId int, userIds []int64, baseInfo int) (map[int64]*User2B, *base_server_sdk.Error) {
	if orgId <= 0 || len(userIds) == 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userIds"] = strings.Join(common.Int64SliceToStringSlice(userIds), ",")
	params["baseInfo"] = strconv.Itoa(baseInfo)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.TbUserServerHost, "user", "getUsersInfo", params)
	if err != nil {
		return nil, err
	}
	users := make(map[int64]*User2B)
	if err := json.Unmarshal(data, &users); err != nil {
		common.ErrorLog("baseServerSdk_GetUsersInfo", params, "unmarshal users fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return users, nil
}

func FindUserList(orgId int, userId int64, nickName, countryCode, phone, email, account string,
	parentAdminId, affiliatedAdminId, beginTime, endTime int64, page, pageSize, baseInfo int) ([]*User2B, *base_server_sdk.Error) {
	if orgId <= 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["nickName"] = nickName
	params["countryCode"] = countryCode
	params["phone"] = phone
	params["email"] = email
	params["account"] = account
	params["parentAdminId"] = strconv.FormatInt(parentAdminId, 10)
	params["affiliatedAdminId"] = strconv.FormatInt(affiliatedAdminId, 10)
	params["beginTime"] = strconv.FormatInt(beginTime, 10)
	params["endTime"] = strconv.FormatInt(endTime, 10)
	params["page"] = strconv.Itoa(page)
	params["pageSize"] = strconv.Itoa(pageSize)
	params["baseInfo"] = strconv.Itoa(baseInfo)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.TbUserServerHost, "user", "findList", params)
	if err != nil {
		return nil, err
	}
	list := make([]*User2B, 0)
	if err := json.Unmarshal(data, &list); err != nil {
		common.ErrorLog("baseServerSdk_FindUserList", params, "unmarshal user list fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return list, nil
}

func RegisterAdmin(admin *Admin2B, code string, currencyTypes []string) (*Admin2B, *base_server_sdk.Error) {
	if admin.OrgId <= 0 || admin.ParentAdminId <= 0 || (admin.Phone == "" && admin.Account == "" && admin.Email == "") || (admin.Account != "" && admin.LoginPwd == "") {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(admin.OrgId)
	params["countryCode"] = admin.CountryCode
	params["phone"] = admin.Phone
	params["email"] = admin.Email
	params["account"] = admin.Account
	params["loginPwd"] = admin.LoginPwd
	params["transPwd"] = admin.TransPwd
	params["nickName"] = admin.NickName
	params["avatar"] = admin.Avatar
	params["sex"] = strconv.Itoa(admin.Sex)
	params["birthDay"] = admin.BirthDay
	params["ext"] = admin.Ext
	params["parentAdminId"] = strconv.FormatInt(admin.ParentAdminId, 10)
	params["isSub"] = strconv.Itoa(admin.IsSub)
	params["code"] = code
	if currencyTypes != nil && len(currencyTypes) > 0 {
		params["currencyTypes"] = strings.Join(currencyTypes, ",")
	}

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.TbUserServerHost, "admin", "register", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, admin); err != nil {
		common.ErrorLog("baseServerSdk_RegisterAdmin", params, "unmarshal admin fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return admin, nil
}

func GetAdminInfo(orgId int, adminId int64, baseInfo int) (*Admin2B, *base_server_sdk.Error) {
	if orgId <= 0 || adminId <= 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["adminId"] = strconv.FormatInt(adminId, 10)
	params["baseInfo"] = strconv.Itoa(baseInfo)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.TbUserServerHost, "admin", "getAdminInfo", params)
	if err != nil {
		return nil, err
	}
	admin := &Admin2B{}
	if err := json.Unmarshal(data, admin); err != nil {
		common.ErrorLog("baseServerSdk_GetAdminInfo", params, "unmarshal admin fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return admin, nil
}

func GetAdminsInfo(orgId int, adminIds []int64, baseInfo int) (map[int64]*Admin2B, *base_server_sdk.Error) {
	if orgId <= 0 || len(adminIds) == 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["adminIds"] = strings.Join(common.Int64SliceToStringSlice(adminIds), ",")
	params["baseInfo"] = strconv.Itoa(baseInfo)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.TbUserServerHost, "admin", "getAdminsInfo", params)
	if err != nil {
		return nil, err
	}
	admins := make(map[int64]*Admin2B)
	if err := json.Unmarshal(data, &admins); err != nil {
		common.ErrorLog("baseServerSdk_GetAdminsInfo", params, "unmarshal admins fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return admins, nil
}

func FindAdminList(orgId int, adminId int64, nickName, countryCode, phone, email, account string,
	roleId, isSub int, parentAdminId, affiliatedAdminId, beginTime, endTime int64, page, pageSize, baseInfo int) ([]*Admin2B, *base_server_sdk.Error) {
	if orgId <= 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["adminId"] = strconv.FormatInt(adminId, 10)
	params["nickName"] = nickName
	params["countryCode"] = countryCode
	params["phone"] = phone
	params["email"] = email
	params["account"] = account
	params["roleId"] = strconv.Itoa(roleId)
	params["isSub"] = strconv.Itoa(isSub)
	params["parentAdminId"] = strconv.FormatInt(parentAdminId, 10)
	params["affiliatedAdminId"] = strconv.FormatInt(affiliatedAdminId, 10)
	params["beginTime"] = strconv.FormatInt(beginTime, 10)
	params["endTime"] = strconv.FormatInt(endTime, 10)
	params["page"] = strconv.Itoa(page)
	params["pageSize"] = strconv.Itoa(pageSize)
	params["baseInfo"] = strconv.Itoa(baseInfo)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.TbUserServerHost, "admin", "findList", params)
	if err != nil {
		return nil, err
	}
	list := make([]*Admin2B, 0)
	if err := json.Unmarshal(data, &list); err != nil {
		common.ErrorLog("baseServerSdk_FindAdminList", params, "unmarshal admin list fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return list, nil
}

func RelationTransfer(orgId int, fromId, toId int64, ttype int, delayTime int64) *base_server_sdk.Error {
	if orgId <= 0 || fromId <= 0 || toId <= 0 || ttype <= 0 {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["fromId"] = strconv.FormatInt(fromId, 10)
	params["toId"] = strconv.FormatInt(toId, 10)
	params["type"] = strconv.Itoa(ttype)
	params["delayTime"] = strconv.FormatInt(delayTime, 10)

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.TbUserServerHost, "relation", "transfer", params)
	return err
}

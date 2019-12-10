package base_server_user

import (
	"github.com/becent/golang-common"
	"github.com/becent/golang-common/base-server-sdk"
	json "github.com/json-iterator/go"
	"strconv"
	"strings"
)

type User struct {
	UserId           int64  `json:"userId"`
	OrgId            int    `json:"orgId"`
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
	CreateTime       int64  `json:"createTime"`
	Ext              string `json:"ext"`
}

func ReserveUserId() (int64, *base_server_sdk.Error) {
	params := make(map[string]string)
	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "reserveUserId", params)
	if err != nil {
		return 0, err
	}

	ret := make(map[string]string)
	if err := json.Unmarshal(data, &ret); err != nil {
		common.ErrorLog("baseServerSdk_ReserveUserId", params, "unmarshal data fail: "+string(data))
		return 0, base_server_sdk.ErrServiceBusy
	}

	userId, e := strconv.ParseInt(ret["userId"], 10, 64)
	if e != nil {
		return 0, base_server_sdk.ErrServiceBusy
	}

	return userId, nil
}

// Register 注册用户
//
// 1、orgId必须大于0
// 2、account、phone、email至少要有一个有值
// 3、account有值时，password必须有值
// 4、code有值时，将会校验短信或者邮件验证码
// 5、其他字段非必填
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1002 账户已经注册
func Register(user *User, code string, currencyTypes []string) (*User, *base_server_sdk.Error) {
	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(user.OrgId)
	params["userId"] = strconv.FormatInt(user.UserId, 10)
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
	params["code"] = code
	if currencyTypes != nil && len(currencyTypes) > 0 {
		params["currencyTypes"] = strings.Join(currencyTypes, ",")
	}

	if params["orgId"] == "0" || (params["phone"] == "" && params["account"] == "" && params["email"] == "") || (params["account"] != "" && params["loginPwd"] == "") {
		return nil, base_server_sdk.ErrInvalidParams
	}

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "register", params)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, user); err != nil {
		common.ErrorLog("baseServerSdk_Register", params, "unmarshal user fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return user, nil
}

// LoginByAccount 通过账号进行登录
//
// 1、password有值时，会进行登录密码校验，为空则略过
//
// 异常返回：
// 1003 账号不存在
// 1004 用户已被冻结
// 1005 密码错误
func LoginByAccount(orgId int, account string, password string) (*User, *base_server_sdk.Error) {
	if orgId <= 0 || account == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["account"] = account
	params["password"] = password

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "loginByAccount", params)
	if err != nil {
		return nil, err
	}

	user := &User{}
	if err := json.Unmarshal(data, user); err != nil {
		common.ErrorLog("baseServerSdk_LoginByAccount", params, "unmarshal user fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return user, nil
}

// LoginByPhone 通过手机号进行登录
//
// 1、code有值时，会进行短信校验，为空则略过
// 2、password有值时，会进行登录密码校验，为空则略过
//
// 异常返回：
// 1003 账号不存在
// 1004 用户已被冻结
// 1005 密码错误
// 1008 短信验证码错误
func LoginByPhone(orgId int, countryCode, phone string, code string, password string) (*User, *base_server_sdk.Error) {
	if orgId <= 0 || phone == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["countryCode"] = countryCode
	params["phone"] = phone
	params["code"] = code
	params["password"] = password

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "loginByPhone", params)
	if err != nil {
		return nil, err
	}

	user := &User{}
	if err := json.Unmarshal(data, user); err != nil {
		common.ErrorLog("baseServerSdk_LoginByPhone", params, "unmarshal user fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return user, nil
}

// LoginByEmail 通过手机号进行登录
//
// 1、code有值时，会进行邮件校验，为空则略过
// 2、password有值时，会进行登录密码校验，为空则略过
//
// 异常返回：
// 1003 账号不存在
// 1004 用户已被冻结
// 1005 密码错误
// 1009 邮件验证码错误
func LoginByEmail(orgId int, email string, code string, password string) (*User, *base_server_sdk.Error) {
	if orgId <= 0 || email == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["email"] = email
	params["code"] = code
	params["password"] = password

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "loginByEmail", params)
	if err != nil {
		return nil, err
	}

	user := &User{}
	if err := json.Unmarshal(data, user); err != nil {
		common.ErrorLog("baseServerSdk_LoginByEmail", params, "unmarshal user fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return user, nil
}

// GetUserInfo 获取用户信息
//
// 异常返回：
// 1003 用户不存在
func GetUserInfo(orgId int, userId int64) (*User, *base_server_sdk.Error) {
	if orgId <= 0 || userId <= 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "getUserInfo", params)
	if err != nil {
		return nil, err
	}

	user := &User{}
	if err := json.Unmarshal(data, user); err != nil {
		common.ErrorLog("baseServerSdk_GetUserInfo", params, "unmarshal user fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return user, nil
}

// GetUsersInfo 批量获取用户信息
//
// 返回值：
// []*User 查询到的用户
// []int64 未找到的用户
func GetUsersInfo(orgId int, userIds []int64) (map[int64]*User, []int64, *base_server_sdk.Error) {
	if orgId <= 0 || len(userIds) == 0 {
		return nil, nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	strUserIds := make([]string, 0)
	for _, userId := range userIds {
		strUserIds = append(strUserIds, strconv.FormatInt(userId, 10))
	}
	params["userIds"] = strings.Join(strUserIds, ",")

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "getUsersInfo", params)
	if err != nil {
		return nil, nil, err
	}

	type Info struct {
		Users    map[int64]*User `json:"users"`
		NotFound []int64         `json:"notFound"`
	}

	info := &Info{}
	if err := json.Unmarshal(data, info); err != nil {
		common.ErrorLog("baseServerSdk_GetUsersInfo", params, "unmarshal users info fail: "+string(data))
		return nil, nil, base_server_sdk.ErrServiceBusy
	}

	return info.Users, info.NotFound, nil
}

// GetBackLoginPwdByPhone 通过手机找回登录密码
//
// 1、code有值的话会进行短信校验，为空则忽略
// 2、password是要设置的新密码
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
// 1008 短信验证码错误
func GetBackLoginPwdByPhone(orgId int, countryCode, phone, code, password string) *base_server_sdk.Error {
	if orgId <= 0 || phone == "" || password == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["countryCode"] = countryCode
	params["phone"] = phone
	params["code"] = code
	params["password"] = password

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "getBackLoginPwd", params)
	if err != nil {
		return err
	}

	return nil
}

// GetBackLoginPwdByEmail 通过邮箱找回登录密码
//
// 1、code有值的话会进行邮箱校验，为空则忽略
// 2、password是要设置的新密码
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
// 1009 邮箱验证码错误
func GetBackLoginPwdByEmail(orgId int, email, code, password string) *base_server_sdk.Error {
	if orgId <= 0 || email == "" || password == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["email"] = email
	params["code"] = code
	params["password"] = password

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "getBackLoginPwd", params)
	if err != nil {
		return err
	}

	return nil
}

// GetBackTransPwdByPhone 通过手机找回支付密码
//
// 1、code有值的话会进行短信校验，为空则忽略
// 2、password是要设置的新密码
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
// 1008 短信验证码错误
func GetBackTransPwdByPhone(orgId int, countryCode, phone, code, password string) *base_server_sdk.Error {
	if orgId <= 0 || phone == "" || password == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["countryCode"] = countryCode
	params["phone"] = phone
	params["code"] = code
	params["password"] = password

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "getBackTransPwd", params)
	if err != nil {
		return err
	}

	return nil
}

// GetBackTransPwdByEmail 通过邮箱找回支付密码
//
// 1、code有值的话会进行邮箱校验，为空则忽略
// 2、password是要设置的新密码
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
// 1009 邮箱验证码错误
func GetBackTransPwdByEmail(orgId int, email, code, password string) *base_server_sdk.Error {
	if orgId <= 0 || email == "" || password == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["email"] = email
	params["code"] = code
	params["password"] = password

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "getBackTransPwd", params)
	if err != nil {
		return err
	}

	return nil
}

// UpdateLoginPwd 更新登录密码
//
// oldPassword有值则校验旧密码，为空则略过
// newPassword必须有值
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
func UpdateLoginPwd(orgId int, userId int64, oldPassword, newPassword string) *base_server_sdk.Error {
	if orgId <= 0 || userId <= 0 || newPassword == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["oldPassword"] = oldPassword
	params["newPassword"] = newPassword

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "updateLoginPwd", params)
	if err != nil {
		return err
	}

	return nil
}

// UpdateTransPwd 更新交易密码
//
// oldPassword有值则校验旧密码，为空则略过
// newPassword必须有值
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
func UpdateTransPwd(orgId int, userId int64, oldPassword, newPassword string) *base_server_sdk.Error {
	if orgId <= 0 || userId <= 0 || newPassword == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["oldPassword"] = oldPassword
	params["newPassword"] = newPassword

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "updateTransPwd", params)
	if err != nil {
		return err
	}

	return nil
}

// AuthRealName 实名认证
//
// idCard有值的话，会进行身份证校验，为空则忽略
// 真实姓名为firstName+lastName
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1010 身份证认证失败
func AuthRealName(orgId int, userId int64, firstName, lastName, idCard string) *base_server_sdk.Error {
	if orgId <= 0 || userId <= 0 {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["firstName"] = firstName
	params["lastName"] = lastName
	params["idCard"] = idCard

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "authRealName", params)
	if err != nil {
		return err
	}

	return nil
}

type Sex int

const (
	_ Sex = iota
	Boy
	Girl
)

type UserFields map[string]string

func (u UserFields) SetNickName(nickName string) {
	u["nickName"] = nickName
}

func (u UserFields) SetAvatar(avatar string) {
	u["avatar"] = avatar
}

func (u UserFields) SetSex(sex Sex) {
	u["sex"] = strconv.Itoa(int(sex))
}

func (u UserFields) SetBirthDay(birthDay string) {
	u["birthDay"] = birthDay
}

func (u UserFields) SetExt(ext string) {
	u["ext"] = ext
}

// UpdateUserInfo 更新用户信息
//
// 可以更新的字段为：
// 1、nickName
// 2、avatar
// 3、sex
// 4、birthDay
// 5、ext
//
// 用例:
// info := make(base_server_user.UserFields)
// info.SetNickName("jak")
// info.SetSex(base_server_user.Boy)
//
// base_server_user.UpdateUserInfo(1, 1000, info)
//
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
func UpdateUserInfo(orgId int, userId int64, info UserFields) (*User, *base_server_sdk.Error) {
	if orgId <= 0 || userId <= 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	for key, val := range info {
		params[string(key)] = val
	}

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "updateUserInfo", params)
	if err != nil {
		return nil, err
	}

	user := &User{}
	if err := json.Unmarshal(data, user); err != nil {
		common.ErrorLog("baseServerSdk_LoginByAccount", params, "unmarshal user fail: "+string(data))
		return nil, base_server_sdk.ErrServiceBusy
	}

	return user, nil
}

// BindAccount 绑定登录账号
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1006 重复绑定
func BindAccount(orgId int, userId int64, account, password string) *base_server_sdk.Error {
	if orgId <= 0 || userId <= 0 || account == "" || password == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["account"] = account
	params["password"] = password

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "bindAccount", params)
	if err != nil {
		return err
	}

	return nil
}

// BindPhone 绑定手机
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1006 重复绑定
func BindPhone(orgId int, userId int64, countryCode, phone, code string) *base_server_sdk.Error {
	if orgId <= 0 || userId <= 0 || phone == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["countryCode"] = countryCode
	params["phone"] = phone
	params["code"] = code

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "bindPhone", params)
	if err != nil {
		return err
	}

	return nil
}

// BindEmail 绑定邮箱
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1006 重复绑定
func BindEmail(orgId int, userId int64, email, code string) *base_server_sdk.Error {
	if orgId <= 0 || userId <= 0 || email == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["email"] = email
	params["code"] = code

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "bindEmail", params)
	if err != nil {
		return err
	}

	return nil
}

// AuthTransPwd 校验支付密码
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1005 密码错误
func AuthTransPwd(orgId int, userId int64, password string) *base_server_sdk.Error {
	if orgId <= 0 || userId <= 0 || password == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["password"] = password

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "authTransPwd", params)
	if err != nil {
		return err
	}

	return nil
}

type Status int

const (
	Normal Status = 1 + iota
	Forbidden
)

// UpdateUserStatus 修改用户状态
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
func UpdateUserStatus(orgId int, userId int64, status Status) *base_server_sdk.Error {
	if orgId <= 0 || userId <= 0 || (status != Normal && status != Forbidden) {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["status"] = strconv.Itoa(int(status))

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "updateStatus", params)
	if err != nil {
		return err
	}

	return nil
}

// UnBindPhone 解绑手机
//
// 1、code有值的话，会校验短信验证码
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1011 手机未绑定
// 1013 解绑手机后将失去所有登录方式
func UnBindPhone(orgId int, userId int64, code string) *base_server_sdk.Error {
	if orgId <= 0 || userId <= 0 {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["code"] = code

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "unBindPhone", params)
	if err != nil {
		return err
	}

	return nil
}

// UnBindEmail 解绑邮箱
//
// 1、code有值的话，会校验邮箱验证码
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
// 1003 用户不存在
// 1012 邮箱未绑定
// 1014 解绑邮箱后将失去所有登录方式
func UnBindEmail(orgId int, userId int64, code string) *base_server_sdk.Error {
	if orgId <= 0 || userId <= 0 {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["code"] = code

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "unBindEmail", params)
	if err != nil {
		return err
	}

	return nil
}

// StoreValAtomic 原子操作，存储用户一些信息
//
// 异常返回：
// 1000 服务繁忙
// 1001 参数异常
func StoreValAtomic(orgId int, userId int64, key, val string) *base_server_sdk.Error {
	if orgId <= 0 || userId <= 0 || key == "" || val == "" {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["key"] = key
	params["val"] = val

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "storeValAtomic", params)
	if err != nil {
		return err
	}

	return nil
}

// GetStoreVal 获取用户存储的信息
func GetStoreVal(orgId int, userId int64, key string) (map[int64]string, *base_server_sdk.Error) {
	if orgId <= 0 || userId <= 0 || key == "" {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["key"] = key

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.UserServerHost, "user", "getStoreVal", params)
	if err != nil {
		return nil, err
	}

	vals := make(map[int64]string, 0)
	if err := json.Unmarshal(data, &vals); err != nil {
		return nil, base_server_sdk.ErrServiceBusy
	}

	return vals, nil
}

func DelStoreVal(orgId int, userId int64, id int64) *base_server_sdk.Error {
	if orgId <= 0 || userId <= 0 || id <= 0 {
		return base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)
	params["userId"] = strconv.FormatInt(userId, 10)
	params["id"] = strconv.FormatInt(id, 10)

	client := base_server_sdk.Instance
	_, err := client.DoRequest(client.Hosts.UserServerHost, "user", "delStoreVal", params)
	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"fmt"
	"github.com/becent/golang-common/base-server-sdk"
	"github.com/becent/golang-common/base-server-sdk/base-server-user"
	"time"
)

func main() {
	base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		AppId:           "10002",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			UserServerHost: "http://127.0.0.1:8081",
		},
		GRpcOnly: false,
	})
	defer base_server_sdk.ReleaseBaseServerSdk()

	now := time.Now()
	defer func(now time.Time) {
		println(time.Since(now).String())
	}(now)

	// 批量查询用户
	users, notFound, err := base_server_user.GetUsersInfo(2, []int64{10000001, 10000002})
	if err != nil {
		println(err.String())
	} else {
		for _, u := range users {
			fmt.Printf("找到用户 %v\n", *u)
		}
		for _, u := range notFound {
			fmt.Printf("未找到 %v\n", u)
		}
	}

	userId, err := base_server_user.ReserveUserId()
	if err != nil {
		println(err.String())
		return
	}
	fmt.Printf("预留userId %v\n", userId)

	// 注册用户
	user, err := base_server_user.Register(&base_server_user.User{
		OrgId:       20,
		UserId:      userId,
		CountryCode: "+88",
		Phone:       "13560487593",
		LoginPwd:    "123456",
		NickName:    "song",
		Avatar:      "shuai.png",
		Ext:         "123",
	}, "", nil)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("注册成功：[%v]\n", *user)
	}

	return

	// 通过手机找回登录密码
	err = base_server_user.GetBackLoginPwdByPhone(5, "+86", "13560487593", "", "654321")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("找回登录密码成功\n")
	}

	// 登录
	user, err = base_server_user.LoginByPhone(5, "+86", "13560487593", "", "654321")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("登录成功：[%v]\n", *user)
	}

	// 获取用户信息
	user, err = base_server_user.GetUserInfo(5, user.UserId)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("获取用户信息成功：[%v]\n", *user)
	}

	// 修改登录密码
	err = base_server_user.UpdateLoginPwd(user.OrgId, user.UserId, "654321", "123456")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("登录密码修改成功\n")
	}

	// 再次登录
	user, err = base_server_user.LoginByPhone(5, "+86", "13560487593", "", "654321")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("登录成功：[%v]\n", *user)
	}

	// 改变密码再次登录
	user, err = base_server_user.LoginByPhone(5, "+86", "13560487593", "", "123456")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("登录成功：[%v]\n", *user)
	}

	// 实名认证
	if err = base_server_user.AuthRealName(5, user.UserId, "song", "liang", "360721199001040204"); err != nil {
		println(err.String())
	} else {
		println("实名认证成功")
	}

	// 获取用户信息
	user, err = base_server_user.GetUserInfo(5, user.UserId)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("获取用户信息成功：[%v]\n", *user)
	}

	// 通过手机找回交易密码
	err = base_server_user.GetBackTransPwdByPhone(5, "+86", "13560487593", "", "asdqwe")
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("找回交易密码成功\n")
	}

	// 验证交易密码
	if err = base_server_user.AuthTransPwd(5, user.UserId, "123456"); err != nil {
		println(err.String())
	} else {
		println("验证交易密码成功")
	}

	// 验证交易密码
	if err = base_server_user.AuthTransPwd(5, user.UserId, "asdqwe"); err != nil {
		println(err.String())
	} else {
		println("验证交易密码成功")
	}

	// 更新交易密码
	if err = base_server_user.UpdateTransPwd(5, user.UserId, "asdqwe", "123465"); err != nil {
		println(err.String())
	} else {
		println("更新交易密码成功")
	}

	info := make(base_server_user.UserFields)
	info.SetNickName("新昵称")
	info.SetBirthDay("2000-10-10")
	info.SetAvatar("new.png")
	info.SetExt("456")
	info.SetSex(base_server_user.Boy)
	if err = base_server_user.UpdateUserInfo(5, user.UserId, info); err != nil {
		println(err.String())
	} else {
		println("更新用户信息成功")
	}

	// 获取用户信息
	user, err = base_server_user.GetUserInfo(5, user.UserId)
	if err != nil {
		println(err.String())
	} else {
		fmt.Printf("获取用户信息成功：[%v]\n", *user)
	}

	// 绑定邮箱
	if err = base_server_user.BindEmail(user.OrgId, user.UserId, "462039092@qq.com", ""); err != nil {
		println(err.String())
	} else {
		fmt.Printf("绑定邮箱成功\n")
	}

	base_server_user.StoreValAtomic(1, 1, "1", "2")

	base_server_user.DelStoreVal(1, 1, 5)

	vals, _ := base_server_user.GetStoreVal(1, 1, "1")
	fmt.Printf("%v\n", vals)
}

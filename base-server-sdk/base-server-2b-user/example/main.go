package main

import (
	"fmt"
	"github.com/becent/golang-common"
	"github.com/becent/golang-common/base-server-sdk"
	"github.com/becent/golang-common/base-server-sdk/base-server-2b-user"
	"time"
)

func main() {
	base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			TbUserServerHost: "http://127.0.0.1:8082",
		},
		GRpcOnly: false,
	})
	defer base_server_sdk.ReleaseBaseServerSdk()

	admin, ex := base_server_2b_user.RegisterAdmin(&base_server_2b_user.Admin2B{
		OrgId:         100,
		Account:       common.Md5Encode(time.Now().String()),
		LoginPwd:      "123456",
		ParentAdminId: 10000075,
		NickName:      "sdk_admin",
	}, "", nil)
	if ex != nil {
		println(ex.String())
		return
	}
	fmt.Printf("%v\n", *admin)

	user, ex := base_server_2b_user.RegisterUser(&base_server_2b_user.User2B{
		OrgId:      100,
		Account:    common.Md5Encode(time.Now().String()),
		LoginPwd:   "123456",
		NickName:   "sdk_user",
		InviteCode: admin.InviteCode,
	}, "", nil)
	if ex != nil {
		println(ex.String())
		return
	}
	fmt.Printf("%v\n", *user)

	adminInfo, ex := base_server_2b_user.GetAdminInfo(100, admin.AdminId, 1)
	if ex != nil {
		println(ex.String())
		return
	}
	fmt.Printf("%v\n", *adminInfo)

	userInfo, ex := base_server_2b_user.GetUserInfo(100, user.UserId, 1)
	if ex != nil {
		println(ex.String())
		return
	}
	fmt.Printf("%v\n", *userInfo)

	adminInfos, ex := base_server_2b_user.GetAdminsInfo(100, []int64{admin.AdminId, 10000070}, 1)
	if ex != nil {
		println(ex.String())
		return
	}
	for _, info := range adminInfos {
		fmt.Printf("%v\n", *info)
	}

	userInfos, ex := base_server_2b_user.GetUsersInfo(100, []int64{user.UserId, 10000082}, 1)
	if ex != nil {
		println(ex.String())
		return
	}
	for _, info := range userInfos {
		fmt.Printf("%v\n", *info)
	}

	userList, ex := base_server_2b_user.FindUserList(100, 0, "", "", "", "", "", 0, 0, 0, 0, 0, 0, 1)
	if ex != nil {
		println(ex.String())
		return
	}
	for _, u := range userList {
		fmt.Printf("%v\n", *u)
	}

	adminList, ex := base_server_2b_user.FindAdminList(100, 0, "", "", "", "", "", 0, 0, 0, 0, 0, 0, 0, 0, 0)
	if ex != nil {
		println(ex.String())
		return
	}
	for _, a := range adminList {
		fmt.Printf("%v\n", *a)
	}

	ex = base_server_2b_user.RelationTransfer(100, admin.AdminId, 10000070, 2, 0)
	if ex != nil {
		println(ex.String())
		return
	}
}

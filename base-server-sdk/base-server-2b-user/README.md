# base_server_2b_user 接口说明文档

## 初始化base_server_sdk
```go
base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
    AppId:           "10000",
    AppSecretKey:    "hiojklsankldlksdnlsdasd",
    RequestTimeout:  5 * time.Second,
    IdleConnTimeout: 10 * time.Minute,
    Hosts: base_server_sdk.Hosts{
        TbUserServerHost: "http://127.0.0.1:8081",
    },
    GRpcOnly: false,
})

// ....

defer base_server_sdk.ReleaseBaseServerSdk()
```

## 相关model

```go
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

type base_server_sdk.Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

```

## 接口说明


### 注册用户

func RegisterUser(user *User2B, code string, currencyTypes []string) (*User2B, *base_server_sdk.Error)

```text
1. orgId必须大于0
2. account, phone, email至少要有一个有值
3. account有值时,password必须有值
4. code有值时,将会校验短信或者邮件验证码
5. currencyTypes 有值的话,会顺便创建相关的account账户
6. inviteCode必须有值
7. 其他字段非必填
```

### 获取用户信息

func GetUserInfo(orgId int, userId int64, baseInfo int) (*User2B, *base_server_sdk.Error)

```text
baseInfo 为1，代表同步查询基础信息
```

### 批量查询用户信息

func GetUsersInfo(orgId int, userIds []int64, baseInfo int) (map[int64]*User2B, *base_server_sdk.Error)

```text
baseInfo 为1，代表同步查询基础信息
```

### 后台查询用户列表

func FindUserList(orgId int, userId int64, nickName, countryCode, phone, email, account string,
	parentAdminId, affiliatedAdminId, beginTime, endTime int64, page, pageSize, baseInfo int) ([]*User2B, *base_server_sdk.Error)
	
```text
1. orgId必填
2. userId有值则查询指定的userId用户
3. nickName模糊查找
4. countryCode和phone搭配使用，按手机查找
5. email按邮箱查找
6. account按账号查找
7. parentAdminId按父级会员查找，相当于查直属
8. affiliatedAdminId按归属查询，相当于查询该会员下属的用户
9. baseInfo为1则同步查询基础信息
```

### 注册会员

func RegisterAdmin(admin *Admin2B, code string, currencyTypes []string) (*Admin2B, *base_server_sdk.Error)

```text
```text
1. orgId必须大于0
2. account, phone, email至少要有一个有值
3. account有值时,password必须有值
4. code有值时,将会校验短信或者邮件验证码
5. currencyTypes 有值的话,会顺便创建相关的account账户
6. parentAdminId必填
7. 其他字段非必填
```

### 获取会员信息

func GetAdminInfo(orgId int, adminId int64, baseInfo int) (*Admin2B, *base_server_sdk.Error)

```text
baseInfo为1则同步查询基础信息
```

### 批量获取会员信息

func GetAdminsInfo(orgId int, adminIds []int64, baseInfo int) (map[int64]*Admin2B, *base_server_sdk.Error)

```text
baseInfo为1则同步查询基础信息
```

### 后台查询会员列表

func FindAdminList(orgId int, adminId int64, nickName, countryCode, phone, email, account string,
	roleId, isSub int, parentAdminId, affiliatedAdminId, beginTime, endTime int64, page, pageSize, baseInfo int) ([]*Admin2B, *base_server_sdk.Error)
	
```text
1. orgId必填
2. userId有值则查询指定的userId用户
3. nickName模糊查找
4. countryCode和phone搭配使用，按手机查找
5. email按邮箱查找
6. account按账号查找
7. roleId按角色查找
8. isSub按是否子账号查找
9. parentAdminId按父级会员查找，相当于查直属
10. affiliatedAdminId按归属查询，相当于查询该会员下属的用户
11. baseInfo为1则同步查询基础信息
```

### 转移从属关系

func RelationTransfer(orgId int, fromId, toId int64, ttype int, delayTime int64) *base_server_sdk.Error

```text
1. fromId代表待转移的用户或者会员id
2. toId代表接收转移的会员id
3. ttype为1代表转移用户、2代表转移会员
4. delayTime代表延迟执行时间，单位为秒
```
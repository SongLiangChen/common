# base_server_user 接口说明文档

## 初始化base_server_sdk
```go
base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
    AppId:           "10000",
    AppSecretKey:    "hiojklsankldlksdnlsdasd",
    RequestTimeout:  5 * time.Second,
    IdleConnTimeout: 10 * time.Minute,
    Hosts: base_server_sdk.Hosts{
        UserServerHost: "http://127.0.0.1:8081",
    },
    GRpcOnly: false,
})

// ....

defer base_server_sdk.ReleaseBaseServerSdk()
```

## 相关model
```go
type User struct {
	UserId     int64  `json:"userId"`
	OrgId      int    `json:"orgId"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Account    string `json:"account"`
	LoginPwd   string `json:"loginPwd"`
	TransPwd   string `json:"transPwd"`
	NickName   string `json:"nickName"`
	Avatar     string `json:"avatar"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	IdCard     string `json:"idCard"`
	Sex        int    `json:"sex"`
	BirthDay   string `json:"birthDay"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"createTime"`
	Ext        string `json:"ext"`
}

type base_server_sdk.Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
```

## 相关错误码
```go
1000 服务繁忙
1001 参数错误
1002 用户已注册
1003 用户不存在
1004 用户状态异常
1005 密码错误
1006 请勿重复绑定
1007 请先设置密码
1008 短信验证码错误
1009 邮件验证码错误
1010 身份认证失败
```


## 接口说明

- 注册用户

func Register(user *User, code string, currencyTypes []string) (*User, *base_server_sdk.Error)

```go
1. orgId必须大于0
2. account, phone, email至少要有一个有值
3. account有值时,password必须有值
4. code有值时,将会校验短信或者邮件验证码
5. currencyTypes 有值的话,会顺便创建相关的account账户
6. 其他字段非必填

异常返回:
1000 服务繁忙
1001 参数异常
1002 账户已经注册
```

- 账号密码登录

func LoginByAccount(orgId int, account string, password string) (*User, *base_server_sdk.Error)

```go
1. password有值时,会进行登录密码校验,为空则直接登录成功

异常返回:
1003 账号不存在
1004 用户已被冻结
1005 密码错误
```

- 手机号登录

func LoginByPhone(orgId int, phone string, code string, password string) (*User, *base_server_sdk.Error)

```go
1. code有值时,会进行短信校验,为空则略过
2. password有值时,会进行登录密码校验,为空则略过

异常返回:
1003 账号不存在
1004 用户已被冻结
1005 密码错误
1008 短信验证码错误
```

- 邮箱登录

func LoginByEmail(orgId int, email string, code string, password string) (*User, *base_server_sdk.Error)

```go
1. code有值时,会进行邮件校验,为空则略过
2. password有值时,会进行登录密码校验,为空则略过

异常返回:
1003 账号不存在
1004 用户已被冻结
1005 密码错误
1009 邮件验证码错误
```

- 获取用户信息

func GetUserInfo(orgId int, userId int64) (*User, *base_server_sdk.Error)

```go
异常返回:
1003 用户不存在
```

- 批量获取用户信息
func GetUsersInfo(orgId int, userIds []int64) (map[int64]*User, []int64, *base_server_sdk.Error)

```go
返回值:
map[int64]*User 查询到的用户
[]int64 未找到的用户
```

- 通过手机号找回登录密码

func GetBackLoginPwdByPhone(orgId int, phone, code, password string) *base_server_sdk.Error

```go
1. code有值的话会进行短信校验,为空则忽略
2. password是要设置的新密码

异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
1005 密码错误
1008 短信验证码错误
```

- 通过邮箱找回登录密码

func GetBackLoginPwdByEmail(orgId int, email, code, password string) *base_server_sdk.Error

```go
1. code有值的话会进行邮箱校验,为空则忽略
2. password是要设置的新密码

异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
1005 密码错误
1009 邮箱验证码错误
```

- 通过手机找回交易密码

func GetBackTransPwdByPhone(orgId int, phone, code, password string) *base_server_sdk.Error

```go
1. code有值的话会进行短信校验,为空则忽略
2. password是要设置的新密码

异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
1005 密码错误
1008 短信验证码错误
```

- 通过邮箱找回交易密码

func GetBackTransPwdByEmail(orgId int, email, code, password string) *base_server_sdk.Error

```go
1. code有值的话会进行邮箱校验,为空则忽略
2. password是要设置的新密码

异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
1005 密码错误
1009 邮箱验证码错误
```

- 更新登录密码

func UpdateLoginPwd(orgId int, userId int64, oldPassword, newPassword string) *base_server_sdk.Error

```go
1. oldPassword有值则校验旧密码,为空则略过
2. newPassword必须有值

异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
1005 密码错误
```

- 更新交易密码

func UpdateTransPwd(orgId int, userId int64, oldPassword, newPassword string) *base_server_sdk.Error

```go
1. oldPassword有值则校验旧密码,为空则略过
2. newPassword必须有值

异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
1005 密码错误
```

- 实名认证

func AuthRealName(orgId int, userId int64, firstName, lastName, idCard string) *base_server_sdk.Error

```go
1. idCard有值的话,会进行身份证校验,为空则忽略
2. 真实姓名为firstName+lastName

异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
1010 身份证认证失败
```

- 更新用户信息

func UpdateUserInfo(orgId int, userId int64, info UserFields) *base_server_sdk.Error

```go
type Sex int

const (
	_ Sex = iota
	Boy
	Girl
)

type UserFields map[string]string

func (u UserFields) SetNickName(nickName string)

func (u UserFields) SetAvatar(avatar string)

func (u UserFields) SetSex(sex Sex)

func (u UserFields) SetBirthDay(birthDay string)

func (u UserFields) SetExt(ext string)
```

```go
可以更新的字段为:
1. nickName
2. avatar
3. sex
4. birthDay
5. ext

用例:
info := make(base_server_user.UserFields)
info.SetNickName("jak")
info.SetSex(base_server_user.Boy)

base_server_user.UpdateUserInfo(1, 1000, info)


异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
```

- 绑定登录账号

func BindAccount(orgId int, userId int64, account, password string) *base_server_sdk.Error

```go
异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
1006 重复绑定
```

- 绑定手机

func BindPhone(orgId int, userId int64, phone, code string) *base_server_sdk.Error

```go
异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
1006 重复绑定
```

- 绑定邮箱

func BindEmail(orgId int, userId int64, email, code string) *base_server_sdk.Error

```go
异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
1006 重复绑定
```

- 校验支付密码

func AuthTransPwd(orgId int, userId int64, password string) *base_server_sdk.Error

```go
异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
1005 密码错误
```

- 修改用户状态

func UpdateUserStatus(orgId int, userId int64, status Status) *base_server_sdk.Error

```go
异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
```

- 解绑手机

func UnBindPhone(orgId int, userId int64, code string) *base_server_sdk.Error

```go
1. code有值的话,会校验短信验证码

异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
1011 手机未绑定
1013 解绑手机后将失去所有登录方式
```

- 解绑邮箱

func UnBindEmail(orgId int, userId int64, code string) *base_server_sdk.Error

```go
1. code有值的话,会校验邮箱验证码

异常返回:
1000 服务繁忙
1001 参数异常
1003 用户不存在
1012 邮箱未绑定
1014 解绑邮箱后将失去所有登录方式
```

- 存储用户业务信息

func StoreValAtomic(orgId int, userId int64, key, val string) *base_server_sdk.Error

```go
原子操作,存储用户一些信息

异常返回:
1000 服务繁忙
1001 参数异常
```

- 获取用户的业务信息

func GetStoreVal(orgId int, userId int64, key string) (map[int64]string, *base_server_sdk.Error)

- 删除用户的业务信息

func DelStoreVal(orgId int, userId int64, id int64) *base_server_sdk.Error
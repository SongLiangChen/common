## base_server_sdk

## 初始化base_server_sdk
```go
base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
    OrgId:           1,
    AppId:           "10000",
    AppSecretKey:    "hiojklsankldlksdnlsdasd",
    RequestTimeout:  5 * time.Second,
    IdleConnTimeout: 10 * time.Minute,
    Hosts: base_server_sdk.Hosts{
        UserServerHost: "http://127.0.0.1:8081",
        AccountServerHost: "http://127.0.0.1:8082",
        StatisticServerHost: "http://127.0.0.1:8083",
        OctopusServerHost: "http://127.0.0.1:8084",
    },
    GRpcOnly: false,
})

// ...

defer base_server_sdk.ReleaseBaseServerSdk()
```

## 测试环境
- http：http://127.0.0.1:5051
- grpc：127.0.0.1:15051

### 相关类型
```go
//Error
type base_server_sdk.Error struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

//业务类型
type BusinessId int

//
type GenerateGaRes struct {
    QrCode string `json:"qrCode"`
    SecretKey string `json:"secretKey"`
}

```

## 业务码常量
```go
const (
    BusinessRegister        BusinessId = 1000 // 注册
    BusinessLogin           BusinessId = 1001 // 登录
    BusinessUpdatePhone     BusinessId = 1002 // 更新手机
    BusinessBindPhone       BusinessId = 1003 // 绑定手机
    BusinessUpdateEmail     BusinessId = 1004 // 更新邮箱
    BusinessBindEmail       BusinessId = 1005 // 绑定邮箱
    BusinessGetBackLoginPwd BusinessId = 1006 // 找回密码
)
```

## 相关错误码
```go
1000    服务繁忙
1001    参数错误
1002    未找到邮件模板
1003    验证码发送太频繁
1004    发送邮件失败
1005    验证码检验失败
1006    无最新校验记录
1007    未找到短信模板
1008    发送短信失败
1009    实名认证失败
1010    生成GA密钥失败
1011    检验GA失败
1012    验证码初始化失败
1013    验证码校验失败
```


## 邮件服务

**发送邮件验证码**
```go
func SendEmailCode(orgId int, businessId BusinessId, email, lang string) *base_server_sdk.Error
```
- 示例
```go
err := base_server_octopus.SendEmailCode(5, base_server_octopus.BusinessLogin, "xxx@qq.com", "zh")
```
- 异常返回
```go
1001 参数错误
1002 未找到邮件模板
1003 验证码发送太频繁
1004 发送邮件失败
```

**校验邮件验证码**
```go
func VerifyEmailCode(orgId int, businessId int, email, code string) (bool, *base_server_sdk.Error)
```
- 示例
```go
ret, err := base_server_octopus.VerifyEmailCode(5,base_server_octopus.BusinessLogin, "xxx@qq.com", "1235")
```

**校验上次邮件验证码是否通过**
```go
func CheckLastEmailVerifyResult(orgId int, businessId BusinessId, email string) (bool, *base_server_sdk.Error)
```
- 示例
```go
ret, err := base_server_octopus.CheckLastEmailVerifyResult(5,base_server_octopus.BusinessLogin,"email")
```


## 短信服务

**发送短信验证码**
```go
//countryCode 国家码, 国内默认86
func SendSimCode(orgId int, businessId BusinessId, countryCode, phone, lang string) *base_server_sdk.Error
```
- 示例
```go
err := base_server_octopus.SendSimCode(5, base_server_octopus.BusinessLogin, "86", "130xxxxxxx", "zh")
```
- 异常返回
```go
1001 参数错误
1007 未找到短信模板
1003 验证码发送太频繁
1008 发送短信失败
```

**校验短信验证码**
```go
func VerifySimCode(orgId int, businessId BusinessId, countryCode phone, code string) (bool, *base_server_sdk.Error)
```
- 示例
```go
ret, err := base_server_octopus.VerifySimCode(5, base_server_octopus.BusinessLogin, "86", "130xxxx1234", "54321")
```

**校验上次短信验证码是否通过**
```go
func CheckLastSimVerifyResult(orgId int, businessId int, countryCode, phone, code string) (bool, *base_server_sdk.Error)
```
- 示例
```go
ret, err := base_server_octopus.CheckLastSimVerifyResult(5,base_server_octopus.BusinessLogin, "86", "130xxxxxxxx", "1235")
```

## 实名验证

**实名验证**
```go
func AuthRealName(orgId int, name string, cardNo string) (bool, *base_server_sdk.Error)
```
- 示例
```go
res, err := base_server_octopus.AuthRealName(5, "张三", "010203201909201234")
```


## 谷歌验证

**谷歌验证初始化（获取密钥）**
```go
func GenerateGa(orgId int, businessId BusinessId, account string) (*GenerateGaRes, *base_server_sdk.Error)
```
- 示例
```go
res, err := base_server_octopus.GenerateGa(5, base_server_octopus.BusinessLogin, "130xxxx1234")
```
- 异常返回
```go
1000 参数错误
1010 生成GA密钥失败
```
- 成功返回
```go
{
"qrCode": "...", // 二维码链接
"secretKey": "base64encodestring" //密钥
}
```

**校验code**
```go
func VerifyGa(orgId int, businessId BusinessId, account string, secret string, gaCode string) (bool, *base_server_sdk.Error)
```
- 示例
```go 
ret, err := base_server_octopus.VerifyGa(5, base_server_octopus.BusinessLogin, "130xxxx1234", "secret", "code")
```


## 极验验证服务

**验证码初始化**
```go
func InitGt(orgId int, businessId BusinessId, account string, ip string) (*InitCaptchaRes, *base_server_sdk.Error)
```
- 示例
```go
res, err := base_server_octopus.InitGt(1, base_server_octopus.BusinessLogin, "130xxxx1234", "127.0.0.1")
```
- 异常返回
```go
1001 参数错误
1012 验证码初始化失败
```
- 成功返回
```go
{
"success": 0/1, //标识是否走本地验证
"gt": "极验账户密钥",
"challenge": "验证码唯一id",
"new_captcha": 0/1 //标识是否走本地验证
}
```

**服务端校验验证码
```go
func VerifyGt(orgId int, businessId BusinessId, account string, ip string, challenge, validate, seccode string) (bool, *base_server_sdk.Error)
```
- 示例
```go
ret, err := base_server_octopus.VerifyGt(1, base_server_octopus.BusinessLogin, "130xxxx1234", "ip", "challenge", "validate", "seccode")
```
- 异常返回
```go
1001 参数错误
1013 验证码校验失败
```

## 验证码服务

**验证码初始化**
```go
//length:验证码长度，默认6位(0~9); width:图片宽度，默认240; height:图片高度，默认80; 传0则用默认值
func InitCaptcha(orgId int, businessId BusinessId, length, width, height int) (*InitCaptchaResponse, *base_server_sdk.Error)
```
- 示例
```go
res, err := base_server_octopus.InitCaptcha(1, base_server_octopus.BusinessLogin, 0, 0, 0)
```
- 异常返回
```go
1001 参数错误
1012 验证码初始化失败
```
- 成功返回
```go
{
"success": ture,
"captchaId": "验证码Id",
"image": "base64图片验证码",
}
```

**刷新验证码**
```go
//length:验证码长度，默认6位(0~9); width:图片宽度，默认240; height:图片高度，默认80; 传0则用默认值
func ReloadCaptcha(orgId int, businessId BusinessId, captchaId string, width, height int) (*InitCaptchaResponse, *base_server_sdk.Error)
```
- 示例
```go
res, err := base_server_octopus.ReloadCaptcha(1, base_server_octopus.BusinessLogin, "captchaId", 0, 0)
```
- 异常返回
```go
1001 参数错误
1017 刷新验证码失败
```
- 成功返回
```go
{
"success": ture,
"captchaId": "验证码Id",
"image": "base64图片验证码",
}
```

**服务端校验验证码
```go
func VerifyCaptcha(orgId int, businessId BusinessId, captchaId string, digits string) (bool, *base_server_sdk.Error)
```
- 示例
```go
ret, err := base_server_octopus.VerifyCaptcha(1, base_server_octopus.BusinessLogin, "captchaId", "535096")
```
- 异常返回
```go
1001 参数错误
1013 验证码校验失败
```

**上传文件到s3
```go
func Upload(orgId int, formFile map[string]string) (map[string]string, *base_server_sdk.Error)
```
- 示例
```go
formFile := make(map[string]string)
formFile["file1"] = "path/to/test.log"
formFile["file2"] = "path/to/test.2.log"
res, err := base_server_octopus.Upload(1, formFile)
// 返回
//	result = {
//		"file1": "https://xxx.com/path/to/file1",
//		"file2": "https://xxx.com/path/to/file2"
//	}
```
- 异常返回
```go
900002 参数错误
900003 文件路径错误
```
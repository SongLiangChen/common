# base_server_user 接口说明文档

## 初始化base_server_sdk
```go
base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		AppId:           "10002",
		AppSecretKey:    "12345678910",
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			AiPayServerHost: "https://t-openapi.aipaybox.com"
		},
		GRpcOnly: false,
	})

// ....

defer base_server_sdk.ReleaseBaseServerSdk()
```

## 相关model
```go
//接口版本
const VERSION = "1.0.1"

//货币
const (
	CURRENCY_CC   = "CC"
	CURRENCY_USD  = "USD"
	CURRENCY_USDT = "USDT"
	CURRENCY_EOS  = "EOS"
	CURRENCY_XRP  = "XRP"
	CURRENCY_BTC  = "BTC"
	CURRENCY_BCH  = "BCH"
	CURRENCY_ETH  = "ETH"
)

//账户余额查询返回
type Account struct {
	MchId        string `json:"mch_id"`
	Currency     string `json:"currency"`
	AvailAmount  string `json:"avail_amount"`
	FreezeAmount string `json:"freeze_amount"`
}

//查询可用支付方式返回
type PayMethod struct {
	MethodCode           string `json:"methodCode"`
	MethodName           string `json:"methodName"`
	MinSingleOrderAmount string `json:"minSingleOrderAmount"`
	MaxSingleOrderAmount string `json:"maxSingleOrderAmount"`
}

//支付通道
type PayChannel struct {
	ChannelId            string `json:"channelId"`
	ChannelDesc          string `json:"channelDesc"`
	MinSingleOrderAmount string `json:"minSingleOrderAmount"`
	MaxSingleOrderAmount string `json:"maxSingleOrderAmount"`
}

//支付订单状态
const (
	PAY_STATUS_SUCCESS = "SUCCESS"
	PAY_STATUS_PAYING  = "PAYING"
	PAY_STATUS_FAIL    = "FAIL"

	WITHDRAW_STATUS_SUCCESS = "SUCCESS"
	WITHDRAW_STATUS_PAYING  = "PAYING"
	WITHDRAW_STATUS_FAIL    = "FAIL"
)

//支付订单
type PayOrder struct {
	PayMethod      string `json:"pay_method"`
	PayChannel     string `json:"pay_channel"`
	MchId          string `json:"mch_id"`
	TransactionId  string `json:"transaction_id"`
	OutTradeNo     string `json:"out_trade_no"`
	NonceStr       string `json:"nonce_str"`
	SignType       string `json:"sign_type"`
	Detail         string `json:"detail"`
	Attach         string `json:"attach"`
	SpbillCreateIp string `json:"spbill_create_ip"`
	NotifyUrl      string `json:"notify_url"`
	UserOutFee     string `json:"user_out_fee"`
	UserOutType    string `json:"user_out_type"`
	MchInFee       string `json:"mch_in_fee"`
	MchInType      string `json:"mch_in_type"`
	TimeStart      int64  `json:"time_start"`
	TimeExpire     int64  `json:"time_expire"`
	TimeEnd        int64  `json:"time_end"`
	TradeStatus    string `json:"trade_status"`
	CodeContent    string `json:"code_content"`
	CodePage       string `json:"code_page"`
	Sign           string `json:"sign"`
	Version        string `json:"version"`
}

//提现银行账户信息
type BankAccount struct {
	BankName       string
	BankUserName   string
	BankUserPhone  string
	BankBranchName string
	BankCardNo     string
	BankProvince   string
	BankCity       string
	QrCode         string
	QrCodeImgType  string
}

//代付订单
type WithdrawOrder struct {
	PayMethod      string `json:"pay_method"`
	PayChannel     string `json:"pay_channel"`
	MchId          string `json:"mch_id"`
	TransactionId  string `json:"transaction_id"`
	OutTradeNo     string `json:"out_trade_no"`
	NonceStr       string `json:"nonce_str"`
	SignType       string `json:"sign_type"`
	Detail         string `json:"detail"`
	Attach         string `json:"attach"`
	SpbillCreateIp string `json:"spbill_create_ip"`
	NotifyUrl      string `json:"notify_url"`
	UserOutFee     string `json:"user_out_fee"`
	UserInType     string `json:"user_in_type"`
	UserInFee      string `json:"user_in_fee"`
	MchOutType     string `json:"mch_out_type"`
	TimeStart      string `json:"time_start"`
	TimeExpire     string `json:"time_expire"`
	TimeEnd        string `json:"time_end"`
	TradeStatus    string `json:"trade_status"`
	Sign           string `json:"sign"`
	Version        string `json:"version"`
}
```

## 相关错误码
```go
//参数错误
const  COMMON_PARAMS_ERROR = 1001;
//版本错误
const  COMMON_VERSION_ERROR = 1002;
//未知错误
const  COMMON_UNKNOW_ERROR = 9999;
//非法操作
const  COMMON_ILLEGAL_OPERATE = 8888;
//请求过于频繁
const  COMMON_REQ_TOO_FREQUENT = 7777;
//请求重复
const  COMMON_REQ_REPEAT = 5555;
//签名错误
const  COMMON_SIGN_ERROR = 6666;
//数据库操作异常
const COMMON_DATABASES_EXCEPTION = 4444;
//支付商户不存在
const PAY_MERCHANT_NOT_EXIST = 3001;
//支付方式错误
const PAY_METHOD_ERROR = 3002;
//支付下单失败
const PAY_SEND_ORDER_FAIL = 3003;
//支付查询失败
const QUERY_ORDER_FAIL = 3004;
//回调订单失败
const CALLBACK_ORDER_FAIL = 3005;
//订单状态异常
const PAY_ORDER_STATUS_ERROR = 3006;
//没有可用金额
const PAY_NO_UNION_AMOUNT = 3007;
//支付方式不存在
const PAY_METHOD_NOT_EXIST = 4001;
//货币类型错误
const PAY_FEE_TYPE_ERROR = 4002;
//签名类型错误
const PAY_SIGN_TYPE_ERROR = 4003;
//通道不存在
const PAY_CHANNEL_NOT_EXIST = 4004;
//回调地址错误
const PAY_NOTIFY_URL_ERROR = 4005;
//订单号错误
const PAY_ORDER_NO_ERROR = 4006;
//金额错误
const PAY_ORDER_AMOUNT_ERROR = 4007;
//到账金额错误
const PAY_ARRIVE_FEE_TYPE_ERROR = 4008;
//ip地址错误
const PAY_IP_ADDRESS_ERROR = 4009;
//随机字符串错误
const PAY_NONCE_STR_ERROR = 4010;
//银行卡参数错误
const PAY_BANK_PARAMS_ERROR = 4011;
//钱包地址错误
const PAY_WALLET_PARAMS_ERROR = 4012;
//支付路由错误
const PAY_ROUTE_ERROR = 4101;
//商户限额错误
const PAY_MERCHANT_LIMIT_ERROR = 4102;
//支付方式限额错误
const PAY_METHOD_LIMIT_ERROR = 4103;
//通道限额错误
const PAY_CHANNEL_LIMIT_ERROR = 4104;
//通道商户限额错误
const PAY_CHANNEL_MERCHANT_LIMIT_ERROR = 4105;
//通道商户apicode错误
const PAY_CHANNEL_MERCHANT_API_CODE_ERROR = 4106;
//通道商户接口调用错误
const PAY_CHANNEL_MERCHANT_API_CALL_ERROR = 4107;

//商户费率错误
const PAY_MCH_RATE_ERROR = 4108;
//通道费率错误
const PAY_CHAN_RATE_ERROR = 4109;
//通道费率错误
const PAY_CHAN_ACCOUNT_ERROR = 4110;
//订单号重复
const PAY_ORDER_REPEAT_ERROR = 4111;
//二维码错误
const PAY_QR_CODE_ERROR = 4112;

//订单不存在
const PAY_ORDER_NOT_EXIST = 4500;
//out_trade_no与transaction_id不能同时为空
const PAY_ORDER_NO_ORDER_NO = 4501;
//订单超时
const PAY_ORDER_EXPIRE = 4502;
//创建订单异常
const WDW_CREATE_ORDER_EXCEPTION = 5001;
//余额不足
const ACCOUNT_NO_BALANCE = 6001;
```


## 接口说明

- 查询账户余额

func QueryBalance(mchId string, currency string, signKey string) (*Account, *base_server_sdk.Error)

```go
1.mchId 商户id
2.currency 货币,使用CURRENCY常量
3.signKey 签名密钥
```
- 查询可用支付方式

func SelectPayMethods(mchId string, userOutType string, userOutFee string, signKey string) ([]PayMethod, *base_server_sdk.Error)

```go
1.mchId 商户id
2.userOutType 用户支出货币,使用CURRENCY常量
3.userOutFee 用户支出金额,元为单位
4.signKey 签名密钥
```

- 查询可用支付通道

func SelectPayChannels(mchId string, payMethod string, userOutFee string, userOutType string, mchInType string, signKey string) ([]PayChannel, *base_server_sdk.Error)

```go
1.mchId 商户id
2.payMethod 支付方式编码
3.userOutType 用户支出货币,使用CURRENCY常量
4.userOutFee 用户支出金额,元为单位
5.mchInType 商户入账货币,使用CURRENCY常量
6.signKey 签名密钥
```

- 提交支付订单

func SubmitPayOrder(mchId string, payMethod string, payChannel string, outTradeNo string, userOutFee string, userOutType string,mchInType string, detail string, attach string, notifyUrl string, signKey string) (*PayOrder, *base_server_sdk.Error)

```go
1.mchId 商户id
2.payMethod 支付方式编码
3.payChannel 支付通道ID
4.outTradeNo 商户订单号
5.userOutType 用户支出货币,使用CURRENCY常量
6.userOutFee 用户支出金额,元为单位
7.mchInType 商户入账货币,使用CURRENCY常量
8.detail 订单详情
9.attach 附带字段,回调时返回
10.notifyUrl 回调地址
11.signKey 签名密钥
```

- 查询支付订单

func QueryPayOrder(mchId string, outTradeNo string, transactionId string, signKey string) (*PayOrder, *base_server_sdk.Error)

```go
1.mchId 商户id
2.outTradeNo 商户订单号
3.transactionId 平台订单号
4.signKey 签名密钥
```

- 生成聚合支付链接

func GenerateUnionPayUrl(mchId string, currency string, userId string, reqTime string,amount string, notifyUrl string, redirectUrl string, signKey string) (string, *base_server_sdk.Error)

```go
1.mchId 商户id
2.currency 货币,使用CURRENCY常量
3.userId 用户id，回调时返回
4.reqTime 请求时间，unix时间戳
5.amount 下单金额,元为单位
6.notifyUrl 回调地址
7.redirectUrl 页面返回地址
8.signKey 签名密钥
```

- 查询代付方式

func SelectWithdrawMethods(mchId string, userInType string, signKey string) ([]PayMethod, *base_server_sdk.Error)

```go
1.mchId 商户id
2.userInType 用户入账货币,使用CURRENCY常量
3.signKey 签名密钥
```

- 查询代付通道

func SelectWithdrawChannels(mchId string, payMethod string, userInFee string, userInType string, mchOutType string, signKey string) ([]PayChannel, *base_server_sdk.Error)

```go
1.mchId 商户id
2.payMethod 支付方式编码
3.userInFee 用户入账金额,单位元
4.userInType 用户入账货币,使用CURRENCY常量
5.mchOutType 商户出账货币,使用CURRENCY常量
6.signKey 签名密钥
```

- 提交代付订单

func SubmitWithdrawOrder(mchId string,payMethod string,payChannel string,outTradeNo string,userInFee string,userInType string,mchOutType string,detail string,attach string,
	notifyUrl string,signKey string, bankAccount *BankAccount) (*WithdrawOrder, *base_server_sdk.Error)

```go
1.mchId 商户id
2.payMethod 支付方式编码
3.payChannel 代付通道ID
4.outTradeNo 商户订单号
5.userInFee 用户入账金额
6.userInType 用户入账货币,使用CURRENCY常量
7.mchOutType 商户出账货币,使用CURRENCY常量
8.detail 订单详情
9.attach 附带字段,回调时返回
10.notifyUrl 回调地址
11.bankAccount 银行账户，参考结构BankAccount
12.signKey 签名密钥
```


- 查询代付订单

func QueryWithdrawOrder(mchId string, outTradeNo string, transactionId string, signKey string) (*WithdrawOrder, *base_server_sdk.Error)

```go
1.mchId 商户id
2.outTradeNo 商户订单号
3.transactionId 平台订单号
4.signKey 签名密钥
```
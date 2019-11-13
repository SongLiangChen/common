## base_server_sdk

## 初始化base_server_sdk
```go
base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
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

defer base_server_sdk.ReleaseBaseServerSdk()

// ...

```
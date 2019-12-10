package base_server_sdk

import (
	"context"
	"github.com/SongLiangChen/grpc_pool"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Hosts struct {
	UserServerHost      string
	AccountServerHost   string
	StatisticServerHost string
	OctopusServerHost   string
	AiPayServerHost     string
	InteractServerHost  string
	GoodsServerHost     string
	TbUserServerHost    string
	// TODO add server host here
}

type BaseServerSdkClient struct {
	gRpcMapPool *grpc_pool.MapPool
	httpClient  *http.Client

	OrgId int

	appId        string
	appSecretKey string

	requestTimeout time.Duration

	cp sync.Pool

	gRpcOnly bool

	Hosts Hosts
}

func DialFunc(addr string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return grpc.DialContext(ctx, addr, grpc.WithBlock(), grpc.WithInsecure())
}

type Config struct {
	// AppId use to generate signature, see makeSignature() for more detail
	// Ignore if GRpcOnly is true
	AppId string
	// AppSecretKey work like AppId
	// Ignore if GRpcOnly is true
	AppSecretKey string
	// RequestTimeout represent both request timeout and response timeout
	RequestTimeout time.Duration
	// IdleConnTimeout is the maximum amount of time an idle
	// connection will remain idle before closing itself
	IdleConnTimeout time.Duration
	// Hosts stores host of base server
	// the host should contain schema such like 'http://' or 'https://' but not contain last '/'
	// for example: 'https://user.baseserver.com'
	Hosts Hosts
	// Will do all request in gRpc way if GRpcOnly is true
	GRpcOnly bool
}

var Instance *BaseServerSdkClient

func InitBaseServerSdk(c *Config) {
	Instance = &BaseServerSdkClient{
		gRpcMapPool: grpc_pool.NewMapPool(DialFunc, 0, c.IdleConnTimeout),

		httpClient: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   c.RequestTimeout,
					KeepAlive: 0,
					DualStack: true,
				}).DialContext,
				IdleConnTimeout:       c.IdleConnTimeout,
				ResponseHeaderTimeout: c.RequestTimeout,
			},
		},

		requestTimeout: c.RequestTimeout,

		appId:        c.AppId,
		appSecretKey: c.AppSecretKey,

		gRpcOnly: c.GRpcOnly,

		Hosts: c.Hosts,
	}

	Instance.cp.New = func() interface{} {
		return &strings.Builder{}
	}
}

func ReleaseBaseServerSdk() {
	Instance.gRpcMapPool.ReleaseAllPool()
}

package main

import (
	"flag"
	"github.com/becent/golang-common"
	"github.com/becent/golang-common/newProject/config"
	"github.com/becent/golang-common/newProject/data"
	"github.com/becent/golang-common/newProject/exception"
	"github.com/becent/golang-common/newProject/gRpcHandler"
	"github.com/becent/golang-common/newProject/handler"
	"github.com/becent/golang-common/newProject/model"
	"github.com/becent/golang-common/newProject/router"
	"github.com/becent/golang-common/newProject/service"
	"os"
	"strings"
	"time"
)

var (
	projectName = flag.String("projectName", "helloWorld", "project name")
	mysqlInfo   = flag.String("mysqlInfo", "", "mysql connect info, format like \"root:root@tcp(127.0.0.1:3306)/user\"")
)

type Func func(string) error

func main() {
	flag.Parse()

	if *mysqlInfo != "" {
		if err := common.AddDB("database", *mysqlInfo, 5, 2, time.Hour); err != nil {
			println(err.Error())
			return
		}
	}

	// 创建项目文件
	if err := os.Mkdir(*projectName, 755); err != nil {
		println(err.Error())
		return
	}

	// 创建restart.sh文件
	restart, err := os.OpenFile(*projectName+"/restart.sh", os.O_CREATE|os.O_RDWR, 755)
	if err != nil {
		println(err.Error())
		return
	}
	defer restart.Close()
	if _, err := restart.WriteString(strings.Replace(restart_temple, "{{projectName}}", *projectName, -1)); err != nil {
		println(err.Error())
		return
	}

	// 创建build.bat文件
	build, err := os.OpenFile(*projectName+"/build.bat", os.O_CREATE|os.O_RDWR, 755)
	if err != nil {
		println(err.Error())
		return
	}
	defer build.Close()
	if _, err := build.WriteString(strings.Replace(build_temple, "{{projectName}}", *projectName, -1)); err != nil {
		println(err.Error())
		return
	}

	// 创建main文件
	file, err := os.OpenFile(*projectName+"/main.go", os.O_CREATE|os.O_RDWR, 755)
	if err != nil {
		println(err.Error())
		return
	}
	defer file.Close()
	if _, err := file.WriteString(strings.Replace(main_temple, "{{projectName}}", *projectName, -1)); err != nil {
		println(err.Error())
		return
	}

	funcs := make([]Func, 0)
	funcs = append(funcs, []Func{
		config.G_config,           // 创建config文件夹
		data.G_data,               // 创建data目录
		exception.G_exception,     // 创建异常目录
		handler.G_handler,         // 创建handler
		gRpcHandler.G_gRpcHandler, // 创建gRpcHandler
		model.G_model,             // 创建model
		router.G_router,           // 创建router
		service.G_service,         // 创建service
	}...)

	for _, f := range funcs {
		if err := f(*projectName); err != nil {
			println(err.Error())
			return
		}
	}

}

var build_temple = `cd ..
cd ..
set GOPATH=%cd%
cd src
cd {{projectName}}
go build -tags="jsoniter" -o main.exe main.go
`

var restart_temple = `#!/bin/bash

PRONAME={{projectName}}

BIN=/data/apps/$PRONAME/$PRONAME
STDLOG=/data/apps/$PRONAME/output.log

if [ $RUNMODE = "pre" ] ; then
        BIN=/data/apps/pre-$PRONAME/$PRONAME
        STDLOG=/data/apps/$PRONAME/output.log
fi

chmod u+x $BIN

ID=$(/usr/sbin/pidof "$BIN")
if [ "$ID" ] ; then
        echo "kill -SIGINT $ID"
        kill -2 $ID
fi

while :
do
        ID=$(/usr/sbin/pidof "$BIN")
        if [ "$ID" ] ; then
                echo "$PRONAME still running...wait"
                sleep 0.1
        else
                echo "$PRONAME service was not started"
                echo "Starting service..."
                
                if [ $RUNMODE = "online" ] ; then
                        nohup $BIN > /dev/null 2>&1 &
                else
                        nohup $BIN > $STDLOG 2>&1 &
                fi
                break
        fi
done

`

var main_temple = `package main

import (
	"context"
	"github.com/becent/golang-common"
	"github.com/becent/golang-common/base-server-sdk"
	"{{projectName}}/config"
	"{{projectName}}/router"
	"github.com/judwhite/go-svc/svc"
	"google.golang.org/grpc"
	"math"
	"net/http"
	"time"
)

type Service struct {
	gRpcSvr    *grpc.Server
	httpServer *http.Server
}

func (s *Service) Init(env svc.Environment) error {
	config.InitConfig()
	println("RunMode:", config.CURMODE)
	for key, val := range config.GetSection("system") {
		println(key, val)
	}

	// init log
	common.ConfigLogger(
		config.CURMODE,
		config.GetConfig("system", "app_name"),
		config.GetConfig("logs", "dir"),
		config.GetConfig("logs", "file_name"),
		config.GetConfigInt("logs", "keep_days"),
		config.GetConfigInt("logs", "rate_hours"),
	)
	println("logger init success")

	// init mysql
	dbInfo := config.GetSection("dbInfo")
	for name, info := range dbInfo {
		if err := common.AddDB(
			name,
			info,
			config.GetConfigInt("mysql", "maxConn"),
			config.GetConfigInt("mysql", "idleConn"),
			time.Hour*time.Duration(config.GetConfigInt("mysql", "maxLeftTime"))); err != nil {
			return err
		}
	}
	println("mysql init success")

	// init redis
	if err := common.AddRedisInstance(
		"",
		config.GetConfig("redis", "addr"),
		config.GetConfig("redis", "port"),
		config.GetConfig("redis", "password"),
		config.GetConfigInt("redis", "db_num")); err != nil {
		return err
	}
	println("redis init success")

	// init base-server-sdk
	base_server_sdk.InitBaseServerSdk(&base_server_sdk.Config{
		AppId:           config.GetConfig("system", "app_id"),
		AppSecretKey:    config.GetConfig("system", "app_secret_key"),
		RequestTimeout:  5 * time.Second,
		IdleConnTimeout: 10 * time.Minute,
		Hosts: base_server_sdk.Hosts{
			AccountServerHost:   config.GetConfig("baseServiceSdk", "base_service_account_host"),
			OctopusServerHost:   config.GetConfig("baseServiceSdk", "base_service_octopus_host"),
			UserServerHost:      config.GetConfig("baseServiceSdk", "base_service_user_host"),
			StatisticServerHost: config.GetConfig("baseServiceSdk", "base_service_statistic_host"),
			InteractServerHost:  config.GetConfig("baseServiceSdk", "base_service_interact_host"),
		},
		GRpcOnly: false,
	})

	return nil
}

func (s *Service) Start() error {
	// launch http server here...
	s.httpServer = &http.Server{
		Addr:        ":" + config.GetConfig("system", "http_listen_port"),
		Handler:     router.NewGinEngine(),
		ReadTimeout: time.Second * 5,
	}

	go func() {
		// Service connections
		if err := s.httpServer.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()
	println("http service start success")

	// launch gRpc service here
	var err error
	s.gRpcSvr, err = router.NewGRpcEngine().Run(
		":"+config.GetConfig("system", "gRpc_listen_port"),
		grpc.MaxRecvMsgSize(math.MaxInt32))
	if err != nil {
		return err
	}
	println("gRpc service start success")

	return nil
}

func (s *Service) Stop() error {
	// stop gRpc server
	s.gRpcSvr.GracefulStop()
	println("grpc server graceful stop")

	// stop http server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		println("Server Shutdown:", err)
	}
	println("http server graceful stop")

	// release source here
	base_server_sdk.ReleaseBaseServerSdk()
	common.ReleaseMysqlDBPool()
	common.ReleaseRedisPool()

	return nil
}

func main() {
	if err := svc.Run(&Service{}); err != nil {
		println(err.Error())
	}
}

`

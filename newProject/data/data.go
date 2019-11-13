package data

import (
	"os"
	"strings"
)

func G_data(projectName string) error {
	if err := os.Mkdir(projectName+"/data", 755); err != nil {
		return err
	}

	if err := os.Mkdir(projectName+"/data/ini", 755); err != nil {
		return err
	}

	if err := os.Mkdir(projectName+"/data/json", 755); err != nil {
		return err
	}

	// online
	file, err := os.OpenFile(projectName+"/data/ini/online.ini", os.O_CREATE|os.O_RDWR, 755)
	if err != nil {
		return err
	}
	if _, err := file.WriteString(strings.Replace(data_temple, "{{projectName}}", projectName, -1)); err != nil {
		return err
	}

	// pre
	file, err = os.OpenFile(projectName+"/data/ini/pre.ini", os.O_CREATE|os.O_RDWR, 755)
	if err != nil {
		return err
	}
	if _, err := file.WriteString(strings.Replace(data_temple, "{{projectName}}", projectName, -1)); err != nil {
		return err
	}

	// test
	file, err = os.OpenFile(projectName+"/data/ini/test.ini", os.O_CREATE|os.O_RDWR, 755)
	if err != nil {
		return err
	}
	if _, err := file.WriteString(strings.Replace(data_temple, "{{projectName}}", projectName, -1)); err != nil {
		return err
	}

	// dev
	file, err = os.OpenFile(projectName+"/data/ini/dev.ini", os.O_CREATE|os.O_RDWR, 755)
	if err != nil {
		return err
	}
	if _, err := file.WriteString(strings.Replace(data_temple, "{{projectName}}", projectName, -1)); err != nil {
		return err
	}

	return nil
}

var data_temple = `[system]
# App名
app_name = {{projectName}}

# 服务器监听的web端口
http_listen_port = 8081

# 服务器监听的gRpc端口
gRpc_listen_port = 18081











[logs]
# 日志目录
dir = /data/logs/{{projectName}}/

# 日志名称
file_name = {{projectName}}.log

# 日志保留时间，单位天
keep_days = 30

# 日志切割间隔，单位小时
rate_hours = 24








[baseServiceSdk]
# 基础账户服务的地址，例如 127.0.0.1:18080
base_service_account_host = 127.0.0.1:15050
base_service_octopus_host = 127.0.0.1:15051
base_service_statistic_host = 127.0.0.1:15052
base_service_user_host = 127.0.0.1:15053
base_service_interact_host = 127.0.0.1:15054












# 数据库相关连接配置
[mysql]
# 连接池最大连接数
maxConn = 100
# 连接池最多的空闲连接
idleConn = 5
# 最大存活时长，单位小时
maxLeftTime = 1


# 数据库连接信息dbName = dbInfo
[dbInfo]
user = root:root@tcp(127.0.0.1:3306)/base?charset=utf8mb4&parseTime=True&loc=Local













[redis]
addr = 127.0.0.1
port = 6379
password =
db_num = 0
`

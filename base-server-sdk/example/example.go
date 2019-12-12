package example

import (
	"encoding/json"
	"github.com/becent/golang-common/base-server-sdk"
	"strconv"
)

func ExampleFunc(a int, b string, c string) (interface{}, *base_server_sdk.Error) {
	params := make(map[string]string)
	params["a"] = strconv.Itoa(a)
	params["b"] = b
	params["c"] = c

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.ExampleServerHost, "example", "test", params)
	if err != nil {
		return nil, err
	}

	var ret interface{}
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, base_server_sdk.ErrServiceBusy
	}

	return ret, nil
}

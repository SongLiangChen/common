package base_server_octopus

import (
	"encoding/json"
	"github.com/becent/golang-common/base-server-sdk"
)

func ExchangeRate(symbol string) (map[string]string, *base_server_sdk.Error) {

	params := make(map[string]string)
	params["symbol"] = symbol

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.OctopusServerHost, "exchange", "rates", params)
	if err != nil {
		return nil, err
	}

	var resp map[string]string
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, base_server_sdk.ErrServiceBusy
	}
	return resp, nil
}

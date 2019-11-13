package gin_handler

type SResponse struct {
	GatewayRet    bool        `json:"gateway-success"`
	GatewayOrgId  int64       `json:"gateway-orgId,string"`
	GatewayUserId int64       `json:"gateway-userId,string"`
	Success       bool        `json:"success"`
	PayLoad       interface{} `json:"payload"`
}

type EResponse struct {
	GatewayRet bool  `json:"gateway-success"`
	Success    bool  `json:"success"`
	Err        Error `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

package base_server_sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/becent/golang-common"
	"github.com/becent/golang-common/gin-handler"
	"github.com/becent/golang-common/grpc-end"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	ErrServiceBusy = &Error{
		Code:    900000,
		Message: "service busy",
	}
	ErrHostEmpty = &Error{
		Code:    900001,
		Message: "host empty, please config it when do InitBaseServerSdk",
	}
	ErrInvalidParams = &Error{
		Code:    900002,
		Message: "params invalid or missed",
	}
	ErrOpenFile = &Error{
		Code:    900003,
		Message: "cannot open the file specified",
	}
)

type Response struct {
	Success bool        `json:"success"`
	PayLoad interface{} `json:"payload"`
	Err     *Error      `json:"error"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) String() string {
	return strconv.Itoa(e.Code) + ": " + e.Message
}

// args is others data that you need. files = args[0]
func (c *BaseServerSdkClient) DoRequest(host string, controller, action string, params map[string]string, args ...map[string]string) ([]byte, *Error) {
	if host == "" {
		return nil, ErrHostEmpty
	}

	var (
		files map[string]string
		data  []byte
		err   error
	)

	if len(args) >= 1 && len(args[0]) > 0 {
		files = args[0]
	}

	if c.gRpcOnly {
		data, err = c.doGRpcRequest(host, controller, action, params, files)
	} else {
		data, err = c.doHttpRequest(host, controller, action, params, files)
	}
	if err != nil {
		common.ErrorLog("baseServerSdk_DoRequest", map[string]interface{}{
			"host":       host,
			"controller": controller,
			"action":     action,
			"params":     params,
		}, err.Error())
		return nil, ErrServiceBusy
	}

	resp := &Response{}
	if err = json.Unmarshal(data, resp); err != nil {
		common.ErrorLog("baseServerSdk_DoRequest", map[string]interface{}{
			"host":       host,
			"controller": controller,
			"action":     action,
			"params":     params,
		}, err.Error())
		return nil, ErrServiceBusy
	}

	if !resp.Success {
		return nil, resp.Err
	}

	if data, err = json.Marshal(resp.PayLoad); err != nil {
		common.ErrorLog("baseServerSdk_DoRequest", map[string]interface{}{
			"host":       host,
			"controller": controller,
			"action":     action,
			"params":     params,
		}, err.Error())
		return nil, ErrServiceBusy
	}

	return data, nil
}

func (c *BaseServerSdkClient) doHttpRequest(host string, controller, action string, params map[string]string, files map[string]string) ([]byte, error) {
	// Assembly body
	contentType, contentReader := newHttpRequestBody(params, files)

	// Make new request
	request, err := http.NewRequest("POST", strings.Join([]string{host, controller, action}, "/"), contentReader)
	if err != nil {
		return nil, err
	}

	// Fill header
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Signature", c.makeSignature())
	request.Header.Set("RequestId", c.requestId())

	// Do http request
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func newHttpRequestBody(params, files map[string]string) (contentType string, contentReader io.Reader) {
	// multipart-formdata
	if len(files) > 0 {
		buf := &bytes.Buffer{}
		writer := multipart.NewWriter(buf)

		//writer text fields
		for key, value := range params {
			writer.WriteField(key, value)
		}

		//writer file fields
		for key, data := range files {
			nn := strings.SplitN(key, gin_handler.FILES_SEPARATOR, -1)
			fw, _ := writer.CreateFormFile(nn[0], nn[1])
			fw.Write([]byte(data))
		}

		contentReader = buf
		// close before call to FormDataContentType ! otherwise its not valid multipart
		writer.Close()
		contentType = writer.FormDataContentType()
		return
	}

	// x-www-form-urlencoded
	v := make(url.Values)
	for key, val := range params {
		v.Set(key, val)
	}
	contentReader = strings.NewReader(v.Encode())
	contentType = "application/x-www-form-urlencoded;charset=utf-8"

	return
}

func (c *BaseServerSdkClient) doGRpcRequest(host string, controller, action string, params map[string]string, files map[string]string) ([]byte, error) {
	// Make new request
	filesMap := make(map[string][]byte, len(files))
	for key, data := range files {
		filesMap[key] = []byte(data)
	}
	request := &grpc_end.Request{
		Controller: controller,
		Action:     action,
		Params:     params,
		Files:      filesMap,
		Header: map[string]string{
			"requestId": c.requestId(),
			"signature": c.makeSignature(),
		},
	}

	// Get a gRpc conn from pool
	pool := c.gRpcMapPool.GetPool(host)
	conn, err := pool.Get()
	if err != nil {
		return nil, err
	}

	// Set the request timeout
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()

	// Do gRpc request
	resp, err := grpc_end.NewEndClient(conn.GetConn()).DoRequest(ctx, request)
	if err != nil {
		pool.DelErrorClient(conn)
		return nil, err
	}
	_ = pool.Put(conn)

	return resp.Data, nil
}

func (c *BaseServerSdkClient) requestId() string {
	return common.Md5Encode(time.Now().String())
}

func (c *BaseServerSdkClient) makeSignature() string {
	now := strconv.FormatInt(time.Now().Unix(), 10)

	b := c.cp.Get().(*strings.Builder)
	defer c.cp.Put(b)
	b.Reset()

	b.WriteString(c.appId)
	b.WriteString("-")
	b.WriteString(now)
	b.WriteString("-")
	b.WriteString(c.appId)

	data := b.String()

	b.Reset()
	b.WriteString(now)
	b.WriteString(":")
	b.WriteString(c.appId)
	b.WriteString(":")
	b.WriteString(common.Hmac(c.appSecretKey, data))

	return b.String()
}

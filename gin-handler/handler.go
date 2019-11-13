package gin_handler

import (
	"errors"
	"github.com/becent/golang-common"
	"github.com/becent/golang-common/exception"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var (
	KEY_APPNAME        = "keyPrefix_AppName"
	KEY_PARAMS         = "keyPrefix_Params"
	KEY_FILES          = "keyPrefix_Files"
	KEY_FILES_SIZE     = "keyPrefix_Files_Size"
	KEY_GIN_CONTROLLER = "keyPrefix_Gin_Controller"
	KEY_RESPONSE       = "keyPrefix_Response"
)

const (
	FILES_SEPARATOR = ";;;"
)

type Handler struct {
	Context *gin.Context
}

type Config struct {
	AppName string

	CheckSignature bool
	AppId          string
	SecretKey      string
}

func NewHandler(cfg *Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.RequestURI == "/" {
			return
		}

		now := time.Now()

		c.Set(KEY_APPNAME, cfg.AppName)

		// get params
		params := make(map[string]string)
		_ = c.Request.ParseForm()
		for key, val := range c.Request.Form {
			if val[0] != "" {
				params[key] = val[0]
			}
		}
		c.Set(KEY_PARAMS, params)

		// if request's Content-Type is multipart/form-data, use MultipartForm() to get text and files.
		if c.ContentType() == gin.MIMEMultipartPOSTForm {
			if form, err := c.MultipartForm(); err == nil {
				for key, val := range form.Value {
					if val[0] != "" {
						params[key] = val[0]
					}
				}

				size := 0
				files := make(map[string]string)
				for fieldName, file := range form.File {
					fileName := file[0].Filename
					if f, err := file[0].Open(); err == nil {
						if content, err := ioutil.ReadAll(f); err == nil {
							key := fieldName + FILES_SEPARATOR + fileName
							files[key] = string(content)
							size += len(files[key])
						}
						f.Close()
					}
				}
				c.Set(KEY_FILES, files)
				c.Set(KEY_FILES_SIZE, size)
			}
		}

		// make a new handler
		h := &Handler{Context: c}
		c.Set(KEY_GIN_CONTROLLER, h)

		// check signature
		if cfg.CheckSignature {
			if err := h.checkSignature(cfg); err != nil {
				h.ErrResponse(&exception.Exception{-1, "签名失败"})
			}
		}

		if !c.IsAborted() {
			c.Next()
		}

		// write log
		resp, _ := c.Get(KEY_RESPONSE)
		entry := log.WithFields(log.Fields{
			"app":      cfg.AppName,
			"action":   c.Request.RequestURI,
			"params":   params,
			"header":   c.Request.Header,
			"clientIp": c.ClientIP(),
			"response": resp,
			"useTime":  time.Since(now).String(),
		})

		err := c.Errors.Last()
		if err != nil && err.Type > 0 {
			entry.Error(err)
			return
		}
		entry.Info("")
	}
}

func DefaultHandler(c *gin.Context) *Handler {
	return c.MustGet(KEY_GIN_CONTROLLER).(*Handler)
}

// ---------------------------------------------------------------------------------------------------------------------
// ---------------------------------------------------------------------------------------------------------------------

// SuccessResponse
// ids: orgId、userId
func (ct *Handler) SuccessResponse(payload interface{}, ids ...int64) {
	if payload == nil {
		payload = make(map[string]interface{})
	}

	resp := SResponse{
		GatewayRet: true,
		Success:    true,
		PayLoad:    payload,
	}
	if len(ids) == 2 {
		resp.GatewayOrgId = ids[0]
		resp.GatewayUserId = ids[1]
	}

	c := ct.Context
	c.Set(KEY_RESPONSE, resp)
	c.JSON(200, resp)
}

// ErrResponse
func (ct *Handler) ErrResponse(exp *exception.Exception, logMessage ...string) {
	ep := EResponse{
		GatewayRet: false,
		Success:    false,
		Err: Error{
			Code:    exp.Code,
			Message: exp.Message,
		},
	}

	msg := exp.Message
	if len(logMessage) > 0 {
		msg += ": " + strings.Join(logMessage, ",")
	}

	c := ct.Context
	c.Set(KEY_RESPONSE, ep)
	_ = c.AbortWithError(200, errors.New(msg)).SetType(gin.ErrorType(exp.Code))
	c.JSON(200, ep)
}

// ---------------------------------------------------------------------------------------------------------------------
// ---------------------------------------------------------------------------------------------------------------------

// IntParam returns int for the given key
// If the value does not exists it returns 0
func (ct *Handler) IntParam(key string) int {
	val, _ := strconv.Atoi(ct.StringParam(key))
	return val
}

// IntParamDefault returns int for the given key
// If the value does not exists it returns defVal
func (ct *Handler) IntParamDefault(key string, defVal int) int {
	ret := ct.IntParam(key)
	if ret == 0 {
		return defVal
	}
	return ret
}

// Int64Param returns int64 for the given key
// If the value does not exists it returns 0
func (ct *Handler) Int64Param(key string) int64 {
	val, _ := strconv.ParseInt(ct.StringParam(key), 10, 64)
	return val
}

// Int64ParamDefault returns int64 for the given key
// If the value does not exists it returns defVal
func (ct *Handler) Int64ParamDefault(key string, defVal int64) int64 {
	ret := ct.Int64Param(key)
	if ret == 0 {
		return defVal
	}
	return ret
}

// StringParam returns string for the given key
// If the value does not exists it returns empty string
func (ct *Handler) StringParam(key string) string {
	c := ct.GetContext()
	params := c.GetStringMapString(KEY_PARAMS)

	ret := params[key]
	if ret == "" {
		ret = c.GetHeader("X-Gw-" + key)
	}

	return ret
}

func (ct *Handler) StringFile(key string) string {
	files := ct.MapFiles()
	return files[key]
}

// GetStringMapString returns the value associated with the key as a map of strings.
func (ct *Handler) MapFiles() (sms map[string]string) {
	c := ct.GetContext()
	files := c.GetStringMapString(KEY_FILES)
	return files
}

// StringParamDefault returns string for the given key
// If the value does not exists it returns defVal
func (ct *Handler) StringParamDefault(key string, defVal string) string {
	v := ct.StringParam(key)
	if v == "" {
		v = defVal
	}
	return v
}

// IsParamExist returns true is request's param exist for the given key
func (ct *Handler) IsParamExist(key string) bool {
	_, ok := ct.GetContext().GetPostForm(key)
	if !ok {
		_, ok = ct.GetContext().GetPostForm(key)
	}
	return ok
}

// ---------------------------------------------------------------------------------------------------------------------
// ---------------------------------------------------------------------------------------------------------------------

// makeRequestSignature returns signature
func (ct *Handler) makeRequestSignature(appId, appSecret string, now string) string {
	if now == "" {
		now = strconv.FormatInt(time.Now().Unix(), 10)
	}
	data := strings.Join([]string{appId, now, appId}, "-")

	return strings.Join([]string{now, appId, common.Hmac(appSecret, data)}, ":")
}

// checkSignature returns nil if signature valid
func (ct *Handler) checkSignature(cfg *Config) error {
	c := ct.GetContext()

	sign := c.GetHeader("Signature")
	if sign == "" {
		return errors.New("签名为空")
	}

	ss := strings.SplitN(sign, ":", -1)
	if len(ss) != 3 {
		return errors.New("签名格式异常")
	}

	st := ss[0]
	if ist, err := strconv.ParseInt(st, 10, 64); err != nil || ss[1] == "" || ss[2] == "" {
		return errors.New("签名格式异常")

	} else {
		sub := time.Now().Unix() - ist
		if sub < 0 {
			sub *= -1
		}
		if sub > 5 {
			return errors.New("时间戳已过期")
		}
	}

	if sign != ct.makeRequestSignature(cfg.AppId, cfg.SecretKey, st) {
		return errors.New("无效的签名")
	}

	return nil
}

// ---------------------------------------------------------------------------------------------------------------------
// ---------------------------------------------------------------------------------------------------------------------

func (ct *Handler) GetContext() *gin.Context {
	return ct.Context
}

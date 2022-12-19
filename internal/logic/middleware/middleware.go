package middleware

import (
	"blog-api/api/admin"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

var authService *jwt.GfJWTMiddleware

// Auth 权限包管理
func Auth() *jwt.GfJWTMiddleware {
	return authService
}

// 初始化
func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		//用户的领域名称，必传
		Realm: "mall",
		// 签名算法
		SigningAlgorithm: "HS256",
		// 签名密钥
		Key: []byte("blog"),
		// 时效
		Timeout: time.Minute * 60 * 6,
		// 	token过期后，可凭借旧token获取新token的刷新时间
		MaxRefresh: time.Minute * 5,
		// 身份验证的key值
		IdentityKey: "user_id",
		//token检索模式，用于提取token-> Authorization
		// TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenLookup: "header: Authorization",
		// token在请求头时的名称，默认值为Bearer.客户端在header中传入"Authorization":"token xxxxxx"
		TokenHeadName: "",
		TimeFunc:      time.Now,
		// 用户标识 map  私有属性
		// 根据登录信息对用户进行身份验证的回调函数
		Authenticator: Authenticator,
		// 处理不进行授权的逻辑
		Unauthorized: Unauthorized,
		//登录期间的设置私有载荷的函数，默认设置Authenticator函数回调的所有内容
		PayloadFunc: PayloadFunc,
		// 解析并设置用户身份信息，并设置身份信息至每次请求中
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}

// PayloadFunc 向webtoken添加额外的有效负载数据。
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	// params := data.(map[string]interface{})
	params := gconv.Map(data)
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler 标识
func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(g.Map{
		"code":    code,
		"message": message,
	})
	r.ExitAll()
}

// Authenticator 用户标识  私有载荷
func Authenticator(ctx context.Context) (interface{}, error) {
	var (
		r  = g.RequestFromCtx(ctx)
		in admin.UserLoginReq
	)
	if err := r.Parse(&in); err != nil {
		return "", err
	}

	user, err := service.User().CheckUserAccount(ctx, model.UserAccountInput{
		Username: in.Username,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}
	data := gconv.Map(user)
	data["user_id"] = data["id"]
	return &data, nil
}

type CustomHandlerResponse struct {
	Code int         `json:"code"    dc:"Error code"`
	Msg  string      `json:"msg" dc:"Error message"`
	Data interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

// =========================================================================================================

type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) Jwt(r *ghttp.Request) {
	Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

// CustomResponseHandler ResponseHandler 返回处理中间件
func (s *sMiddleware) CustomResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	//var (
	//	msg  = "success"
	//	err  = r.GetError()
	//	res  = r.GetHandlerResponse()
	//	code = gerror.Code(err)
	//)
	//if err != nil {
	//	if code == gcode.CodeNil {
	//		code = gcode.CodeInternalError
	//	}
	//	msg = err.Error()
	//} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
	//	msg = http.StatusText(r.Response.Status)
	//	switch r.Response.Status {
	//	case http.StatusNotFound:
	//		code = gcode.CodeNotFound
	//	case http.StatusForbidden:
	//		code = gcode.CodeNotAuthorized
	//	default:
	//		code = gcode.CodeUnknown
	//	}
	//} else {
	//	code = gcode.New(200, msg, "")
	//}
	//r.Response.WriteJson(CustomHandlerResponse{
	//	Code: code.Code(),
	//	Msg:  code.Message(),
	//	Data: res,
	//})

	var (
		err             = r.GetError()
		res             = r.GetHandlerResponse()
		code gcode.Code = gcode.CodeOK
	)
	if err != nil {
		code = gerror.Code(err)
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		JsonExit(r, code.Code(), err.Error())
		r.Exit()
	} else {
		JsonExit(r, 200, "", res)
	}
}

// JsonRes 数据返回通用JSON数据结构
type JsonRes struct {
	Code    int         `json:"code"` // 默认200
	Message string      `json:"msg"`  // 提示信息
	Data    interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
}

// Json 返回标准JSON数据。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = g.Map{}
	}
	r.Response.WriteJson(JsonRes{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// JsonExit 返回标准JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, code int, message string, data ...interface{}) {
	Json(r, code, message, data...)
	r.Exit()
}

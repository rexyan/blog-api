// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ExceptionLog is the golang structure for table exception_log.
type ExceptionLog struct {
	Id          int64       `json:"id"          description:""`
	Uri         string      `json:"uri"         description:"请求接口"`
	Method      string      `json:"method"      description:"请求方式"`
	Param       string      `json:"param"       description:"请求参数"`
	Description string      `json:"description" description:"操作描述"`
	Error       string      `json:"error"       description:"异常信息"`
	Ip          string      `json:"ip"          description:"ip"`
	IpSource    string      `json:"ipSource"    description:"ip来源"`
	Os          string      `json:"os"          description:"操作系统"`
	Browser     string      `json:"browser"     description:"浏览器"`
	CreateTime  *gtime.Time `json:"createTime"  description:"操作时间"`
	UserAgent   string      `json:"userAgent"   description:"user-agent用户代理"`
}

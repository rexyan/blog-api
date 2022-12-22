package admin

import "github.com/gogf/gf/v2/frame/g"

type GetJobLogListReq struct {
	g.Meta        `path:"/job/logs" tags:"日志" method:"get" summary:"定时任务日志列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Page
	Date string `json:"date"`
}

type GetJobLogListRes struct {
	CommonPageHelper
	List []struct {
		Job
		Error string `json:"error"`
		Times int    `json:"times"`
		LogId int    `json:"logId"`
	} `json:"list"`
}

type DeleteJobLogReq struct {
	g.Meta        `path:"/job/log" tags:"日志" method:"delete" summary:"删除定时任务日志"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	LogId         int    `json:"logId"`
}

type DeleteJobLogRes struct {
}

type GetLoginLogListReq struct {
	g.Meta        `path:"/loginLogs" tags:"日志" method:"get" summary:"登录日志列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Page
	Date string `json:"date"`
}

type GetLoginLogListRes struct {
	CommonPageHelper
	List []struct {
		ID          int         `json:"id"`
		Username    string      `json:"username"`
		IP          string      `json:"ip"`
		IPSource    string      `json:"ipSource"`
		Os          string      `json:"os"`
		Browser     string      `json:"browser"`
		Status      bool        `json:"status"`
		Description string      `json:"description"`
		CreateTime  string      `json:"createTime"`
		UserAgent   interface{} `json:"userAgent"`
	} `json:"list"`
}

type DeleteLoginLogReq struct {
	g.Meta        `path:"/loginLog" tags:"日志" method:"delete" summary:"删除定时任务日志"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	LogId         int    `json:"id"`
}

type DeleteLoginLogRes struct {
}

type GetOperationLogListReq struct {
	g.Meta        `path:"/operationLogs" tags:"日志" method:"get" summary:"操作日志列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Page
	Date string `json:"date"`
}

type GetOperationLogListRes struct {
	CommonPageHelper
	List []struct {
		ID          int         `json:"id"`
		Username    string      `json:"username"`
		URI         string      `json:"uri"`
		Method      string      `json:"method"`
		Param       string      `json:"param"`
		Description string      `json:"description"`
		IP          string      `json:"ip"`
		IPSource    string      `json:"ipSource"`
		Os          string      `json:"os"`
		Browser     string      `json:"browser"`
		Times       int         `json:"times"`
		CreateTime  string      `json:"createTime"`
		UserAgent   interface{} `json:"userAgent"`
	} `json:"list"`
}

type DeleteOperationLogReq struct {
	g.Meta        `path:"/operationLog" tags:"日志" method:"delete" summary:"删除操作日志"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	LogId         int    `json:"id"`
}

type DeleteOperationLogRes struct {
}

type GetExceptionLogListReq struct {
	g.Meta        `path:"/exceptionLogs" tags:"日志" method:"get" summary:"异常日志列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Page
	Date string `json:"date"`
}

type GetExceptionLogListRes struct {
	CommonPageHelper
	List []struct {
		ID          int         `json:"id"`
		URI         string      `json:"uri"`
		Method      string      `json:"method"`
		Param       string      `json:"param"`
		Description string      `json:"description"`
		Error       string      `json:"error"`
		IP          string      `json:"ip"`
		IPSource    string      `json:"ipSource"`
		Os          string      `json:"os"`
		Browser     string      `json:"browser"`
		CreateTime  string      `json:"createTime"`
		UserAgent   interface{} `json:"userAgent"`
	} `json:"list"`
}

type DeleteExceptionLogReq struct {
	g.Meta        `path:"/exceptionLog" tags:"日志" method:"delete" summary:"删除异常日志"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	LogId         int    `json:"id"`
}

type DeleteExceptionLogRes struct {
}

type GetVisitLogListReq struct {
	g.Meta        `path:"/visitLogs" tags:"日志" method:"get" summary:"访问日志列表"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	Page
	Date      string `json:"date"`
	VisitorId string `json:"uuid"`
}

type GetVisitLogListRes struct {
	CommonPageHelper
	List []struct {
		ID         int         `json:"id"`
		UUID       string      `json:"uuid"`
		URI        string      `json:"uri"`
		Method     string      `json:"method"`
		Param      string      `json:"param"`
		Behavior   string      `json:"behavior"`
		Content    string      `json:"content"`
		Remark     string      `json:"remark"`
		IP         string      `json:"ip"`
		IPSource   string      `json:"ipSource"`
		Os         string      `json:"os"`
		Browser    string      `json:"browser"`
		Times      int         `json:"times"`
		CreateTime string      `json:"createTime"`
		UserAgent  interface{} `json:"userAgent"`
	} `json:"list"`
}

type DeleteVisitLogReq struct {
	g.Meta        `path:"/visitLog" tags:"日志" method:"delete" summary:"删除访问日志"`
	Authorization string `json:"Authorization" in:"header"  dc:"Authorization"`
	LogId         int    `json:"id"`
}

type DeleteVisitLogRes struct {
}

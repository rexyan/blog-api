// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SiteSetting is the golang structure of table site_setting for DAO operations like Where/Data.
type SiteSetting struct {
	g.Meta `orm:"table:site_setting, do:true"`
	Id     interface{} //
	NameEn interface{} //
	NameZh interface{} //
	Value  interface{} //
	Type   interface{} // 1基础设置，2页脚徽标，3资料卡，4友链信息
}

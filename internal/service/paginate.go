// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"blog-api/internal/model"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gpage"
)

type IPaginate interface {
	Paginate(res *model.CommonPageHelper, count int64, page model.PageInput, result gdb.Result, pageInfo *gpage.Page) *model.CommonPageHelper
}

var localPaginate IPaginate

func Paginate() IPaginate {
	if localPaginate == nil {
		panic("implement not found for interface IPaginate, forgot register?")
	}
	return localPaginate
}

func RegisterPaginate(i IPaginate) {
	localPaginate = i
}
package paginate

import (
	"blog-api/internal/model"
	"blog-api/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gpage"
)

type sPaginate struct{}

func init() {
	service.RegisterPaginate(New())
}

func New() *sPaginate {
	return &sPaginate{}
}

func (s *sPaginate) Paginate(res *model.CommonPageHelper, count int64, page model.PageInput, result gdb.Result, pageInfo *gpage.Page) *model.CommonPageHelper {
	// ==============  分页信息 ==============
	res.Total = gconv.Int(count)
	res.PageNum = page.PageNum
	res.PageSize = page.PageSize
	res.Size = len(result)
	res.StartRow = (page.PageSize-1)*page.PageNum + 1
	res.EndRow = page.PageSize * page.PageNum
	res.Pages = pageInfo.TotalPage
	res.PrePage = pageInfo.CurrentPage - 1

	// 是否是第一个页
	if pageInfo.CurrentPage == 0 {
		res.IsFirstPage = true
		res.HasPreviousPage = false
		res.NavigateFirstPage = 1
		res.PrePage = 0
	} else {
		res.IsFirstPage = false
		res.HasPreviousPage = true
		res.NavigateFirstPage = pageInfo.CurrentPage - 1
		res.PrePage = pageInfo.CurrentPage - 1
	}
	// 是否是最后一页
	if pageInfo.CurrentPage != pageInfo.TotalPage {
		res.IsLastPage = false
		res.HasNextPage = true
		res.NavigateLastPage = pageInfo.CurrentPage + 1
		res.NextPage = pageInfo.CurrentPage + 1
	} else {
		res.IsLastPage = true
		res.HasNextPage = false
		res.NavigateLastPage = pageInfo.TotalPage
		res.NextPage = 0
	}
	res.NavigatePages = 8
	for i := 1; i <= pageInfo.TotalPage; i++ {
		res.NavigatepageNums = append(res.NavigatepageNums, i)
	}
	return res
}

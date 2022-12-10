package blog

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sBlog struct{}

func init() {
	service.RegisterBlog(New())
}
func New() *sBlog {
	return &sBlog{}
}

// GetBlogCount 获取博客计数
func (s *sBlog) GetBlogCount(ctx context.Context) (count int64) {
	count, err := dao.Blog.Ctx(ctx).Count()
	if err != nil {
		return
	}
	return
}

// GetBlogCountByTagId 根据 TagId 获取博客计数
func (s *sBlog) GetBlogCountByTagId(ctx context.Context, TagId int64) (count int64) {
	count, err := dao.BlogTag.Ctx(ctx).Count(dao.BlogTag.Columns().TagId, TagId)
	if err != nil {
		return 0
	}
	return
}

// GetBlogCountByCategoryId 根据 CateGoryId 获取博客计数
func (s *sBlog) GetBlogCountByCategoryId(ctx context.Context, CategoryID int64) (count int64) {
	count, err := dao.Blog.Ctx(ctx).Count(dao.Blog.Columns().CategoryId, CategoryID)
	if err != nil {
		return 0
	}
	return
}

// CreateBlog 创建博客文章
func (s *sBlog) CreateBlog(ctx context.Context, in model.CreateBlogInput) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		blogCls := dao.Blog.Columns()
		blogLastInsertId, err := tx.Ctx(ctx).InsertAndGetId(dao.Blog.Table(), g.Map{
			blogCls.IsAppreciation:   in.IsAppreciation,
			blogCls.CategoryId:       in.CategoryId,
			blogCls.IsCommentEnabled: in.IsCommentEnabled,
			blogCls.Content:          in.Content,
			blogCls.Description:      in.Description,
			blogCls.FirstPicture:     in.FirstPicture,
			blogCls.Password:         in.Password,
			blogCls.IsPublished:      in.IsPublished,
			blogCls.ReadTime:         in.ReadTime,
			blogCls.IsRecommend:      in.IsRecommend,
			blogCls.Title:            in.Title,
			blogCls.IsTop:            in.IsTop,
			blogCls.Views:            in.Views,
			blogCls.Words:            in.Words,
			blogCls.UserId:           in.UserId,
			blogCls.CreateTime:       gtime.Datetime(),
			blogCls.UpdateTime:       gtime.Datetime(),
		})
		if err != nil {
			return err
		}

		// 文章关联标签
		for _, tagId := range in.TagList {
			_, err := s.BlogAssociationTag(ctx, gconv.Int(blogLastInsertId), tagId)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}

// BlogAssociationTag 新增博客关联标签
func (s *sBlog) BlogAssociationTag(ctx context.Context, BlogId int, TagId int) (res int, err error) {
	lastInsertId, err := dao.BlogTag.Ctx(ctx).InsertAndGetId(g.Map{
		dao.BlogTag.Columns().BlogId: BlogId,
		dao.BlogTag.Columns().TagId:  TagId,
	})
	if err != nil {
		return 0, err
	}
	return gconv.Int(lastInsertId), nil
}

func (s *sBlog) GetBlogs(ctx context.Context, title string, page model.PageInput) (res *model.BlogListOutput, err error) {
	r := g.RequestFromCtx(ctx)
	res = &model.BlogListOutput{}
	query := dao.Blog.Ctx(ctx)

	// 是否根据 title 过滤
	if len(title) > 0 {
		query = query.WhereLike(dao.Blog.Columns().Title, title)
	}
	result, err := query.Page(page.PageNum, page.PageSize).All()
	if err != nil {
		return &model.BlogListOutput{}, err
	}

	// 总数
	count, err := query.Count()
	if err != nil {
		return nil, err
	}
	// 页码信息
	pageInfo := r.GetPage(gconv.Int(count), page.PageSize)
	err = gconv.Scan(result, &res.List)
	if err != nil {
		return nil, err
	}
	// 查询每个 blog 的 category 信息
	err = dao.Category.Ctx(ctx).
		Where(dao.Category.Columns().Id, gdb.ListItemValuesUnique(res.List, "CategoryId")).
		ScanList(&res.List, "Category")
	if err != nil {
		return nil, err
	}

	// ==============  分页信息 ==============
	res.Total = gconv.Int(count)
	res.PageNum = page.PageNum
	res.PageSize = page.PageSize
	res.Size = len(result)
	res.StartRow = (page.PageSize-1)*page.PageNum + 1
	res.EndRow = page.PageSize * page.PageNum
	res.Pages = pageInfo.TotalPage
	res.PrePage = pageInfo.CurrentPage - 1
	res.NextPage = pageInfo.CurrentPage + 1
	// 是否是第一个页
	if pageInfo.CurrentPage == 0 {
		res.IsFirstPage = true
		res.HasPreviousPage = false
		res.NavigateFirstPage = 1
	} else {
		res.IsFirstPage = false
		res.HasPreviousPage = true
		res.NavigateFirstPage = pageInfo.CurrentPage - 1
	}
	// 是否是最后一页
	if pageInfo.CurrentPage != pageInfo.TotalPage {
		res.IsLastPage = false
		res.HasNextPage = true
		res.NavigateLastPage = pageInfo.CurrentPage + 1
	} else {
		res.IsLastPage = true
		res.HasNextPage = false
		res.NavigateLastPage = pageInfo.TotalPage
	}
	res.NavigatePages = 8
	for i := 1; i <= pageInfo.TotalPage; i++ {
		res.NavigatepageNums = append(res.NavigatepageNums, i)
	}
	return
}

func (s *sBlog) GetBlogDetail(ctx context.Context, BlogId int) (res *model.BlogDetailOutput, err error) {
	res = &model.BlogDetailOutput{}

	err = dao.Blog.Ctx(ctx).Where(dao.Blog.Columns().Id, BlogId).Scan(&res)
	if err != nil {
		return nil, err
	}

	// 此种写法存在问题（猜测可能是 BlogDetailOutput 中存在 Struct 继承关系）
	//err = dao.Category.Ctx(ctx).
	//	Where(dao.Category.Columns().Id, res.CategoryId).
	//	ScanList(&res, "Category")
	//if err != nil {
	//	return nil, err
	//}

	err = dao.Category.Ctx(ctx).Where(dao.Category.Columns().Id, res.CategoryId).Scan(&res.Category)
	if err != nil {
		return nil, err
	}

	tagList, err := service.Tag().GetTagIdsByBlogId(ctx, BlogId)
	if err != nil {
		return nil, err
	}

	err = dao.Tag.Ctx(ctx).Where(dao.Tag.Columns().Id, gdb.ListItemValuesUnique(tagList, "TagId")).Scan(&res.Tags)
	if err != nil {
		return nil, err
	}

	return
}

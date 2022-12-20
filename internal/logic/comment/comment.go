package comment

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sComment struct{}

func init() {
	service.RegisterComment(New())
}

func New() *sComment {
	return &sComment{}
}

// GetCommentCount 评论计数
func (s *sComment) GetCommentCount(ctx context.Context) (count int64) {
	count, err := dao.Comment.Ctx(ctx).Count()
	if err != nil {
		return
	}
	return
}

func (s *sComment) GetCommentList(ctx context.Context, page model.PageInput) (res *model.GetCommentListOutput, err error) {
	r := g.RequestFromCtx(ctx)
	res = &model.GetCommentListOutput{}
	query := dao.Comment.Ctx(ctx)

	result, err := query.Page(page.PageNum, page.PageSize).All()
	if err != nil {
		return &model.GetCommentListOutput{}, err
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
	// 查找子评论
	for i := 0; i < len(res.List); i++ {
		DepthComment(ctx, &res.List[i])
	}

	service.Paginate().Paginate(&res.CommonPageHelper, count, page, result, pageInfo)
	return
}

func DepthComment(ctx context.Context, commentItem *model.CommentItem) {
	replyComments := []model.CommentItem{}
	// 查询该条评论的子评论
	parentComments, _ := dao.Comment.Ctx(ctx).Where(dao.Comment.Columns().ParentCommentId, commentItem.ID).All()
	if parentComments == nil {
		commentItem.ReplyComments = replyComments
	} else {
		_ = gconv.Scan(parentComments, &replyComments)
		commentItem.ReplyComments = replyComments
		// 循环每个子评论
		for i := 0; i < len(replyComments); i++ {
			// 递归
			DepthComment(ctx, &replyComments[i])
		}
	}
}

func (s *sComment) UpdateCommentNotice(ctx context.Context, CommentId int, NoticeStatus bool) (err error) {
	_, err = dao.Comment.Ctx(ctx).Data(dao.Comment.Columns().IsNotice, NoticeStatus).Where(dao.Comment.Columns().Id, CommentId).Update()
	return
}

func (s *sComment) UpdateCommentPublish(ctx context.Context, CommentId int, PublishStatus bool) (err error) {
	_, err = dao.Comment.Ctx(ctx).Data(dao.Comment.Columns().IsPublished, PublishStatus).Where(dao.Comment.Columns().Id, CommentId).Update()
	return
}

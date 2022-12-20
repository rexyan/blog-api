package friend

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sFriend struct{}

func init() {
	service.RegisterFriend(New())
}

func New() *sFriend {
	return &sFriend{}
}

func (s *sFriend) GetFriendInfo(ctx context.Context) (res *model.GetFriendInfoOutput, err error) {
	res = &model.GetFriendInfoOutput{}
	// 查询 friendContent
	friendContent, err := dao.SiteSetting.Ctx(ctx).
		Fields(dao.SiteSetting.Columns().Value).
		Where(dao.SiteSetting.Columns().NameEn, "friendContent").
		All()

	if err != nil {
		return nil, err
	}
	if len(friendContent) > 0 {
		res.Content = gconv.String(friendContent[0])
	} else {
		res.Content = ""
	}

	// 查询 friendCommentEnabled
	friendCommentEnabled, err := dao.SiteSetting.Ctx(ctx).
		Fields(dao.SiteSetting.Columns().Value).
		Where(dao.SiteSetting.Columns().NameEn, "friendCommentEnabled").
		All()
	if err != nil {
		return nil, err
	}
	if len(friendCommentEnabled) > 0 {
		res.CommentEnabled = gconv.Bool(friendCommentEnabled[0]["value"])
	} else {
		res.CommentEnabled = false
	}
	return
}

func (s *sFriend) UpdateFriendInfo(ctx context.Context, Content string) (err error) {
	_, err = dao.SiteSetting.Ctx(ctx).Data(dao.SiteSetting.Columns().Value, Content).Where(dao.SiteSetting.Columns().NameEn, "friendContent").Update()
	return
}

func (s *sFriend) GetFriendList(ctx context.Context, page model.PageInput) (res *model.GetFriendListOutput, err error) {
	r := g.RequestFromCtx(ctx)
	res = &model.GetFriendListOutput{}
	query := dao.Friend.Ctx(ctx)

	result, err := query.Page(page.PageNum, page.PageSize).All()
	if err != nil {
		return &model.GetFriendListOutput{}, err
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

	service.Paginate().Paginate(&res.CommonPageHelper, count, page, result, pageInfo)
	return
}

func (s *sFriend) UpdateFriendPublishStatus(ctx context.Context, FriendId int, PublishStatus bool) (err error) {
	_, err = dao.Friend.Ctx(ctx).Data(g.Map{
		dao.Friend.Columns().IsPublished: PublishStatus,
	}).Where(dao.Friend.Columns().Id, FriendId).Update()
	return
}

func (s *sFriend) UpdateFriend(ctx context.Context, in *model.UpdateFriendInput) (err error) {
	friendCls := dao.Friend.Columns()
	_, err = dao.Friend.Ctx(ctx).Data(g.Map{
		friendCls.Avatar:      in.Avatar,
		friendCls.CreateTime:  in.CreateTime,
		friendCls.Description: in.Description,
		friendCls.Nickname:    in.Nickname,
		friendCls.IsPublished: in.IsPublished,
		friendCls.Views:       in.Views,
		friendCls.Website:     in.Website,
	}).Where(friendCls.Id, in.Id).Update()
	return
}

func (s *sFriend) CreateFriend(ctx context.Context, in *model.CreateFriendInput) (err error) {
	friendCls := dao.Friend.Columns()
	_, err = dao.Friend.Ctx(ctx).Data(g.Map{
		friendCls.Avatar:      in.Avatar,
		friendCls.Description: in.Description,
		friendCls.Nickname:    in.Nickname,
		friendCls.IsPublished: in.IsPublished,
		friendCls.Website:     in.Website,
		friendCls.Views:       0,
		friendCls.CreateTime:  gtime.Now(),
	}).Insert()
	return
}

func (s *sFriend) UpdateCommentEnabled(ctx context.Context, CommentEnabled bool) (err error) {
	_, err = dao.SiteSetting.Ctx(ctx).Data(dao.SiteSetting.Columns().Value, CommentEnabled).Where(dao.SiteSetting.Columns().NameEn, "friendCommentEnabled").Update()
	return
}

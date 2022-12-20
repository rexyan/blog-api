package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/model"
	"blog-api/internal/model/entity"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var Friend = cFriend{}

type cFriend struct{}

func (c *cFriend) GetFriendInfo(ctx context.Context, req *admin.GetFriendInfoReq) (res *admin.GetFriendInfoRes, err error) {
	friendInfo, err := service.Friend().GetFriendInfo(ctx)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(friendInfo, &res)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cFriend) UpdateFriendInfo(ctx context.Context, req *admin.UpdateFriendInfoReq) (res *admin.UpdateFriendInfoRes, err error) {
	err = service.Friend().UpdateFriendInfo(ctx, req.Content)
	return
}

func (c *cFriend) GetFriendList(ctx context.Context, req *admin.GetFriendListReq) (res *admin.GetFriendListRes, err error) {
	friendList, err := service.Friend().GetFriendList(ctx, model.PageInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(friendList, &res)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cFriend) UpdateFriendPublishStatus(ctx context.Context, req *admin.UpdateFriendPublishStatusReq) (res *admin.UpdateFriendPublishStatusRes, err error) {
	err = service.Friend().UpdateFriendPublishStatus(ctx, req.ID, req.IsPublished)
	return
}

func (c *cFriend) UpdateFriend(ctx context.Context, req *admin.UpdateFriendReq) (res *admin.UpdateFriendRes, err error) {
	err = service.Friend().UpdateFriend(ctx, &model.UpdateFriendInput{
		Friend: entity.Friend{
			Id:          gconv.Int64(req.ID),
			Nickname:    req.Nickname,
			Description: req.Description,
			Website:     req.Website,
			Avatar:      req.Avatar,
			IsPublished: gconv.Int(req.IsPublished),
			Views:       req.Views,
			CreateTime:  gtime.New(req.CreateTime),
		},
	})
	return
}

func (c *cFriend) CreateFriend(ctx context.Context, req *admin.CreateFriendReq) (res *admin.CreateFriendRes, err error) {
	err = service.Friend().CreateFriend(ctx, &model.CreateFriendInput{
		Friend: entity.Friend{
			Nickname:    req.Nickname,
			Description: req.Description,
			Website:     req.Website,
			Avatar:      req.Avatar,
			IsPublished: gconv.Int(req.IsPublished),
		},
	})
	return
}

func (c *cFriend) UpdateCommentEnabled(ctx context.Context, req *admin.UpdateCommentEnabledReq) (res *admin.UpdateCommentEnabledRes, err error) {
	err = service.Friend().UpdateCommentEnabled(ctx, req.CommentEnabled)
	return
}

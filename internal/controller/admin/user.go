package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) Login(ctx context.Context, req *admin.UserLoginReq) (res *admin.UserLoginRes, err error) {
	res = &admin.UserLoginRes{}
	user, err := service.User().CheckUserAccount(ctx, model.UserAccountInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	res.Token = service.User().GenerateJwtToken(ctx)
	err = gconv.Scan(user, &res.User)
	if err != nil {
		return nil, err
	}
	return
}

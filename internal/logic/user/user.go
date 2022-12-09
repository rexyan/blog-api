package user

import (
	"blog-api/internal/dao"
	"blog-api/internal/logic/middleware"
	"blog-api/internal/model"
	"blog-api/internal/model/entity"
	"blog-api/internal/service"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// HmacSha256
func (s *sUser) hmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// GenerateJwtToken 生成 jwt token
func (s *sUser) GenerateJwtToken(ctx context.Context) string {
	tokenString, _ := middleware.Auth().LoginHandler(ctx)
	return "Bearer " + tokenString
}

// CheckUserAccount 检查用户账号密码是否正确
func (s *sUser) CheckUserAccount(ctx context.Context, in model.UserAccountInput) (user *entity.User, err error) {
	secretKey, err := g.Cfg().Get(ctx, "service.secretKey")
	if err != nil {
		return nil, err
	}
	hashPassWord := s.hmacSha256(in.Password, gconv.String(secretKey))
	err = dao.User.Ctx(ctx).Where(dao.User.Columns().Username, in.Username).Where(dao.User.Columns().Password, hashPassWord).Scan(&user)
	if err != nil {
		return nil, nil
	}
	return
}

// GetUserInfo 获取用户信息
func (s *sUser) GetUserInfo(ctx context.Context, userId int32) (user *entity.User) {
	err := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, userId).Scan(&user)
	if err != nil {
		return nil
	}
	return
}

package about

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAbout struct{}

func init() {
	service.RegisterAbout(New())
}
func New() *sAbout {
	return &sAbout{}
}

func (s *sAbout) GetAbout(ctx context.Context) (res *model.GetAboutOutput, err error) {
	all, err := dao.About.Ctx(ctx).All()
	if err != nil {
		return nil, err
	}
	res = &model.GetAboutOutput{}
	for i := 0; i < len(all); i++ {
		switch gconv.String(all[i]["name_en"]) {
		case "title":
			res.Title = gconv.String(all[i]["value"])
		case "musicId":
			res.MusicId = gconv.String(all[i]["value"])
		case "content":
			res.Content = gconv.String(all[i]["value"])
		case "commentEnabled":
			res.CommentEnabled = gconv.String(all[i]["value"])
		}
	}
	return
}

func (s *sAbout) UpdateAbout(ctx context.Context, in *model.UpdateAboutInput) (err error) {
	_, err = dao.About.Ctx(ctx).Data(dao.About.Columns().Value, in.Title).Where(dao.About.Columns().NameEn, "title").Update()
	if err != nil {
		return err
	}
	_, err = dao.About.Ctx(ctx).Data(dao.About.Columns().Value, in.MusicId).Where(dao.About.Columns().NameEn, "musicId").Update()
	if err != nil {
		return err
	}
	_, err = dao.About.Ctx(ctx).Data(dao.About.Columns().Value, in.Content).Where(dao.About.Columns().NameEn, "content").Update()
	if err != nil {
		return err
	}
	_, err = dao.About.Ctx(ctx).Data(dao.About.Columns().Value, in.CommentEnabled).Where(dao.About.Columns().NameEn, "commentEnabled").Update()
	if err != nil {
		return err
	}
	return
}

package setSetting

import (
	"blog-api/internal/dao"
	"blog-api/internal/model"
	"blog-api/internal/model/entity"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSetSetting struct{}

func init() {
	service.RegisterSetSetting(New())
}
func New() *sSetSetting {
	return &sSetSetting{}
}

func (s *sSetSetting) GetSetSetting(ctx context.Context, SetSetting string) (res []*entity.SiteSetting, err error) {
	all, err := dao.SiteSetting.Ctx(ctx).Where(dao.SiteSetting.Columns().Type, SetSetting).All()
	err = gconv.Scan(all, &res)
	return
}

func (s *sSetSetting) UpdateSetSetting(ctx context.Context, in *model.UpdateSetSettingInput) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		siteSettingCls := dao.SiteSetting.Columns()
		// 删除
		for _, deleteId := range in.DeleteIds {
			_, err = tx.Ctx(ctx).Delete(dao.SiteSetting.Table(), siteSettingCls.Id, deleteId)
			if err != nil {
				return err
			}
		}

		// 修改，新增
		for _, setting := range in.Settings {
			if setting.Id != 0 {
				// 更新
				_, err = dao.SiteSetting.Ctx(ctx).Data(g.Map{
					siteSettingCls.Type:   setting.Type,
					siteSettingCls.Value:  setting.Value,
					siteSettingCls.NameEn: setting.NameEn,
					siteSettingCls.NameZh: setting.NameZh,
				}).Where(siteSettingCls.Id, setting.Id).Update()
				if err != nil {
					return err
				}
			} else {
				// 新增
				_, err = dao.SiteSetting.Ctx(ctx).Data(g.Map{
					siteSettingCls.Type:   setting.Type,
					siteSettingCls.Value:  setting.Value,
					siteSettingCls.NameEn: setting.NameEn,
					siteSettingCls.NameZh: setting.NameZh,
				}).Insert()
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	return
}

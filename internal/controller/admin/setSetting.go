package admin

import (
	"blog-api/api/admin"
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

type cSetSetting struct{}

var SetSetting = cSetSetting{}

func (c *cSetSetting) GetSetSetting(ctx context.Context, req *admin.GetSiteSettingsReq) (res *admin.GetSiteSettingsRes, err error) {
	res = &admin.GetSiteSettingsRes{}
	setting1, err := service.SetSetting().GetSetSetting(ctx, "1")
	if err != nil {
		return nil, err
	}
	setting2, err := service.SetSetting().GetSetSetting(ctx, "2")
	if err != nil {
		return nil, err
	}
	setting3, err := service.SetSetting().GetSetSetting(ctx, "3")
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(setting1, &res.Type1)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(setting2, &res.Type2)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(setting3, &res.Type3)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cSetSetting) UpdateSetSetting(ctx context.Context, req *admin.UpdateSetSettingsReq) (res *admin.UpdateSetSettingsRes, err error) {
	var updateSettingInput = &model.UpdateSetSettingInput{}
	err = gconv.Scan(req.Settings, &updateSettingInput.Settings)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(req.DeleteIds, &updateSettingInput.DeleteIds)
	if err != nil {
		return nil, err
	}
	err = service.SetSetting().UpdateSetSetting(ctx, updateSettingInput)
	if err != nil {
		return nil, err
	}
	return
}

// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"blog-api/internal/model"
	"context"
)

type IAbout interface {
	GetAbout(ctx context.Context) (res *model.GetAboutOutput, err error)
	UpdateAbout(ctx context.Context, in *model.UpdateAboutInput) (err error)
}

var localAbout IAbout

func About() IAbout {
	if localAbout == nil {
		panic("implement not found for interface IAbout, forgot register?")
	}
	return localAbout
}

func RegisterAbout(i IAbout) {
	localAbout = i
}

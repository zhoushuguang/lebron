package logic

import (
	"context"

	"github.com/zhoushuguang/lebron/apps/cart/admin/internal/svc"
	"github.com/zhoushuguang/lebron/apps/cart/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLogic {
	return &AdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLogic) Admin(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}

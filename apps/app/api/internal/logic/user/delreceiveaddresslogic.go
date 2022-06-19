package user

import (
	"context"

	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelReceiveAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelReceiveAddressLogic {
	return &DelReceiveAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelReceiveAddressLogic) DelReceiveAddress(req *types.UserReceiveAddressDelReq) (resp *types.UserReceiveAddressDelRes, err error) {
	// todo: add your logic here and delete this line

	return
}

package user

import (
	"context"

	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddReceiveAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReceiveAddressLogic {
	return &AddReceiveAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddReceiveAddressLogic) AddReceiveAddress(req *types.UserReceiveAddressAddReq) (resp *types.UserReceiveAddressAddRes, err error) {

	return
}

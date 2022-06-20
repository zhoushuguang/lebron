package logic

import (
	"context"

	"github.com/zhoushuguang/lebron/apps/user/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserReceiveAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserReceiveAddressLogic {
	return &DelUserReceiveAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DelUserReceiveAddress 删除收获地址
func (l *DelUserReceiveAddressLogic) DelUserReceiveAddress(in *user.UserReceiveAddressDelReq) (*user.UserReceiveAddressDelRes, error) {
	// todo: add your logic here and delete this line

	return &user.UserReceiveAddressDelRes{}, nil
}

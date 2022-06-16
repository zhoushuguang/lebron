package user

import (
	"context"

	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRceiveAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRceiveAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRceiveAddressListLogic {
	return &UserRceiveAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRceiveAddressListLogic) UserRceiveAddressList(req *types.UserReceiveAddressListReq) (resp *types.UserReceiveAddressListRes, err error) {
	// todo: add your logic here and delete this line

	return
}

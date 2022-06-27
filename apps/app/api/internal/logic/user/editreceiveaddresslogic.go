package user

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditReceiveAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditReceiveAddressLogic {
	return &EditReceiveAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditReceiveAddressLogic) EditReceiveAddress(req *types.UserReceiveAddressEditReq) (resp *types.UserReceiveAddressEditRes, err error) {
	var editRpcReq user.UserReceiveAddressEditReq
	errCopy := copier.Copy(&editRpcReq, req)
	if errCopy != nil {
		return nil, errCopy
	}
	_, err = l.svcCtx.UserRPC.EditUserReceiveAddress(l.ctx, &editRpcReq)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

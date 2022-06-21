package logic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zhoushuguang/lebron/apps/user/model"
	"github.com/zhoushuguang/lebron/apps/user/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"
	"github.com/zhoushuguang/lebron/pkg/xerr"

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
	dbAddress := new(model.UserReceiveAddress)
	dbAddress.Id = in.Id
	dbAddress.IsDelete = 1
	err := l.svcCtx.UserReceiveAddressModel.Update(l.ctx, dbAddress)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "DelUserReceiveAddress Database Exception : %+v , err: %v", dbAddress, err)
	}
	return &user.UserReceiveAddressDelRes{}, nil
}

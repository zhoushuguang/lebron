package logic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhoushuguang/lebron/apps/user/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/user/rpc/model"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"
	"github.com/zhoushuguang/lebron/pkg/xerr"
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
	_, err := l.svcCtx.UserReceiveAddressModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errors.Wrap(xerr.NewErrMsg("数据不存在"), "收获地址不存在")
		}
		return nil, err
	}

	dbAddress := new(model.UserReceiveAddress)
	dbAddress.Id = in.GetId()
	dbAddress.IsDelete = 1
	err = l.svcCtx.UserReceiveAddressModel.UpdateIsDelete(l.ctx, dbAddress)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "DelUserReceiveAddress Database Exception : %+v , err: %v", dbAddress, err)
	}
	return &user.UserReceiveAddressDelRes{}, nil
}

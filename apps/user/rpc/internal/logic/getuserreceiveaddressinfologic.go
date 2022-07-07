package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhoushuguang/lebron/apps/user/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/user/rpc/model"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"
	"github.com/zhoushuguang/lebron/pkg/xerr"
)

type GetUserReceiveAddressInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserReceiveAddressInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserReceiveAddressInfoLogic {
	return &GetUserReceiveAddressInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据主键id,查询收获地址
func (l *GetUserReceiveAddressInfoLogic) GetUserReceiveAddressInfo(in *user.UserReceiveAddressInfoReq) (*user.UserReceiveAddress, error) {
	receiveAddress, err := l.svcCtx.UserReceiveAddressModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			//return nil, status.Error(100, "收获地址数据不存在")
			return nil, errors.Wrap(xerr.NewErrMsg("收获地址数据不存在"), "收获地址数据不存在")
		}
		return nil, err
	}
	var resp user.UserReceiveAddress
	_ = copier.Copy(&resp, receiveAddress)
	resp.CreateTime = receiveAddress.CreateTime.Unix()
	resp.UpdateTime = receiveAddress.UpdateTime.Unix()
	return &resp, nil
}

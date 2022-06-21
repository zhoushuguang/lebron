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

type EditUserReceiveAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditUserReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserReceiveAddressLogic {
	return &EditUserReceiveAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// EditUserReceiveAddress 编辑收获地址
func (l *EditUserReceiveAddressLogic) EditUserReceiveAddress(in *user.UserReceiveAddressEditReq) (*user.UserReceiveAddressEditRes, error) {
	//todo 限制,只能有一个默认地址
	if in.IsDefault == 1 {

	}

	dbAddress := new(model.UserReceiveAddress)
	dbAddress.Id = in.Id
	dbAddress.Name = in.GetName()
	dbAddress.Phone = in.GetPhone()
	dbAddress.Province = in.GetProvince()
	dbAddress.City = in.GetCity()
	dbAddress.IsDefault = in.GetIsDefault()
	dbAddress.PostCode = in.GetPostCode()
	dbAddress.Region = in.GetRegion()
	dbAddress.DetailAddress = in.GetDetailAddress()
	_, err := l.svcCtx.UserReceiveAddressModel.Insert(l.ctx, dbAddress)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "AddUserReceiveAddress Database Exception : %+v , err: %v", dbAddress, err)
	}
	return &user.UserReceiveAddressEditRes{}, nil
}

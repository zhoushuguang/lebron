package logic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zhoushuguang/lebron/apps/user/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/user/rpc/model"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"
	"github.com/zhoushuguang/lebron/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserReceiveAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserReceiveAddressLogic {
	return &AddUserReceiveAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddUserReceiveAddress 添加收获地址
func (l *AddUserReceiveAddressLogic) AddUserReceiveAddress(in *user.UserReceiveAddressAddReq) (*user.UserReceiveAddressAddRes, error) {
	dbAddress := new(model.UserReceiveAddress)
	dbAddress.Uid = in.GetUid()
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
	return &user.UserReceiveAddressAddRes{}, nil
}

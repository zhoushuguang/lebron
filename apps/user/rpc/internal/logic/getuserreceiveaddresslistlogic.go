package logic

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zhoushuguang/lebron/apps/user/model"
	"github.com/zhoushuguang/lebron/apps/user/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"
	"github.com/zhoushuguang/lebron/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserReceiveAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserReceiveAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserReceiveAddressListLogic {
	return &GetUserReceiveAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取收获地址列表
func (l *GetUserReceiveAddressListLogic) GetUserReceiveAddressList(in *user.UserReceiveAddressListReq) (*user.UserReceiveAddressListRes, error) {
	whereBuilder := l.svcCtx.UserReceiveAddressModel.RowBuilder().Where(squirrel.Eq{"uid": in.Uid})
	receiveAddressesList, err := l.svcCtx.UserReceiveAddressModel.FindAll(l.ctx, whereBuilder)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Failed to get user's homestay order err : %v , in :%+v", err, in)
	}
	var resp []*user.UserReceiveAddress
	if len(receiveAddressesList) > 0 {
		for _, receiveAddresses := range receiveAddressesList {
			var pbAddress user.UserReceiveAddress
			_ = copier.Copy(&pbAddress, receiveAddresses)

			//pbHomestayOrder.CreateTime = homestayOrder.CreateTime.Unix()
			//pbHomestayOrder.LiveStartDate = homestayOrder.LiveStartDate.Unix()
			//pbHomestayOrder.LiveEndDate = homestayOrder.LiveEndDate.Unix()

			resp = append(resp, &pbAddress)
		}
	}
	return &user.UserReceiveAddressListRes{
		List: resp,
	}, nil
}

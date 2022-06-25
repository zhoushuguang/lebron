package user

import (
	"context"
	"encoding/json"

	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"
	"github.com/zhoushuguang/lebron/pkg/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserReceiveAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserReceiveAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserReceiveAddressListLogic {
	return &UserReceiveAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserReceiveAddressListLogic) UserReceiveAddressList(req *types.UserReceiveAddressListReq) (resp *types.UserReceiveAddressListRes, err error) {
	var addressListReq user.UserReceiveAddressListReq
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Error! get uid from token"), "Failed toget uid from token err : %v ,req:%+v", err, req)
	}
	addressListReq.Uid = uid
	rpcRes, err := l.svcCtx.UserRPC.GetUserReceiveAddressList(l.ctx, &addressListReq)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Error! Function UserReceiveAddressList"), "Failed to get user addrerss  list err : %v ,req:%+v", err, req)
	}
	var addressList []types.UserReceiveAddress
	for _, rpcAddress := range rpcRes.List {
		var addressVo types.UserReceiveAddress
		_ = copier.Copy(&addressVo, rpcAddress)
		addressList = append(addressList, addressVo)
	}
	return &types.UserReceiveAddressListRes{List: addressList}, nil
}

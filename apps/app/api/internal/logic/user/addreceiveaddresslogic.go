package user

import (
	"context"
	"encoding/json"

	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddReceiveAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddReceiveAddressLogic {
	return &AddReceiveAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddReceiveAddressLogic) AddReceiveAddress(req *types.UserReceiveAddressAddReq) (resp *types.UserReceiveAddressAddRes, err error) {
	//l.ctx.Value("uid")  get jwt token uid
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	_, err = l.svcCtx.UserRPC.AddUserReceiveAddress(l.ctx, &user.UserReceiveAddressAddReq{
		Uid:           int32(uid),
		Name:          req.Name,
		Phone:         req.Phone,
		Province:      req.Province,
		City:          req.City,
		IsDefault:     req.IsDefault,
		PostCode:      req.PostCode,
		Region:        req.Region,
		DetailAddress: req.DetailAddress,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

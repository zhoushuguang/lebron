package user

import (
	"context"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"

	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCollectionDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCollectionDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCollectionDelLogic {
	return &UserCollectionDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCollectionDelLogic) UserCollectionDel(req *types.UserCollectionDelReq) (resp *types.UserCollectionDelRes, err error) {
	_, err = l.svcCtx.UserRPC.DelUserCollection(l.ctx, &user.UserCollectionDelReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

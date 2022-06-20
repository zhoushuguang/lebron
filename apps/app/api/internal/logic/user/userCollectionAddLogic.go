package user

import (
	"context"

	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCollectionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCollectionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCollectionAddLogic {
	return &UserCollectionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCollectionAddLogic) UserCollectionAdd(req *types.CollectionAddReq) (resp *types.CollectionAddRes, err error) {
	res, err := l.svcCtx.UserRPC.AddCollection(l.ctx, &user.CollectionAddReq{Collection: &user.Collection{
		Uid:       req.Collection.Uid,
		ProductId: req.Collection.ProductId,
	}})
	if err != nil {
		return nil, err
	}
	return &types.CollectionAddRes{
		Code:    int16(res.Code),
		Message: res.Message,
	}, nil
}

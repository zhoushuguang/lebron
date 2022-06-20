package logic

import (
	"context"

	"github.com/zhoushuguang/lebron/apps/user/model"
	"github.com/zhoushuguang/lebron/apps/user/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollectionLogic {
	return &AddCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  添加收藏
func (l *AddCollectionLogic) AddCollection(in *user.CollectionAddReq) (*user.CollectionAddRes, error) {
	collectionIn := new(model.Collection)
	collectionIn.Uid = in.Collection.Uid
	collectionIn.ProductId = in.Collection.ProductId
	_, err := l.svcCtx.CollectionModel.Insert(l.ctx, collectionIn)
	if err != nil {
		return nil, err
	}
	return &user.CollectionAddRes{
		Code:    200,
		Message: "操作成功",
	}, nil
}

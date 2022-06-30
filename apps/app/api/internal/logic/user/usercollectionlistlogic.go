package user

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"
	"github.com/zhoushuguang/lebron/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCollectionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCollectionListLogic {
	return &UserCollectionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCollectionListLogic) UserCollectionList(req *types.UserCollectionListReq) (resp *types.UserCollectionListRes, err error) {
	var collectionListReq user.UserCollectionListReq
	uid, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Error! get uid from token"), "Failed toget uid from token err : %v ,req:%+v", err, req)
	}
	collectionListReq.Uid = uid
	rpcRes, err := l.svcCtx.UserRPC.GetUserCollectionList(l.ctx, &collectionListReq)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Error! Function UserCollectionList"), "Failed to get user UserCollectionList  err : %v ,req:%+v", err, req)
	}
	var productId1 []int64
	for _, rpcList := range rpcRes.ProductId {
		productId1 = append(productId1, rpcList)
	}

	return &types.UserCollectionListRes{
		ProductId: productId1,
	}, nil
}

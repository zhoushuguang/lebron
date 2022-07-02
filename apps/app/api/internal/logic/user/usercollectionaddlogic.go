package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"

	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"

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

func (l *UserCollectionAddLogic) UserCollectionAdd(req *types.UserCollectionAddReq) (resp *types.UserCollectionAddRes, err error) {
	//get jwt token uid
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	var addRpcReq user.UserCollectionAddReq
	fmt.Println("uid ---", uid)
	addRpcReq.Uid = uid
	errCopy := copier.Copy(&addRpcReq, req)
	if err != nil {
		return nil, errCopy
	}
	_, err = l.svcCtx.UserRPC.AddUserCollection(l.ctx, &addRpcReq)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

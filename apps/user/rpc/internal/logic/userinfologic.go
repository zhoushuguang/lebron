package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zhoushuguang/lebron/apps/user/model"
	"google.golang.org/grpc/status"

	"github.com/zhoushuguang/lebron/apps/user/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	umsMember, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	var resp user.UserInfoResponse
	_ = copier.Copy(&resp, umsMember)

	return &resp, nil
}

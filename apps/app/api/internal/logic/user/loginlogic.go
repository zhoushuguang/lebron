package user

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"
	"github.com/zhoushuguang/lebron/pkg/jwtx"

	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	var loginReq user.LoginRequest
	loginReq.Username = req.Username
	loginReq.Password = req.Password
	res, err := l.svcCtx.UserRPC.Login(l.ctx, &loginReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}

	//generate token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}

	return &types.LoginResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
	}, nil
}

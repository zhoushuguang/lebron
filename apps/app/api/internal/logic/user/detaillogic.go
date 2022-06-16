package user

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"

	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	//l.ctx.Value("uid")  get jwt token uid
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	userInfo, err := l.svcCtx.UserRPC.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.UserInfoResp)
	_ = copier.Copy(&resp, userInfo)
	return resp, nil
}

package logic

import (
	"context"

	"github.com/zhoushuguang/lebron/apps/reply/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/reply/rpc/reply"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentsLogic {
	return &CommentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentsLogic) Comments(in *reply.CommentsRequest) (*reply.CommentsResponse, error) {
	// todo: add your logic here and delete this line

	return &reply.CommentsResponse{}, nil
}

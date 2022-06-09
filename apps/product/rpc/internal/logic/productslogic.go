package logic

import (
	"context"

	"github.com/zhoushuguang/lebron/apps/product/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.ProductRequest) (*product.ProductResponse, error) {
	if in.ProductIds == "1" {
		products := make(map[int64]*product.ProductItem)
		products[1] = &product.ProductItem{
			ProductId: 1,
			Name:      "测试商品名称",
		}
		return &product.ProductResponse{Products: products}, nil
	}
	return &product.ProductResponse{}, nil
}

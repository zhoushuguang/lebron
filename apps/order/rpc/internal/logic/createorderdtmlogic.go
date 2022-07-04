package logic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhoushuguang/lebron/apps/order/rpc/internal/svc"
	"github.com/zhoushuguang/lebron/apps/order/rpc/model"
	"github.com/zhoushuguang/lebron/apps/order/rpc/order"
	"github.com/zhoushuguang/lebron/apps/product/rpc/product"
	"github.com/zhoushuguang/lebron/apps/user/rpc/user"
	"github.com/zhoushuguang/lebron/pkg/snowflake"
	"github.com/zhoushuguang/lebron/pkg/xerr"
)

type CreateOrderDTMLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderDTMLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderDTMLogic {
	return &CreateOrderDTMLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderDTMLogic) CreateOrderDTM(in *order.AddOrderReq) (*order.AddOrderResp, error) {
	//check user
	var userReq user.UserInfoRequest
	userReq.Id = in.Userid
	userRpcRes, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userReq)
	if err != nil {
		return nil, err
	}
	if userRpcRes == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DataNoExistError), "Error! 用户不存在 Exception : %+v  ", userRpcRes)
	}

	//check product
	var productReq product.ProductItemRequest
	productReq.ProductId = in.Productid
	productRpcRes, err := l.svcCtx.ProductRpc.Product(l.ctx, &productReq)
	if err != nil {
		return nil, err
	}
	if productRpcRes == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DataNoExistError), "Error! 商品不存在 Exception : %+v  ", productRpcRes)
	}

	//check product stock
	if productRpcRes.Stock <= 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("商品库存不足"), "商品库存不足")
	}

	//check user_receive_address
	var userReceiveAddressInfoReq user.UserReceiveAddressInfoReq
	userReceiveAddressInfoReq.Id = in.ReceiveAddressId
	receiveAddress, err := l.svcCtx.UserRpc.GetUserReceiveAddressInfo(l.ctx, &userReceiveAddressInfoReq)
	if err != nil {
		return nil, err
	}
	if receiveAddress == nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DataNoExistError), "Error! 用户收获地址 Exception : %+v  ", receiveAddress)
	}

	//generate new order id
	orderId := snowflake.GenIDString()

	//create new shipping
	var insertShipping model.Shipping
	insertShipping.Orderid = orderId
	insertShipping.Userid = in.Userid
	insertShipping.ReceiverName = receiveAddress.Name
	insertShipping.ReceiverPhone = receiveAddress.Phone
	insertShipping.ReceiverMobile = receiveAddress.Phone
	insertShipping.ReceiverProvince = receiveAddress.Province
	insertShipping.ReceiverCity = receiveAddress.City
	insertShipping.ReceiverDistrict = receiveAddress.Region
	insertShipping.ReceiverAddress = receiveAddress.DetailAddress
	insertShippingRes, err := l.svcCtx.ShippingModel.Insert(l.ctx, &insertShipping)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "create new shipping Database Exception : %+v , err: %v", insertShipping, err)
	}
	newShippingId, err := insertShippingRes.LastInsertId()
	if err != nil {
		return nil, err
	}
	logx.Info("New Shipping Primary Id:", newShippingId)

	//create new orderitem
	insertOrderitem := model.Orderitem{
		OrderId:      orderId,
		UserId:       in.Userid,
		ProductId:    in.Productid,
		ProductName:  productRpcRes.Name,
		ProductImage: productRpcRes.ImageUrl,
		CurrentPrice: productRpcRes.Price,
		Quantity:     in.Quantity,
		//TotalPrice:   float64(mathin.Quantity * productRpcRes.GetPrice()),
	}
	insertOrderitemRes, err := l.svcCtx.OrderitemModel.Insert(l.ctx, &insertOrderitem)
	if err != nil {
		return nil, err
	}
	newOrderitemId, err := insertOrderitemRes.LastInsertId()
	if err != nil {
		return nil, err
	}
	logx.Info("New newOrderitem Primary Id:", newOrderitemId)

	//create new order
	insertOrder := model.Orders{
		Id:         orderId,
		Userid:     in.Userid,
		Shoppingid: newShippingId,
		Postage:    in.Postage,
	}
	_, err = l.svcCtx.OrderModel.Insert(l.ctx, &insertOrder)
	if err != nil {
		return nil, err
	}

	//update product stock
	updateProductStockRequest := product.UpdateProductStockRequest{
		ProductId: in.Productid,
		Num:       in.Quantity,
	}
	_, err = l.svcCtx.ProductRpc.UpdateProductStock(l.ctx, &updateProductStockRequest)
	if err != nil {
		return nil, err
	}
	return &order.AddOrderResp{
		Id: orderId,
	}, nil
}

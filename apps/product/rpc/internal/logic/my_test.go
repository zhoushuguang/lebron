package logic

import (
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zhoushuguang/lebron/apps/product/rpc/internal/model"
	"github.com/zhoushuguang/lebron/apps/product/rpc/product"
	"testing"
)

func TestName(t *testing.T) {
	products := make(map[int64]*product.ProductItem)
	pdis := []int64{1}
	ps, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, pid := range pdis {
			source <- pid
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		pid := item.(int64)
		println(pid)

		p := &model.Product{
			Id: 1,
		}
		writer.Write(p)
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		var r []*model.Product
		for p := range pipe {
			r = append(r, p.(*model.Product))
		}
		writer.Write(r)
	})
	if err != nil {
		println(err)
	}

	for _, p := range ps.([]*model.Product) {
		products[p.Id] = &product.ProductItem{
			ProductId: p.Id,
		}
	}
}

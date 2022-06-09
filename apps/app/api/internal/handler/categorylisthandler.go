package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/logic"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/svc"
	"github.com/zhoushuguang/lebron/apps/app/api/internal/types"
)

func CategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.CategoryList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

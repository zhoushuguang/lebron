package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zhoushuguang/lebron/apps/product/admin/internal/logic"
	"github.com/zhoushuguang/lebron/apps/product/admin/internal/svc"
)

func UploadImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadImageLogic(r.Context(), svcCtx, r)
		resp, err := l.UploadImage()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

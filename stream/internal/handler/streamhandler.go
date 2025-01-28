package handler

import (
	"net/http"

	"gateway-zero-ex/stream/internal/logic"
	"gateway-zero-ex/stream/internal/svc"
	"gateway-zero-ex/stream/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func StreamHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewStreamLogic(r.Context(), svcCtx)
		resp, err := l.Stream(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

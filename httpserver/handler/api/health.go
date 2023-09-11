package api

import (
	"github.com/inazak/training-go-httpserver/common/jsonhelper"
	"net/http"
)

func (h *Handler) ServeHealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := h.backend.HealthCheck(ctx)
	if err != nil {
		_ = jsonhelper.WriteJSONResponse(
			ctx,
			w,
			&jsonhelper.ErrorResponse{Message: err.Error()},
			http.StatusInternalServerError)
		return
	}

	rsp := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}
	_ = jsonhelper.WriteJSONResponse(ctx, w, rsp, http.StatusOK)
}

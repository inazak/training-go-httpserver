package api

import (
  "net/http"
  "github.com/inazak/training-go-httpserver/common/jsonhelper"
)

func (h *Handler) ServeHealthCheck(w http.ResponseWriter, r *http.Request) {
  ctx := r.Context()

  err := h.backend.HealthCheck(ctx)
  if err != nil {
    _ = jsonhelper.WriteJSONResponse(
      ctx,
      w,
      &jsonhelper.ErrorResponse{ Message: err.Error(), },
      http.StatusInternalServerError)
    return
  }

  rsp := struct{
    status string `json:"status"`
  }{
    status: "ok",
  }
  _ = jsonhelper.WriteJSONResponse(ctx, w ,rsp, http.StatusOK)
}


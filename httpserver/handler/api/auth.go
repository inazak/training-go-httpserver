package api

import (
	"context"
	"github.com/inazak/training-go-httpserver/common/jsonhelper"
	"net/http"
)

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, err := h.backend.ValidateToken(ctx, r)
		if err != nil {
			_ = jsonhelper.WriteJSONResponse(
				ctx,
				w,
				&jsonhelper.ErrorResponse{Message: "fail to authorize"},
				http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(ctx, contextKeyUserID, id)
		r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

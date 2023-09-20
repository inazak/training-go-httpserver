package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/inazak/training-go-httpserver/service"
)

type contextKey string

var (
	contextKeyUserID contextKey = "userid"
)

var (
	validation = validator.New()
)

type Handler struct {
	backend service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{
		backend: svc,
	}
}

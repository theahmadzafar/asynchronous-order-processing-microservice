package handlers

import (
	"asynchronous-order-processing-microservice/internal/entities"
	"asynchronous-order-processing-microservice/internal/services"
	"asynchronous-order-processing-microservice/internal/transport/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	os *services.OrderService
}

func NewOrderHandler(os *services.OrderService) http.Handler {
	return &OrderHandler{
		os: os,
	}
}

func (h *OrderHandler) Register(route *gin.RouterGroup) {
	route.POST("submitOrder", h.submitOrder)
}

func (h *OrderHandler) submitOrder(ctx *gin.Context) {
	req := entities.Order{}
	if err := ctx.ShouldBind(&req); err != nil {
		http.Error(ctx, nil)
		return
	}
	if err := h.os.Create(ctx, req); err != nil {
		http.Error(ctx, nil)
		return
	}

	http.OK(ctx, HealthResponse{Success: "ok"})
}

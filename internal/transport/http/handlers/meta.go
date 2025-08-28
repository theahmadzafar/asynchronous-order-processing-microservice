package handlers

import (
	"asynchronous-order-processing-microservice/internal/transport/http"

	"github.com/gin-gonic/gin"
)

var tag = "no tag"

type HealthResponse struct {
	Success string `json:"success"`
}

type InfoResponse struct {
	Tag string `json:"tag"`
}

type MetaHandler struct {
}

func NewMetaHandler() http.Handler {
	return &MetaHandler{}
}

func (h *MetaHandler) Register(route *gin.RouterGroup) {
	route.GET("health", h.health)
	route.GET("info", h.info)
}

func (h *MetaHandler) health(ctx *gin.Context) {
	http.OK(ctx, HealthResponse{Success: "ok"})
}

func (h *MetaHandler) info(ctx *gin.Context) {
	http.OK(ctx, InfoResponse{Tag: tag})
}

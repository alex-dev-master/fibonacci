package handler

import (
	"github.com/alex-dev-master/fibonacci.git/intrernal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/get-fibonacci", h.getFibonacci)

	}

	return router
}
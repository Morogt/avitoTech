package handler

import (
	"avitoTech/pkg/service"
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

	balance := router.Group("/balance")
	{
		balance.GET("/info", h.info)
		balance.POST("/addFunds", h.addMoney)
		balance.POST("/withdrawalOfMoney", h.withdrawalOfMoney)
		balance.POST("/transfer", h.transfer)
		balance.POST("/pay", h.pay)
		balance.GET("/history", h.history)
	}

	//services := router.Group("/services")
	//{
	//	services.POST("/addService", h.addService)
	//}

	reports := router.Group("/reports")
	{
		reports.GET("/get", h.get)
	}

	user := router.Group("/user", h.create)
	{
		user.POST("/create")
	}

	return router
}

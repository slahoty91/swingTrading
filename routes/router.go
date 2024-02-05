package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/slahoty91/tradingBot/controller"
)

func InitilizeRoutes(router *gin.Engine) {
	nseController := controller.NewNSEController()
	ordCont := controller.NewOrdersController()
	api := router.Group("/api")

	api.GET("/getName/:inst_token", nseController.GetName)
	api.GET("/serachInstrument/:query", nseController.FuzzySearch)
	api.POST("/placeOrder", ordCont.InsertOrder)
}

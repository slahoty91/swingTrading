package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/slahoty91/tradingBot/model"
	"github.com/slahoty91/tradingBot/services"
)

type NSEController struct {
	nseService services.NSEservice
}

func NewNSEController() *NSEController {
	return &NSEController{
		nseService: *services.NewNSEService(),
	}
}

type OrdersController struct {
	ordersService services.Orderservice
}

func NewOrdersController() *OrdersController {
	return &OrdersController{
		ordersService: *services.NewOrderService(),
	}
}

func (nc *NSEController) GetName(c *gin.Context) {
	ins_token := c.Param("inst_token")
	token, err := strconv.ParseInt(ins_token, 10, 64)
	if err != nil {
		fmt.Println("Error converting string to int64:", err)
		return
	}
	name, err := nc.nseService.GetName(token)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"name": name})
}

func (nc *NSEController) FuzzySearch(c *gin.Context) {
	query := c.Param("query")
	fmt.Println(query, "from controller")
	results, err := nc.nseService.FuzzySearch(query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"results": results})
}

func (oc *OrdersController) InsertOrder(c *gin.Context) {
	body := c.Request.Body
	fmt.Println(body, "from controller")
	var order model.Orders
	if err := c.BindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON format"})
		return
	}
	err := oc.ordersService.CreateOrders(order)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"results": order})
}

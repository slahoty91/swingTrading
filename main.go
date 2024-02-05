package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/slahoty91/tradingBot/routes"
)

func main() {
	fmt.Println("THIS is a test")
	r := gin.Default()
	routes.InitilizeRoutes(r)
	r.Run(":8000")
}

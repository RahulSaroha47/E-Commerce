package main

import (
	"e-commerce/middleware"
	"e-commerce/routes"
	"log"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port ="8080"
	}

	app := controllers.NewApplication()

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart",app.AddToCart())
	router.GET("/removeitem",app.RemoveItem())
	router.GET("/cartcheckout",app.BuyFromCart())
	router.GET("/instantbuy",app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
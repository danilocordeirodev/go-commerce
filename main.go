package main

import (
	"log"
	"os"

	"github.com/danilocordeirodev/go-commerce/controllers"
	"github.com/danilocordeirodev/go-commerce/database"
	"github.com/danilocordeirodev/go-commerce/middleware"
	"github.com/danilocordeirodev/go-commerce/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Cliente, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("cartcheckout", app.BuyFromCart())
	router.GET("instanbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}

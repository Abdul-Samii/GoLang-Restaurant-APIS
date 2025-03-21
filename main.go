package main

import (
	"golang-restuarant-management/database"
	"golang-restuarant-management/middlewares"
	"golang-restuarant-management/routes"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.UserRouter(routes.Routes)
	router.Use(middlewares.Authentication())

	routes.foodRoutes(router)
	routes.orderRoutes(router)
	routes.userRoutes(router)
	routes.tableRoutes(router)
	routes.invoiceRoutes(router)
	routes.menuRoutes(router)
	routes.orderItemRoutes(router)

	router.Run(":" + port)
}

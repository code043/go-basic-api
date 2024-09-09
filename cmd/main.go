package main

import (
	"github.com/code043/go-basic-api/controllers"
	"github.com/code043/go-basic-api/database"
	"github.com/code043/go-basic-api/repositories"
	"github.com/code043/go-basic-api/usecases"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(cors.Default())

	//server.GET("/ping", ping.Ping)
	dbConnection, err := database.ConnectDB()
	if err != nil{
		panic(err)
	}
	productRepository := repositories.NewProductRepository(dbConnection)
	productUseCase := usecases.NewProductUsecase(productRepository)
	productController := controllers.NewProductController(productUseCase)

	server.GET("/products", productController.GetProducts)
	server.POST("/products", productController.CreateProduct)
	server.GET("/products/:productId", productController.GetProductById)

	server.Run(":8000")
}
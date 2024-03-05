package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/valentineejk/gomongo/internal/database"
	"github.com/valentineejk/gomongo/internal/handler"
)

func main() {
	//load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGO_URI")

	err := database.ConnectDb(uri, "mako")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("..")
	fmt.Println(".....")
	fmt.Println("........")
	fmt.Println("connected to db")

	defer func() {
		err := database.CloseDb()
		if err != nil {
			fmt.Println(err)
		}
	}()

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	gin.Logger()

	r.GET("/products", handler.GetProducts)
	r.POST("/products", handler.AddProducts)
	r.GET("/products/:id", handler.GetProductById)
	r.PATCH("/products/:id/stock", handler.UpdateProductStockById)
	r.PATCH("/products/:id/price", handler.UpdateProductPriceById)
	r.DELETE("/products/:id", handler.DeleteProductById)

	r.Run(":8080")

}

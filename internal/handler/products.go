package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valentineejk/gomongo/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProducts(c *gin.Context) {
	cursor, err := database.Products.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	//products
	var products []database.Product
	if err = cursor.All(c, &products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "unable to fetch products",
		})
		return
	}

	//ok
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   products,
	})
}

func AddProducts(c *gin.Context) {

	var body database.CreateProductRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid request body",
		})
		return
	}

	res, err := database.Products.InsertOne(c, body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "error inserting body",
		})
		return
	}

	product := database.Product{
		ID:       res.InsertedID.(primitive.ObjectID),
		Name:     body.Name,
		Category: body.Category,
		Stock:    body.Stock,
		Price:    body.Price,
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data":   product,
	})
}

func GetProductById(c *gin.Context) {

	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid id",
		})
		return
	}

	res := database.Products.FindOne(c, primitive.M{"_id": _id})
	product := database.Product{}
	err = res.Decode(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "unable to find product",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": product,
	})
}

func UpdateProductStockById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid id",
		})
		return
	}

	var body struct {
		Stock int `json:"stock" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid request body",
		})
		return
	}

	_, err = database.Products.UpdateOne(c, bson.M{"_id": _id}, bson.M{"$set": bson.M{"stock": body.Stock}})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "error updating stock",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "stock updated",
	})
}

func UpdateProductPriceById(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid id",
		})
		return
	}

	var body struct {
		Price float32 `json:"price" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid request body",
		})
		return
	}

	_, err = database.Products.UpdateOne(c, bson.M{"_id": _id}, bson.M{"$set": bson.M{"price": body.Price}})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "error updating stock",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "price updated",
	})
}

func DeleteProductById(c *gin.Context) {
	id := c.Param("id")

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid id",
		})
		return
	}

	res, err := database.Products.DeleteOne(c, bson.M{"_id": _id})
	if res.DeletedCount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "product found",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "product deleted",
	})

}

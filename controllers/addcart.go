// add book to cart with session and inner join
package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"web/utils/token"
	"web/models"
)
type Cart struct {
	ID uint `gorm:"primary_key" json:"id"`
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
	Quantity uint `json:"quantity"`
}
// add book to cart with session and inner join
type AddCartInput struct {
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
	Quantity uint `json:"quantity"`
}

func AddCart(c *gin.Context) {
	//extract token
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var input AddCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//check if book is already in cart
	var cart models.Cart
	if err := models.DB.Where("user_id = ? AND book_id = ?", user_id, input.BookID).First(&cart).Error; err != nil {
		//if not in cart, add it
		cart := models.Cart{
			UserID: user_id,
			BookID: input.BookID,
			Quantity: input.Quantity,
		}
		if err := models.DB.Create(&cart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book added to cart"})
		return
	}
	//if already in cart, update quantity
	if err := models.DB.Model(&cart).Where("user_id = ? AND book_id = ?", user_id, input.BookID).Update("quantity", cart.Quantity + input.Quantity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book quantity updated"})
}
// view cart
func ViewCart(c *gin.Context) {
	//extract token
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var cart []models.Cart
	if err := models.DB.Where("user_id = ?", user_id).Find(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": cart})
}
package handler

import (
	"log"
	"product-service/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase *usecase.ProductUsecase
}

func NewProductHandler(u *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		usecase: u,
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.usecase.GetProducts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"data": products})
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Failed casting string id to int")
	}

	product, err := h.usecase.GetByID(uint(idInt))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{"data": product})
}

package handler

import (
	"ecommerce-go/internal/core/domain"
	"ecommerce-go/internal/core/ports"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionHandler struct {
	service ports.EcommerceService
}

func NewTransactionHandler(service ports.EcommerceService) *TransactionHandler {

	return &TransactionHandler{
		service: service,
	}
}

// GetProducts godoc
// @Summary Obtain all products
// @Description This endpoint gets all the products
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {array} domain.Product
// @Router /products [get]
// @Security Bearer
func (th *TransactionHandler) GetProducts(ctx *gin.Context) {
	products, err := th.service.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"products": products})
}

// PostProduct godoc
// @Summary Create a product
// @Description Endpoint for loading a new product
// @Tags example
// @Accept json
// @Produce json
// @Success 200
// @Param product body domain.Product true "Product data"
// @Router /products [post]
// @Security Bearer
func (th *TransactionHandler) PostProduct(ctx *gin.Context) {
	var newProduct domain.Product
	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := th.service.PostProduct(newProduct)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, gin.H{"product": nil})
}

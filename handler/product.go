package handler

import (
	"go-jwt/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//ProductHandler --> interface to Product handler
type ProductHandler interface {
	GetProduct(*gin.Context)
	GetAllProducts(*gin.Context)
}

type productHandler struct {
	repo repository.ProductRepository
}

//NewProductHandler --> returns new handler for product entity
func NewProductHandler() ProductHandler {
	return &productHandler{
		repo: repository.NewProductRepository(),
	}
}


func(h *productHandler) GetProduct(ctx *gin.Context){
	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":  err.Error()})
		return
	}
	product, _ := h.repo.GetProduct(ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (h * productHandler) GetAllProducts(ctx *gin.Context) {
	product, err := h.repo.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}
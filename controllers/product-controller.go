package controllers

import (
	"net/http"
	"strconv"

	"github.com/code043/go-basic-api/models"
	"github.com/code043/go-basic-api/usecases"
	"github.com/gin-gonic/gin"
)

type ProductController struct{
	productUseCase usecases.ProductUsecase
}

func NewProductController(usecase usecases.ProductUsecase) ProductController{
	return ProductController{
		productUseCase: usecase,
	}
}
func (p *ProductController) GetProducts(ctx *gin.Context){
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}
func (p *ProductController) CreateProduct(ctx *gin.Context){
	var product models.Product
	err := ctx.BindJSON(&product)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, insertedProduct)
}
func (p *ProductController) GetProductById(ctx *gin.Context){
	id := ctx.Param("productId")
	if id == ""{
		response := models.Response{
			Message: "Product ID cannot be null!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	producId, err := strconv.Atoi(id)
	if err != nil {
		response := models.Response{
			Message: "Product ID must be a number!",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	product, err := p.productUseCase.GetProductById(producId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	if product == nil {
		response := models.Response{
			Message: "Product was not found in the database!",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	ctx.JSON(http.StatusOK, product)
}
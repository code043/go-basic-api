package usecases

import (
	"fmt"

	"github.com/code043/go-basic-api/models"
	"github.com/code043/go-basic-api/repositories"
)
type ProductUsecase struct{
	repository repositories.ProductRepository
}

func NewProductUsecase(repo repositories.ProductRepository) ProductUsecase{

	return ProductUsecase{
		repository: repo,
	}
}
func (pU *ProductUsecase) GetProducts() ([]models.Product, error){
	return pU.repository.GetProducts()
}
func (pU *ProductUsecase) CreateProduct(product models.Product) (models.Product, error){
	productId, err := pU.repository.CreateProduct(product)
	if err != nil{
		fmt.Println(err)
		return models.Product{}, err
	}
	product.ID = productId
	return product, nil
}
func (pu *ProductUsecase) GetProductById(id_product int)(*models.Product, error){
	product, err := pu.repository.GetProductById(id_product)
	if err != nil{
		return nil, nil
	}
	return product, nil
}
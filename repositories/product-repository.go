package repositories

import (
	"database/sql"
	"fmt"

	"github.com/code043/go-basic-api/models"
)

type ProductRepository struct{
	connection *sql.DB
}
func NewProductRepository(connection *sql.DB) ProductRepository{
	return ProductRepository{
		connection: connection,
	}
}
func (pR *ProductRepository) GetProducts() ([]models.Product, error){
	query := "select id, product_name, price from product"
	rows, err := pR.connection.Query(query)
	if err != nil{
		fmt.Println(err)
		return []models.Product{}, err
	}
	var productsList []models.Product 
	var productObj models.Product

	for rows.Next(){
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price,
		)
		if err != nil {
			fmt.Println(err)
			return []models.Product{}, err
		}
		productsList = append(productsList, productObj)
	}
	rows.Close()
	return productsList, nil
}
func (pR *ProductRepository) CreateProduct(product models.Product) (int, error){
	var id int
	query, err := pR.connection.Prepare("insert into product (product_name, price) values ($1, $2) returning id")
	if err != nil{
		fmt.Println(err)
		return 0, err
	}
	err = query.QueryRow(product.Name, product.Price).Scan(
		&id,
	)
	if err != nil{
		fmt.Println(err)
		return 0, err
	}
	query.Close()
	return id, nil
}
func (pR *ProductRepository) GetProductById(id_product int) (*models.Product, error){
	query, err := pR.connection.Prepare("select * from product where id = $1")
	if err != nil{
		fmt.Println(err)
		return nil, err
	}
	var product models.Product
	err = query.QueryRow(id_product).Scan(
		&product.ID,
		&product.Name,
		&product.Price,
	)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, nil
		}
		return nil, err
	}
	query.Close()
	return &product, nil
}
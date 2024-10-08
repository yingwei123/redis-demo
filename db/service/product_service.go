package service

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"redis-demo/db/model"
	"redis-demo/db/repository"
	"strconv"
	"time"
)

type ProductService interface {
	CreateProduct(product model.Product) (model.Product, error)
	GetProduct(id uint) (model.Product, error)
	UpdateProduct(id uint, toUpdate map[string]interface{}) error
	DeleteProduct(id uint) error
	GetAllProducts() ([]model.Product, error)
	ReadCSVToProducts(csvPath string) ([]model.Product, error)
	GetAllProductsWithRedis() ([]model.Product, error)
	GetProductWithRedis(id uint) (model.Product, error)
	DeleteProductWithRedis(id uint) error
	UpdateProductWithRedis(id uint, toUpdate map[string]interface{}) error
}

type ProductServiceImpl struct {
	ProductRepo repository.ProductRepository
}

func NewProductService(ProductRepo repository.ProductRepository) ProductService {
	return &ProductServiceImpl{ProductRepo: ProductRepo}
}

func (u *ProductServiceImpl) CreateProduct(product model.Product) (model.Product, error) {
	if product.Name == "" || product.Price == 0 || product.Stock == 0 {
		return model.Product{}, errors.New("all fields are required")
	}

	return u.ProductRepo.CreateProduct(product)
}

func (u *ProductServiceImpl) GetProduct(id uint) (model.Product, error) {
	if id == 0 {
		return model.Product{}, errors.New("id is required")
	}

	return u.ProductRepo.GetProduct(id)
}

func (u *ProductServiceImpl) UpdateProduct(id uint, toUpdate map[string]interface{}) error {
	if toUpdate["Name"] == "" && toUpdate["Price"] == 0 && toUpdate["Description"] == "" {
		return errors.New("no fields to update")
	}

	return u.ProductRepo.UpdateProduct(id, toUpdate)
}

func (u *ProductServiceImpl) DeleteProduct(id uint) error {
	if id == 0 {
		return errors.New("id is required")
	}

	return u.ProductRepo.DeleteProduct(id)
}

func (u *ProductServiceImpl) GetAllProducts() ([]model.Product, error) {
	return u.ProductRepo.GetAllProducts()
}

func (u *ProductServiceImpl) GetAllProductsWithRedis() ([]model.Product, error) {
	return u.ProductRepo.GetAllProductsWithRedis()
}

func (u *ProductServiceImpl) GetProductWithRedis(id uint) (model.Product, error) {
	return u.ProductRepo.GetProductWithRedis(id)
}

func (u *ProductServiceImpl) DeleteProductWithRedis(id uint) error {
	return u.ProductRepo.DeleteProductWithRedis(id)
}

func (u *ProductServiceImpl) UpdateProductWithRedis(id uint, toUpdate map[string]interface{}) error {
	return u.ProductRepo.UpdateProductWithRedis(id, toUpdate)
}

// ReadCSVToProducts reads a CSV file and converts each row into a Product model
func (u *ProductServiceImpl) ReadCSVToProducts(csvPath string) ([]model.Product, error) {
	// Open the CSV file
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close()

	// Initialize CSV reader
	reader := csv.NewReader(file)

	// Read the CSV header (skipping the first line)
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %v", err)
	}

	var products []model.Product

	// Iterate through the rows
	for {
		// Read each row
		record, err := reader.Read()
		if err != nil {
			// End of file is reached
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("failed to read CSV row: %v", err)
		}

		// Convert each field from the CSV row to the appropriate type
		price, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid price value: %v", err)
		}

		stock, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, fmt.Errorf("invalid stock value: %v", err)
		}

		createdAt, err := time.Parse("2006-01-02 15:04:05", record[4])
		if err != nil {
			return nil, fmt.Errorf("invalid created_at value: %v", err)
		}

		updatedAt, err := time.Parse("2006-01-02 15:04:05", record[5])
		if err != nil {
			return nil, fmt.Errorf("invalid updated_at value: %v", err)
		}

		// Construct the Product struct
		product := model.Product{
			Name:        record[0],
			Description: record[1],
			Price:       price,
			Stock:       stock,
			CreatedAt:   &createdAt,
			UpdatedAt:   &updatedAt,
		}

		// Append to the slice of products
		products = append(products, product)
	}

	u.ProductRepo.BulkCreateProducts(products, 1000)

	return products, nil
}

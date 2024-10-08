package repository

import (
	"fmt"
	"redis-demo/db/model"
	"strconv"
	"time"
)

type ProductRepository interface {
	CreateProduct(Product model.Product) (model.Product, error)
	GetProduct(id uint) (model.Product, error)
	UpdateProduct(id uint, toUpdate map[string]interface{}) error
	DeleteProduct(id uint) error
	GetAllProducts() ([]model.Product, error)
	BulkCreateProducts(products []model.Product, chunkSize int) error
	UpdateProductWithRedis(id uint, toUpdate map[string]interface{}) error
	GetProductWithRedis(id uint) (model.Product, error)
	DeleteProductWithRedis(id uint) error
	GetAllProductsWithRedis() ([]model.Product, error)
}

type ProductRepoImpl struct {
	Client *DBClient
}

func NewProductRepository(db *DBClient) ProductRepository {
	return &ProductRepoImpl{Client: db}
}

func (u *ProductRepoImpl) CreateProduct(Product model.Product) (model.Product, error) {
	err := u.Client.db.Create(&Product).Error
	if err != nil {
		return model.Product{}, err
	}

	return Product, nil
}

func (u *ProductRepoImpl) GetProduct(id uint) (model.Product, error) {
	var Product model.Product
	err := u.Client.db.Where("id = ?", id).First(&Product).Error
	if err != nil {
		return model.Product{}, err
	}

	return Product, nil
}

func (u *ProductRepoImpl) GetProductWithRedis(id uint) (model.Product, error) {
	var product model.Product

	err := u.Client.redisClient.GetStruct(strconv.Itoa(int(id)), &product)
	if err == nil {
		return product, nil
	}

	// Cache miss or Redis error, fetch from DB
	err = u.Client.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return model.Product{}, err
	}

	// After fetching from DB, store the result in Redis for future requests
	err = u.Client.redisClient.SetJSON(strconv.Itoa(int(id)), product, 10*time.Minute)
	if err != nil {
		fmt.Printf("Failed to cache product in Redis: %v\n", err) // Log error but don't stop execution
	}

	return product, nil
}

func (u *ProductRepoImpl) UpdateProductWithRedis(id uint, toUpdate map[string]interface{}) error {
	err := u.Client.db.Model(&model.Product{}).Where("id = ?", id).Updates(toUpdate).Error
	if err != nil {
		return err
	}

	err = u.Client.redisClient.InvalidateCacheKey(fmt.Sprintf("%d", id))
	if err != nil {
		fmt.Printf("Failed to invalidate cache for product %d: %v\n", id, err)
		return err
	}

	product, err := u.GetProduct(id)
	if err != nil {
		fmt.Printf("failed to get product %d: %v\n", id, err)
		return err
	}

	err = u.Client.redisClient.SetJSON(fmt.Sprintf("%d", product.ID), product, 10*time.Minute)
	if err != nil {
		fmt.Printf("Failed to cache product in Redis: %v\n", err)
		return err
	}

	return nil
}

func (u *ProductRepoImpl) UpdateProduct(id uint, toUpdate map[string]interface{}) error {
	return u.Client.db.Model(&model.Product{}).Where("id = ?", id).Updates(toUpdate).Error
}

func (u *ProductRepoImpl) DeleteProduct(id uint) error {
	err := u.Client.db.Unscoped().Where("id = ?", id).Delete(&model.Product{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *ProductRepoImpl) DeleteProductWithRedis(id uint) error {
	err := u.Client.db.Unscoped().Where("id = ?", id).Delete(&model.Product{}).Error
	if err != nil {
		return err
	}

	err = u.Client.redisClient.InvalidateCacheKey(fmt.Sprintf("%d", id))
	if err != nil {
		fmt.Printf("failed to invalidate cache for product %d: %v\n", id, err)
	}

	return nil
}

func (u *ProductRepoImpl) GetAllProducts() ([]model.Product, error) {
	var Products []model.Product
	err := u.Client.db.Find(&Products).Error
	if err != nil {
		return nil, err
	}

	return Products, nil
}

func (u *ProductRepoImpl) GetAllProductsWithRedis() ([]model.Product, error) {
	var allProducts []model.Product

	err := u.Client.redisClient.GetStruct("all_products", &allProducts)
	if err == nil {
		return allProducts, nil
	}

	err = u.Client.db.Find(&allProducts).Error
	if err != nil {
		return nil, err
	}

	err = u.Client.redisClient.SetJSON("all_products", allProducts, 10*time.Minute)
	if err != nil {
		fmt.Printf("Failed to cache products in Redis: %v\n", err)
	}

	return allProducts, nil
}

// BulkCreateProducts adds products in chunks to avoid overloading memory or database limits
func (u *ProductRepoImpl) BulkCreateProducts(products []model.Product, chunkSize int) error {
	total := len(products)

	// Loop over the products in chunks
	for i := 0; i < total; i += chunkSize {
		end := i + chunkSize
		if end > total {
			end = total
		}

		chunk := products[i:end]

		// Bulk insert the current chunk
		if err := u.Client.db.Create(&chunk).Error; err != nil {
			return fmt.Errorf("failed to insert products chunk: %v", err)
		}
	}

	return nil
}

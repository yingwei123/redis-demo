package handler

import (
	"encoding/json"
	"net/http"
	"redis-demo/db/model"
	"redis-demo/db/service"
	"redis-demo/server/util"
)

type Product struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
}

func CreateProduct(p service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var product Product

		err := util.ParseFromRequest(r, &product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		p, err := p.CreateProduct(model.Product{
			Name:        product.Name,
			Price:       float64(product.Price),
			Stock:       product.Stock,
			Description: product.Description,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(p)
	}
}

func GetProduct(p service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := util.ParseIdFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		p, err := p.GetProduct(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	}
}

func GetProductWithRedis(p service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := util.ParseIdFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		p, err := p.GetProductWithRedis(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	}
}

func UpdateProduct(p service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := util.ParseIdFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var product Product

		err = util.ParseFromRequest(r, &product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = p.UpdateProduct(uint(id), map[string]interface{}{
			"Name":        product.Name,
			"Price":       product.Price,
			"Stock":       product.Stock,
			"Description": product.Description,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Product updated successfully"))
	}
}

func UpdateProductWithRedis(p service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := util.ParseIdFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var product Product

		err = util.ParseFromRequest(r, &product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = p.UpdateProductWithRedis(uint(id), map[string]interface{}{
			"Name":        product.Name,
			"Price":       product.Price,
			"Stock":       product.Stock,
			"Description": product.Description,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Product updated successfully"))
	}
}

func DeleteProduct(p service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := util.ParseIdFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = p.DeleteProduct(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteProductWithRedis(p service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := util.ParseIdFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = p.DeleteProductWithRedis(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func GetAllProducts(p service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := p.GetAllProducts()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	}
}

func GetAllProductsWithRedis(p service.ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := p.GetAllProductsWithRedis()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	}
}

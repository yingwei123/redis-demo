package server

import (
	"redis-demo/db/service"
	"redis-demo/server/handler.go"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router, p service.ProductService) {
	r.HandleFunc("/product", handler.CreateProduct(p)).Methods("POST")
	r.HandleFunc("/product/{id}", handler.GetProduct(p)).Methods("GET")
	r.HandleFunc("/product/{id}", handler.UpdateProduct(p)).Methods("PUT")
	r.HandleFunc("/product/{id}", handler.DeleteProduct(p)).Methods("DELETE")
	r.HandleFunc("/products", handler.GetAllProducts(p)).Methods("GET")
	r.HandleFunc("/products/redis", handler.GetAllProductsWithRedis(p)).Methods("GET")
	r.HandleFunc("/product/{id}/redis", handler.GetProductWithRedis(p)).Methods("GET")
	r.HandleFunc("/product/{id}/redis", handler.UpdateProductWithRedis(p)).Methods("PUT")
	r.HandleFunc("/product/{id}/redis", handler.DeleteProductWithRedis(p)).Methods("DELETE")
}

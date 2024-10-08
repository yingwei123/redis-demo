package server

import (
	"redis-demo/db/repository"
	"redis-demo/db/service"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Server struct {
	Server *negroni.Negroni
}

func CreateServer(db *repository.DBClient) *Server {
	router := mux.NewRouter()

	productService := service.NewProductService(repository.NewProductRepository(db))

	RegisterProductRoutes(router, productService)

	n := negroni.New()
	n.UseHandler(router)
	n.Use(negroni.NewLogger())

	return &Server{
		Server: n,
	}
}

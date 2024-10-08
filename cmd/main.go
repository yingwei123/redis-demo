package main

import (
	"fmt"
	"log"
	"net/http"
	"redis-demo/config"
	"redis-demo/datagen"
	"redis-demo/db/repository"
	"redis-demo/db/service"
	"redis-demo/rclient"
	"redis-demo/server"
)

func main() {
	cfg, err := config.LoadENV()
	if err != nil {
		log.Fatal(err)
	}

	redisClient, err := rclient.CreateRedisClient(cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.CreateDBConnection(*cfg.DB, redisClient)
	if err != nil {
		log.Fatal(err)
	}

	//uncomment this and run it if you want to test with a larger dataset, replace x with the number of data you want to generate
	// GenerateDummyData(x)

	//uncomment this and run it once to init data to the product table for testing.
	//This repo comes with product.csv, replace this with products.csv or the csv file you generate with the GenerateDummyData function
	// InitProductData(db, "products.csv")

	server := server.CreateServer(db)

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), server.Server)
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateDummyData(num int) {
	datagen.GenerateProductData("products-extended.csv", num)
}

// this function is used to initialize the product data in the database
func InitProductData(db *repository.DBClient, filePath string) error {
	productService := service.NewProductService(repository.NewProductRepository(db))
	p, err := productService.ReadCSVToProducts(filePath)
	if err != nil {
		return err
	}

	println("Number of products created: ", len(p))

	return nil
}

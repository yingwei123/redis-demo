package datagen

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

// GenerateProductData generates product data and saves it as a CSV file
func GenerateProductData(csvName string, numProducts int) error {
	// Seed random generator
	rand.Seed(time.Now().UnixNano())

	// Create a CSV file with the specified name
	file, err := os.Create(csvName)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Create CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	header := []string{"Name", "Description", "Price", "Stock", "CreatedAt", "UpdatedAt"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("error writing header: %v", err)
	}

	// Generate product data
	for i := 0; i < numProducts; i++ {
		// Generate random product data
		name := gofakeit.ProductName()
		description := gofakeit.Sentence(10)
		price := fmt.Sprintf("%.2f", gofakeit.Price(10, 1500))
		stock := strconv.Itoa(gofakeit.Number(1, 100))
		createdAt := gofakeit.Date().Format("2006-01-02 15:04:05")
		updatedAt := gofakeit.Date().Format("2006-01-02 15:04:05")

		// Write row to CSV
		row := []string{name, description, price, stock, createdAt, updatedAt}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error writing row: %v", err)
		}
	}

	fmt.Printf("Generated %d products in %s\n", numProducts, csvName)
	return nil
}

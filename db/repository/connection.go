package repository

import (
	"fmt"
	"redis-demo/config"
	"redis-demo/db/model"
	"redis-demo/rclient"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBClient struct {
	db          *gorm.DB
	redisClient *rclient.RedisClient
}

// CreateDBConnection creates a connection to the PostgreSQL database using the provided parameters
func CreateDBConnection(config config.DBConfig, redisClient *rclient.RedisClient) (*DBClient, error) {
	// Create DSN using formatted string
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.User, config.Password, config.DBName, config.Host, config.DBPort)

	// Connect to the database using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	//configure connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB from gorm.DB: %w", err)
	}

	// Configure the connection pool
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)

	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %w", err)
	}

	println("Successfully connected to PostgreSQL")

	return &DBClient{db: db, redisClient: redisClient}, nil
}

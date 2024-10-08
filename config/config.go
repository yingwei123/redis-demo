package config

import (
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Config struct {
	Server *ServerConfig
	DB     *DBConfig
	Redis  *RedisConfig
}

type ServerConfig struct {
	Port uint `env:"PORT,default=3000"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST,default=localhost"`
	Port     string `env:"REDIS_PORT,default=6379"`
	Password string `env:"REDIS_PASSWORD,default="`
	DB       int    `env:"REDIS_DB,default=0"`
}

type DBConfig struct {
	Host            string        `env:"HOST,required"`
	User            string        `env:"USER,required"`
	Password        string        `env:"PASSWORD,required"`
	DBName          string        `env:"DB_NAME,required"`
	DBPort          string        `env:"DB_PORT,required"`
	MaxIdleConns    int           `env:"MAX_IDLE_CONNS, default=10"`
	MaxOpenConns    int           `env:"MAX_OPEN_CONNS, default=10"`
	ConnMaxLifetime time.Duration `env:"CONN_MAX_LIFETIME, default=0"`
}

// LoadENV loads the environment variables from the .env file and returns a Config struct
func LoadENV() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	serverConfig, err := LoadServerConfig()
	if err != nil {
		return nil, err
	}

	dbConfig, err := LoadDBConfig()
	if err != nil {
		return nil, err
	}

	redisConfig, err := LoadRedisConfig()
	if err != nil {
		return nil, err
	}

	return &Config{DB: dbConfig, Server: serverConfig, Redis: redisConfig}, nil
}

// LoadRedisConfig loads the Redis configuration from the .env file and returns a RedisConfig struct
func LoadRedisConfig() (*RedisConfig, error) {
	var cfg RedisConfig
	if err := loadConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// LoadServerConfig loads the server configuration from the .env file and returns a ServerConfig struct
func LoadServerConfig() (*ServerConfig, error) {
	var cfg ServerConfig
	if err := loadConfig(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// LoadDBConfig loads the database configuration from the .env file and returns a DBConfig struct
func LoadDBConfig() (*DBConfig, error) {
	var cfg DBConfig
	if err := loadConfig(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func loadConfig(cfg interface{}) error {
	return envdecode.Decode(cfg)
}

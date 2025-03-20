package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// AppConfig represents the application's configuration
type AppConfig struct {
	Server   ServerConfig   `json:"server"`
	MongoDB  MongoDBConfig  `json:"mongodb"`
	Redis    RedisConfig    `json:"redis"`
	Security SecurityConfig `json:"security"`
}

// ServerConfig represents the HTTP server configuration
type ServerConfig struct {
	Port         string `json:"port"`
	ReadTimeout  int    `json:"readTimeout"`
	WriteTimeout int    `json:"writeTimeout"`
}

// MongoDBConfig represents the MongoDB connection configuration
type MongoDBConfig struct {
	URI             string `json:"uri"`
	Database        string `json:"database"`
	CertificatePath string `json:"certificatePath"`
}

// RedisConfig represents the Redis connection configuration
type RedisConfig struct {
	URI      string `json:"uri,omitempty"`
	Addr     string `json:"addr,omitempty"`
	Password string `json:"password,omitempty"`
	DB       int    `json:"db,omitempty"`
}

// SecurityConfig represents security-related configurations
type SecurityConfig struct {
	JWTSecret string `json:"jwtSecret"`
}

var (
	config     *AppConfig
	configOnce sync.Once
)

// GetConfig loads and returns the application config
func GetConfig() (*AppConfig, error) {
	var err error
	configOnce.Do(func() {
		config = &AppConfig{}

		// First try to load from config file
		configFile := getConfigPath()
		err = loadConfigFromFile(configFile, config)
		if err != nil {
			// If file doesn't exist or has an error, use default values
			setDefaults(config)
		}

		// Override with environment variables if set
		overrideWithEnv(config)
	})

	return config, err
}

// getConfigPath returns the configuration file path
func getConfigPath() string {
	// Check for custom config path in environment
	customPath := os.Getenv("CONFIG_PATH")
	if customPath != "" {
		return customPath
	}

	// Default to config/app.config.json in the project root
	return filepath.Join("config", "app.config.json")
}

// loadConfigFromFile loads config from a JSON file
func loadConfigFromFile(path string, cfg *AppConfig) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(cfg)
}

// setDefaults sets default configuration values
func setDefaults(cfg *AppConfig) {
	// Server defaults
	cfg.Server.Port = "8080"
	cfg.Server.ReadTimeout = 10
	cfg.Server.WriteTimeout = 10

	// MongoDB defaults - empty, require explicit configuration
	cfg.MongoDB.URI = ""
	cfg.MongoDB.Database = "balancelife"
	cfg.MongoDB.CertificatePath = ""

	// Redis defaults - empty, require explicit configuration
	cfg.Redis.URI = ""
	cfg.Redis.Addr = ""
	cfg.Redis.Password = ""
	cfg.Redis.DB = 0

	// Security defaults - empty, require explicit configuration
	cfg.Security.JWTSecret = ""
}

// overrideWithEnv overrides config with environment variables
func overrideWithEnv(cfg *AppConfig) {
	// Server settings
	if val := os.Getenv("SERVER_PORT"); val != "" {
		cfg.Server.Port = val
	}

	// MongoDB settings
	if val := os.Getenv("MONGODB_URI"); val != "" {
		cfg.MongoDB.URI = val
	}
	if val := os.Getenv("MONGODB_DATABASE"); val != "" {
		cfg.MongoDB.Database = val
	}
	if val := os.Getenv("MONGODB_CERT_PATH"); val != "" {
		cfg.MongoDB.CertificatePath = val
	}

	// Redis settings
	if val := os.Getenv("REDIS_URI"); val != "" {
		cfg.Redis.URI = val
	}
	if val := os.Getenv("REDIS_ADDR"); val != "" {
		cfg.Redis.Addr = val
	}
	if val := os.Getenv("REDIS_PASSWORD"); val != "" {
		cfg.Redis.Password = val
	}
	if val := os.Getenv("REDIS_DB"); val != "" {
		var db int
		if _, err := fmt.Sscanf(val, "%d", &db); err == nil {
			cfg.Redis.DB = db
		}
	}

	// Security settings
	if val := os.Getenv("JWT_SECRET"); val != "" {
		cfg.Security.JWTSecret = val
	}
}

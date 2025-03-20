package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load environment variables
	envPath := "config/.env"
	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
		log.Println("Using environment variables instead")
	}

	fmt.Println("=====================================================")
	fmt.Println("         BalanceLife Database Connection Test        ")
	fmt.Println("=====================================================")
	fmt.Println(os.Getenv("MONGODB_CERT_PATH"))

	// Test MongoDB connection
	fmt.Println("🔍 TESTING MONGODB CONNECTION:")
	fmt.Println("----------------------------------------------------")
	testMongoDBConnection()
	fmt.Println()

	// Test Redis connection
	fmt.Println("🔍 TESTING REDIS CONNECTION:")
	fmt.Println("----------------------------------------------------")
	testRedisConnection()
}

func testMongoDBConnection() {
	// Get MongoDB connection details from environment
	uri := os.Getenv("MONGODB_URI")
	database := os.Getenv("MONGODB_DATABASE")
	certPath := os.Getenv("MONGODB_CERT_PATH")

	if uri == "" {
		log.Println("❌ ERROR: MONGODB_URI is not set in environment variables or .env file")
		return
	}

	if database == "" {
		database = "balancelife"
		log.Printf("⚠️ MONGODB_DATABASE not set, using default: %s", database)
	}

	log.Printf("MongoDB URI: %s", maskConnectionString(uri))
	log.Printf("Database: %s", database)

	// Check certificate file if using X509 auth
	if certPath != "" {
		_, err := os.Stat(certPath)
		if os.IsNotExist(err) {
			log.Printf("❌ ERROR: Certificate file does not exist at path: %s", certPath)
			return
		}
		log.Printf("Using X.509 certificate: %s", certPath)
	}

	// Configure MongoDB connection
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)

	// Add TLS certificate if provided
	if certPath != "" {
		clientOptions.SetTLSConfig(nil) // Setting to nil lets the driver use system certs
		clientOptions = clientOptions.ApplyURI(fmt.Sprintf(
			"%s&tlsCertificateKeyFile=%s",
			uri,
			certPath,
		))
	}

	// Set a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	log.Println("Attempting to connect to MongoDB...")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Printf("❌ ERROR: Failed to connect to MongoDB: %v", err)
		return
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("⚠️ Warning: Error disconnecting from MongoDB: %v", err)
		}
	}()

	// Ping the database to verify connection
	log.Println("Pinging MongoDB server...")
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("❌ ERROR: Failed to ping MongoDB: %v", err)
		return
	}

	// Get database stats
	log.Printf("Connection successful!")
	db := client.Database(database)

	// Try to list collections
	collections, err := db.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		log.Printf("⚠️ Warning: Failed to list collections: %v", err)
	} else {
		if len(collections) == 0 {
			log.Printf("No collections found in database '%s'", database)
		} else {
			log.Printf("Collections in database '%s':", database)
			for i, name := range collections {
				log.Printf("  %d. %s", i+1, name)
			}
		}
	}

	// Get server stats
	var result bson.M
	err = db.RunCommand(ctx, bson.D{{Key: "buildInfo", Value: 1}}).Decode(&result)
	if err != nil {
		log.Printf("⚠️ Warning: Cannot get server info: %v", err)
	} else {
		if version, ok := result["version"].(string); ok {
			log.Printf("MongoDB server version: %s", version)
		}
	}

	log.Println("✅ MongoDB connection test completed successfully!")
}

func testRedisConnection() {
	// Get Redis connection details from environment
	addr := os.Getenv("REDIS_ADDR")
	password := os.Getenv("REDIS_PASSWORD")

	if addr == "" {
		log.Println("❌ ERROR: REDIS_ADDR is not set in environment variables or .env file")
		return
	}

	log.Printf("Redis Address: %s", addr)
	if password != "" {
		log.Println("Redis Password: [REDACTED]")
	}

	// Configure Redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    password,
		DB:          0, // default DB
		DialTimeout: 10 * time.Second,
	})

	// Set a timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ping the Redis server to verify connection
	log.Println("Pinging Redis server...")
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("❌ ERROR: Failed to ping Redis: %v", err)
		return
	}

	log.Printf("Redis responded: %s", pong)

	// Try basic operations
	log.Println("Testing basic Redis operations...")

	// Set a test key
	testKey := "balance_life:connection_test"
	testValue := "Connected at " + time.Now().Format(time.RFC3339)

	err = rdb.Set(ctx, testKey, testValue, 1*time.Minute).Err()
	if err != nil {
		log.Printf("❌ ERROR: Failed to set test key: %v", err)
		return
	}
	log.Printf("Successfully set test key '%s'", testKey)

	// Get the test key
	val, err := rdb.Get(ctx, testKey).Result()
	if err != nil {
		log.Printf("❌ ERROR: Failed to get test key: %v", err)
		return
	}
	log.Printf("Retrieved test key value: %s", val)

	// Delete the test key
	err = rdb.Del(ctx, testKey).Err()
	if err != nil {
		log.Printf("⚠️ Warning: Failed to delete test key: %v", err)
	} else {
		log.Printf("Successfully deleted test key '%s'", testKey)
	}

	// Get Redis info
	infoStr, err := rdb.Info(ctx).Result()
	if err != nil {
		log.Printf("⚠️ Warning: Cannot get Redis info: %v", err)
	} else {
		// Extract and print Redis version
		lines := strings.Split(infoStr, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "redis_version:") {
				parts := strings.Split(line, ":")
				if len(parts) > 1 {
					log.Printf("Redis server version: %s", strings.TrimSpace(parts[1]))
					break
				}
			}
		}
	}

	// Clean up
	log.Println("Closing Redis connection...")
	err = rdb.Close()
	if err != nil {
		log.Printf("⚠️ Warning: Error closing Redis connection: %v", err)
	}

	log.Println("✅ Redis connection test completed successfully!")
}

// Helper function to mask sensitive parts of connection strings
func maskConnectionString(uri string) string {
	// For MongoDB URI with username and password
	if strings.Contains(uri, "@") {
		parts := strings.Split(uri, "@")
		if len(parts) >= 2 {
			authParts := strings.Split(parts[0], "://")
			if len(authParts) >= 2 {
				return authParts[0] + "://[credentials-hidden]@" + parts[1]
			}
		}
	}
	return uri
}

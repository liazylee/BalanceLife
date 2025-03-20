package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/zhenyili/BalanceLife/src/config"
	"github.com/zhenyili/BalanceLife/src/db"
	"github.com/zhenyili/BalanceLife/src/handlers"

	// Import the docs package
	_ "github.com/zhenyili/BalanceLife/docs"
)

// @title           BalanceLife API
// @version         1.0
// @description     A calorie tracking application backend with meal and workout tracking
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    https://github.com/zhenyili/BalanceLife
// @contact.email  support@balancelife.example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Load environment variables from .env file if it exists
	if err := godotenv.Load("config/.env"); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
		log.Println("Using environment variables or defaults")
	}

	// Load application configuration
	cfg, err := config.GetConfig()
	if err != nil {
		log.Printf("Warning: Error loading config: %v, using defaults", err)
	}

	// Initialize data store with MongoDB
	mongoStore, err := db.NewMongodbStore(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v. MongoDB is required to run the application.", err)
	}

	log.Println("Successfully connected to MongoDB")
	store := mongoStore

	// Ensure we close database connections when the application exits
	defer func() {
		log.Println("Closing database connections...")
		if err := mongoStore.Close(); err != nil {
			log.Printf("Error closing database connections: %v", err)
		}
	}()

	// Initialize router
	router := gin.Default()

	// Setup CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Basic health check route
	// @Summary      Health Check
	// @Description  Get the status of the server and its dependencies
	// @Tags         system
	// @Produce      json
	// @Success      200  {object}  map[string]interface{}
	// @Router       /health [get]
	router.GET("/health", func(c *gin.Context) {
		usingMongoDB := mongoStore.HasMongoDB()

		c.JSON(200, gin.H{
			"status":        "ok",
			"message":       "BalanceLife API is running",
			"store_type":    "mongodb",
			"using_mongodb": usingMongoDB,
		})
	})

	// API routes
	api := router.Group("/api")

	// Initialize handlers and register routes
	userHandler := handlers.NewUserHandler(store)
	userHandler.RegisterRoutes(api)

	mealHandler := handlers.NewMealHandler(store)
	mealHandler.RegisterRoutes(api)

	workoutHandler := handlers.NewWorkoutHandler(store)
	workoutHandler.RegisterRoutes(api)

	// Get port from config or use default
	port := cfg.Server.Port
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Starting BalanceLife API server on port %s...", port)
	log.Printf("Swagger documentation available at http://localhost:%s/swagger/index.html", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

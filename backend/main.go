package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"github.com/yourorg/auth-service/internal/api"
	"github.com/yourorg/auth-service/internal/cache"
	"github.com/yourorg/auth-service/internal/service"
)

func main() {
	// Get Redis URL from environment variable
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379/0"
	}

	// Initialize cache
	ruleCache, err := cache.NewRuleCache(redisURL)
	if err != nil {
		log.Fatalf("Failed to initialize cache: %v", err)
	}

	// Initialize validator
	validator := service.NewValidator(ruleCache)

	// Initialize handler
	apiHandler := api.NewHandler(validator)

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// Setup routes with CORS
	mux := http.NewServeMux()
	mux.HandleFunc("/validate", apiHandler.ValidateAccess)

	// Wrap with CORS handler
	handler := c.Handler(mux)

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Start server
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

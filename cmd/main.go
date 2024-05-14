package main

import (
    "log"
    "os"
    "os/signal"
	"net/http"

    "github.com/gin-gonic/gin"
    
)

func main() {
    r := gin.Default()

    // Set up routes
    r.POST("/upload", func(c *gin.Context) {
		// Upload file
	})

    // r.GET("/chat", )

    // Start the server
    srv := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }
 			   go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Failed to start server: %v", err)
        }
    }()

quit := make(chan os.Signal, 1)
signal.Notify(quit, os.Interrupt)
<-quit
log.Println("Shutting down server...")

if err := srv.Close(); err != nil {
	log.Fatalf("Error shutting down server: %v", err)
}
package startup

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BevisDev/BevisBot/internal/app/config"
	"github.com/BevisDev/BevisBot/internal/app/router"
	"github.com/gin-gonic/gin"
)

// Run starts the application, sets up signal handling, and ensures graceful shutdown.
func Run() {
	// Initialize context for application lifecycle
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up signal handling for graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Start application
	err := Initialize(ctx)
	if err != nil {
		return
	}

	// init engine
	var r *gin.Engine
	cf := config.AppConfig.Server

	if cf.Profile == "prod" || cf.Profile == "prod-job" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()

		// Handle panics
		r.Use(gin.Recovery())
	} else {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	}

	// register router
	router.RegisterRouter(r)
	r.Run()

	// Close connections on exit
	//defer func() {
	//	lib.Logger.Sync()
	//	lib.DbODS.Close()
	//	lib.DatabaseDE.Close()
	//	lib.Rdb01.Close()
	//	if !config.SystemConfig.RabbitMQ.IsDisable {
	//		lib.RabbitMQ.Close()
	//	}
	//}()

	// Set trusted proxies
	if len(cf.TrustedProxies) > 0 {
		if err := r.SetTrustedProxies(cf.TrustedProxies); err != nil {
			log.Printf("SetTrustedProxies error: %v", err)
		}
	}

	// Configure server
	srv := &http.Server{
		Addr:    cf.Port,
		Handler: r,
	}

	// Channel to signal server shutdown completion
	serverShutdown := make(chan bool, 1)

	// Start server
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("ListenAndServe error: %v", err)
		}
		serverShutdown <- true
	}()

	// Handle shutdown signal
	go func() {
		<-sig
		log.Println("Received shutdown signal...")

		// Create context with timeout for graceful shutdown
		shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 30*time.Second)
		defer shutdownCancel()

		// Attempt graceful shutdown
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("Server forced to shutdown due to timeout: %v", err)

			// Force close if graceful shutdown fails
			if closeErr := srv.Close(); closeErr != nil {
				log.Printf("Error force closing server:: %v", err)
			}
		} else {
			log.Println("Server shutdown completed")
		}

		// Cancel application context after server shutdown
		cancel()
	}()

	// Wait for either context cancellation or server shutdown
	select {
	case <-ctx.Done():
		log.Println("Application context cancelled, shutting down...")
	case <-serverShutdown:
		log.Println("Server shutting down...")
	}
}

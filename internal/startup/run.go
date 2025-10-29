package startup

import (
	"github.com/BevisDev/BevisBot/internal/app/router"
	"github.com/BevisDev/godev/utils"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

// Run starts the application, sets up signal handling, and ensures graceful shutdown.
func Run() {
	// Initialize context for application lifecycle
	ctx, cancel := utils.NewCtxCancel(nil)
	defer cancel()

	// Get application state
	//var state = utils.GetState(ctx)

	// Set up signal handling for graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Start application
	err := Initialize(ctx)
	if err != nil {
		//lib.Logger.Fatal(state, "engine is nil")
		return
	}

	// init engine
	var r *gin.Engine
	//cf := config.SystemConfig.ServerConfig
	//
	//if cf.Profile == "prod" || cf.Profile == "prod-job" {
	//	gin.SetMode(gin.ReleaseMode)
	//	r = gin.New()
	//
	//	// Handle panics
	//	r.Use(gin.Recovery())
	//} else {
	gin.SetMode(gin.DebugMode)
	gin.ForceConsoleColor()
	r = gin.Default()
	//}

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

	// get config
	//cf := config.SystemConfig.ServerConfig

	// Set trusted proxies
	//if len(cf.TrustedProxies) > 0 {
	//	if err := r.SetTrustedProxies(cf.TrustedProxies); err != nil {
	//		lib.Logger.Fatal(state, "error while setting trustedProxies: {}", err)
	//	}
	//}

	// Configure server
	//srv := &http.Server{
	//	Addr:    cf.Port,
	//	Handler: r,
	//}

	// Channel to signal server shutdown completion
	//serverShutdown := make(chan bool, 1)

	// Start server
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
	//		lib.Logger.Fatal(state, "server error: {}", err)
	//	}
	//	serverShutdown <- true
	//}()

	// Handle shutdown signal
	//go func() {
	//	<-sig
	//	lib.Logger.Info(state, "Received shutdown signal...")
	//
	//	// Create context with timeout for graceful shutdown
	//	shutdownCtx, shutdownCancel := utils.NewCtxTimeout(nil, 30)
	//	defer shutdownCancel()
	//
	//	// Attempt graceful shutdown
	//	if err := srv.Shutdown(shutdownCtx); err != nil {
	//		lib.Logger.Error(state, "Server forced to shutdown due to timeout: {}", err)
	//
	//		// Force close if graceful shutdown fails
	//		if closeErr := srv.Close(); closeErr != nil {
	//			lib.Logger.Error(state, "Error force closing server: {}", closeErr)
	//		}
	//	} else {
	//		lib.Logger.Info(state, "Server shutdown completed")
	//	}
	//
	//	// Cancel application context after server shutdown
	//	cancel()
	//}()
	//
	//// Wait for either context cancellation or server shutdown
	//select {
	//case <-ctx.Done():
	//	lib.Logger.Info(state, "Application context cancelled, shutting down...")
	//case <-serverShutdown:
	//	lib.Logger.Info(state, "Server shutting down...")
	//}
}

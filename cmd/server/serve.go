package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinoctavian/evodka_backend/pkg/config"
	"github.com/kevinoctavian/evodka_backend/pkg/middleware"
	"github.com/kevinoctavian/evodka_backend/pkg/router"
	"github.com/kevinoctavian/evodka_backend/platform/database"
	"github.com/kevinoctavian/evodka_backend/platform/logger"
)

func Serve() {
	appCfg := config.AppCfg()

	logger.SetUpLogger()
	logr := logger.GetLogger()

	if err := database.ConnectDB(); err != nil {
		logr.Panicf("Oops... database is not connected! error: %v", err)
		return
	}

	fiberCfg := config.FiberConfig()
	app := fiber.New(fiberCfg)

	middleware.FiberMiddleware(app) // Use the Logger middleware
	router.PublicRoutes(app)        // Set up the routes
	router.PrivateRoutes(app)       // Set up the private routes with JWT protection

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		logr.Infoln("Shutting down server...")
		_ = app.Shutdown()
	}()

	// start http server
	serverAddr := fmt.Sprintf("%s:%d", appCfg.Host, appCfg.Port)
	if err := app.Listen(serverAddr); err != nil {
		logr.Errorf("Oops... server is not running! error: %v", err)
	}
}

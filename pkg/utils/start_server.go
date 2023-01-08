package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for indle connection
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signal
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %+v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Oops.. Server in not running! Reason: %+v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting simple server
func StartServer(a *fiber.App) {
	// Run server
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Ooops.. Server is not running! Reason: %+v", err)
	}
}

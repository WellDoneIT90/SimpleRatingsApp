package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoute(a *fiber.App) {
	// Creating route group
	route := a.Group("/swagger")

	// route
	route.Get("*", swagger.HandlerDefault)
}

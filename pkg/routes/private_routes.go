package routes

import (
	"github.com/WellDoneIT90/SimpleRatingsApp/app/controllers"
	"github.com/WellDoneIT90/SimpleRatingsApp/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	// Creating route group
	route := a.Group("/api/v1")

	// Create route POST
	route.Post("/rating", middleware.JWTProtected(), controllers.CreateRating)

	// Create route DELETE
	route.Delete("/rating/:id", middleware.JWTProtected(), controllers.DeleteRating)

}

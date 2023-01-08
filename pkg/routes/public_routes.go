package routes

import (
	"github.com/WellDoneIT90/SimpleRatingsApp/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	// Craete Route Group
	route := a.Group("/api/v1")

	// Create Routes
	route.Get("/Ratings", controllers.GetRatings)
	route.Get("/Rating/:id", controllers.GetRating)
	route.Get("/token/new", controllers.GetNewAccessToken)
}

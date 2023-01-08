package routes

import "github.com/gofiber/fiber/v2"

func NotFoundRoute(a *fiber.App) {
	// Register new special route

	a.Use(
		// Annonimus function
		func(c *fiber.Ctx) error {
			// Return Status 404 error
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "sorry, endpoint not found",
			})
		},
	)
}

package main

import (
	"fmt"

	"github.com/WellDoneIT90/SimpleRatingsApp/pkg/configs"
	"github.com/WellDoneIT90/SimpleRatingsApp/pkg/middleware"
	"github.com/WellDoneIT90/SimpleRatingsApp/pkg/routes"
	"github.com/WellDoneIT90/SimpleRatingsApp/pkg/utils"
	"github.com/gofiber/fiber/v2"

	_ "github.com/WellDoneIT90/SimpleRatingsApp/docs"
	_ "github.com/joho/godotenv/autoload"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	fmt.Println("Ratings App in GO using Fiber and PostgreSQl deployed locally in Docker!")

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	// Routes
	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	// Start server with graceful shutdown
	utils.StartServerWithGracefulShutdown(app)
}

package controllers

import (
	"time"

	"github.com/WellDoneIT90/SimpleRatingsApp/app/models"
	"github.com/WellDoneIT90/SimpleRatingsApp/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetRatings func gets all ratings in Ratings
func GetRatings(c *fiber.Ctx) error {
	// Create database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all books
	ratings, err := db.GetRatings()
	if err != nil {
		// Return 404, if book is not found
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "Ratings where not found",
			"count":   0,
			"ratings": nil,
		})
	}

	// Return 200 OK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(ratings),
		"ratings": ratings,
	})
}

// GetRating funx gets one rating by given ID
func GetRating(c *fiber.Ctx) error {
	// Catch rating id from URL
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// create Database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get one rating by given ID
	rating, err := db.GetRating(id)
	if err != nil {
		// Return 404 not found, when rating was not found in Ratings
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "rating was not found in Ratings",
			"rating": nil,
		})
	}

	// Return 200 OK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"rating": rating,
	})
}

// CreateRating func to create a new rating
func CreateRating(c *fiber.Ctx) error {
	// Get time now
	now := time.Now().Unix()

	// Get claims from JWT
	claims, err := utils.ExtractTokenMetadata()
	if err != nil {
		// Return status 500 and JWT parse error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT to current rating
	expires = claims.Expires

	// Checking if Now is greather then expiration from JWT
	if now > expires {
		// Return status 401 unauthorized error
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration of your token",
		})
	}

	// Create a Ratings struct
	rating := &models.Ratings{}

	// Check received JSON from URL Body is valid
	if err := c.BodyParser(rating); err != nil {
		// Return status 400 error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Ratings model
	validate := utils.NewValidator()

	// Set initialized default data for ratings
	rating.ID = uuid.New()
	rating.CreatedAt = time.Now()

	// Validate rating fields
	if err := validate.Struct(rating); err != nil {
		// Return 400 if some fields are not valide
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create new rating
	if err := db.CreateRating(rating); err != nil {
		// Return status 500 error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status OK
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":  true,
		"msg":    nil,
		"rating": rating,
	})
}

// DeleteRating func to delete one rating by given ID from Ratings

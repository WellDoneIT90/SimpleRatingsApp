package controllers

import (
	"time"

	"github.com/WellDoneIT90/SimpleRatingsApp/app/models"
	"github.com/WellDoneIT90/SimpleRatingsApp/pkg/utils"
	"github.com/WellDoneIT90/SimpleRatingsApp/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetRatings func gets all ratings in Ratings
// @Description Get all existing ratings.
// @Summary get all existing ratings
// @Tags Ratings
// @Accept json
// @Produce json
// @Success 200 {array} models.Rating
// @Router /api/v1/ratings [get]
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

	// Get all ratings
	ratings, err := db.GetRatings()
	if err != nil {
		// Return 404, if ratings are not found
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "Ratings where not found: " + err.Error(),
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

// GetRating func gets one rating by given ID
// @Description Get rating by given ID.
// @Summary get rating by given ID
// @Tags Rating
// @Accept json
// @Produce json
// @Param id path string true "Rating ID"
// @Success 200 {object} models.Rating
// @Router /api/v1/rating/{id} [get]
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
			"msg":    "rating was not found in Ratings: " + err.Error(),
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
// @Description Create a new rating.
// @Summary create a new rating
// @Tags Rating
// @Accept json
// @Produce json
// @Param company body string true "Company"
// @Param author body string true "Author"
// @Param author_role body string true "AuthorRole"
// @Param company_rating body integer true "CompanyRating"
// @Param description body string true "Description"
// @Success 200 {object} models.Rating
// @Security ApiKeyAuth
// @Router /api/v1/rating [post]
func CreateRating(c *fiber.Ctx) error {
	// Get time now
	now := time.Now().Unix()

	// Get claims from JWT
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT to current rating
	expires := claims.Expires

	// Checking if Now is greather then expiration from JWT
	if now > expires {
		// Return status 401 unauthorized error
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration of your token",
		})
	}

	// Create a Ratings struct
	rating := &models.Rating{}

	// Check, if received JSON data is valid
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
		"error":  false,
		"msg":    nil,
		"rating": rating,
	})
}

// DeleteRating func to delete one rating by given ID from Ratings
// @Description Delete rating by given ID.
// @Summary delete rating by given ID
// @Tags Rating
// @Accept json
// @Produce json
// @Param id path string true "Rating ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /api/v1/rating/{id} [delete]
func DeleteRating(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current ratings.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new Ratings struct
	rating := &models.Rating{}

	// Catch rating id from URL
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Rating model.
	validate := utils.NewValidator()

	// Validate only one rating field ID.
	if err := validate.StructPartial(rating, id.String()); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if rating with given exists.
	foundedBook, err := db.GetRating(id)
	if err != nil {
		// Return status 404 and book not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "book with this ID not found",
		})
	}

	// Delete book by given ID.
	if err := db.DeleteRating(foundedBook.ID); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}

package controllers

import (
	"github.com/WellDoneIT90/SimpleRatingsApp/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// @Description Create a new access token.
// @Summary create a new access token
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Router /v1/token/new [get]
func GetNewAccessToken(c *fiber.Ctx) error {
	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}

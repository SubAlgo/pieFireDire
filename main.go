package main

import (
	"net/http"
	"pieFireDire/models"
	"pieFireDire/services"

	"github.com/gofiber/fiber/v2"
)

var r struct {
}

func main() {
	app := fiber.New()

	app.Get("/beef/summary", func(c *fiber.Ctx) error {
		serviceBeef := services.NewBeefService()
		err := serviceBeef.Read("./files/beef.txt")
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": err,
			})
		}

		res := models.Response{Beef: serviceBeef.Get()}
		return c.JSON(res)
	})

	app.Listen(":8000")
}

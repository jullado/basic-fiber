package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func MiddlewareMain(app fiber.Router) {

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	// app.Use("/", func(c *fiber.Ctx) error {
	// 	c.Next()
	// 	log.Print(c.Path())
	// 	return nil
	// })
}

package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func MiddlewareBasic(app fiber.Router) {

	app.Use("/wildcards", func(c *fiber.Ctx) error {
		c.Locals("name", "jullado") // กำหนดค่าตัวแปร (key, value) จาก middleware ไปใช้ใน route
		fmt.Println("before route")
		c.Next()
		fmt.Println("after route")
		return nil
	})
}

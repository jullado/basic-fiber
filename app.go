package main

import (
	"log"
	"test-fiber/config"
	"test-fiber/database"
	middlewares "test-fiber/middleware"
	"test-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	config.InitEnvironment()
	database.InitialDBMongo()
	database.InitialRedis()
}

func main() {
	app := fiber.New()

	// app.Server().MaxConnsPerIP = 1 // ตั้งค่าจำนวน connection ต่อ IP ** ถ้า route นั้นยังทำงานอยู่จะนับเป็น 1 connection เครื่องเดิมจะไม่สามารถเปิดแทบเชื่อมต่อเข้าไปอีกได้ถ้าเกิน

	app.Get("/", func(c *fiber.Ctx) error {
		go func() {
			log.Print(c.Path())
		}()
		return c.SendString("Hello, World!")
	})

	// group api
	basicGroup := app.Group("/api/v1/basic", func(c *fiber.Ctx) error {
		c.Set("version", "basic") // set header ได้ เมื่อเข้า group นี้
		return c.Next()
	})

	mongoGroup := app.Group("/mongo")
	requestGroup := app.Group("/request")
	centrifugalGroup := app.Group("/centrifugal")
	emailGroup := app.Group("/email")
	redisGroup := app.Group("/redis")

	// middleware
	middlewares.MiddlewareMain(app)
	middlewares.MiddlewareBasic(basicGroup)

	// routes
	routes.RouterBasic(basicGroup)
	routes.RouterMongo(mongoGroup)
	routes.RouterRequest(requestGroup)
	routes.RouterCentrifugal(centrifugalGroup)
	routes.RouterEmail(emailGroup)
	routes.RouterRedis(redisGroup)

	// Start server
	env := viper.GetString("app.env")
	port := viper.GetString("app.port")
	if env == "dev" || env == "development" {
		app.Listen("localhost:" + port)
	} else {
		app.Listen(":" + port)
	}
}

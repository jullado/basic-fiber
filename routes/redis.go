package routes

import (
	"encoding/json"
	"test-fiber/database"
	"test-fiber/model"

	"github.com/gofiber/fiber/v2"
)

func RouterRedis(app fiber.Router) {
	app.Get("/get", func(c *fiber.Ctx) error {

		getDataStr, err := database.Redis_Client.Get("name").Result()
		if err != nil {
			panic(err)
		}

		// get data **ถ้าข้อมูลเป็น json จะเป็น string แบบ json
		getDataJson, err := database.Redis_Client.Get("id123").Result()
		if err != nil {
			panic(err)
		}

		json.Unmarshal([]byte(getDataJson), &model.BodyData)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "get redis", "data": fiber.Map{"str": getDataStr, "json": model.BodyData}})
	})

	app.Get("/set", func(c *fiber.Ctx) error {
		// ใส่ key, value, expire

		// set data string
		err := database.Redis_Client.Set("name", "Julladith", 0).Err()
		if err != nil {
			panic(err)
		}

		// set data json
		data := struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{
			Name: "julladith",
			Age:  15,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		err = database.Redis_Client.Set("id123", jsonData, 0).Err()
		if err != nil {
			panic(err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "set redis", "data": ""})
	})
}

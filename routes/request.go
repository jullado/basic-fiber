package routes

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"test-fiber/model"

	"github.com/gofiber/fiber/v2"
)

func RouterRequest(app fiber.Router) {
	app.Get("/get", func(c *fiber.Ctx) error {
		url := "https://localhost:3002/api/v1/healthcheck"
		client := &http.Client{}

		// สร้าง request method
		req, _ := http.NewRequest("GET", url, nil)

		// ส่ง request
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// อ่านค่า response
		body, _ := ioutil.ReadAll(res.Body)

		// แปลง response เป็น map
		json.Unmarshal(body, &model.BodyData)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "get request", "data": model.BodyData})
	})

	app.Get("/post", func(c *fiber.Ctx) error {
		url := "https://localhost:3002/api/v1/authen"
		data_payload := fiber.Map{"username": "julladith", "password": "boss025614105"}
		payload, _ := json.Marshal(data_payload)
		client := &http.Client{}

		// สร้าง request method
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		req.Header.Add("Content-Type", "application/json")

		// ส่ง request
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// อ่านค่า response
		body, _ := ioutil.ReadAll(res.Body)

		// แปลง response เป็น map
		json.Unmarshal(body, &model.BodyData)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "post request", "data": model.BodyData})
	})
}

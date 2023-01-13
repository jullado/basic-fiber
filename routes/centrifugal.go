package routes

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"test-fiber/model"

	"github.com/gofiber/fiber/v2"
)

func RouterCentrifugal(app fiber.Router) {

	// send publish with channel
	app.Post("/publish", func(c *fiber.Ctx) error {
		url := "http://localhost:8000/api"
		api_key := "d7627bb6-2292-4911-82e1-615c0ed3eebb"
		data_payload := fiber.Map{
			"method": "publish",
			"params": fiber.Map{
				"channel": "new2",
				"data": fiber.Map{
					"data": "from backend",
				},
			},
		}
		payload, _ := json.Marshal(data_payload)
		client := &http.Client{}

		// สร้าง request method
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "apikey "+api_key)

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

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "request Centrifugal", "data": model.BodyData})
	})

	app.Post("/gen_token", func(c *fiber.Ctx) error {

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "request Centrifugal", "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyODI4MiIsImV4cCI6MTY2NjQyNjA5MSwiaWF0IjoxNjY1ODIxMjkxfQ.T18TED7QDwvLXjK2rBt_XR2t4KYY-AwhEH6yNC73oRg"})
	})
}

package routes

import (
	"fmt"
	"test-fiber/model"

	"github.com/gofiber/fiber/v2"
)

func RouterBasic(app fiber.Router) {

	// รับค่า params ปกติ ** http://localhost:8080/params2?type=json&num=555
	app.Get("/params1", func(c *fiber.Ctx) error {
		types := c.Query("type")
		number := c.Query("num")
		return c.JSON(fiber.Map{"message": types + number})
	})

	// รับค่า params ผ่านชื่อ path ** http://localhost:8080/params1/json/jullado  ** nameจะมีหรือไม่มีก็ได้
	app.Get("/params2/:type/:name?", func(c *fiber.Ctx) error {
		types := c.Params("type")
		name := c.Params("name")
		return c.SendString(types + name)
	})

	// รับค่า params ผ่านชื่อ path เป็น int เท่านั้น ** http://localhost:8080/params3/1
	app.Get("/params3/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.ErrBadRequest
		}
		return c.SendString(fmt.Sprintf("%v", id))
	})

	// รับค่า params ทั้งหมด เพื่อยัดเข้าไปใน struct ตาม key ที่กำหนด ** http://localhost:8080/params4?id=11name=jull
	app.Get("/params4", func(c *fiber.Ctx) error {
		req := model.Params4{}
		c.QueryParser(&req)
		return c.JSON(req)
	})

	// รับค่า path ทั้งหมด เป็น string ต่อกัน ** http://localhost:8080/wildcards/boss/jull   แสดง boss/jull
	app.Get("/wildcards/*", func(c *fiber.Ctx) error {
		wildcard := c.Params(("*"))
		name := c.Locals("name") // ค่าตัวแปรที่สร้างใน middleware
		fmt.Println("in route", name)
		return c.SendString(wildcard)
	})

	// แสดงหน้าเว็บ html
	app.Static("/webstatic", "../static", fiber.Static{
		Index: "test.html", // ถ้าไม่ใส่จะใช้เป็น index.html
	})

	// ดึงค่า env จาก fiber
	app.Get("/env", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"BaseURL":     c.BaseURL(),
			"Hostname":    c.Hostname(),
			"IP":          c.IP(),
			"IPs":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Subdomains":  c.Subdomains(),
		})
	})

	// รับค่า body json
	app.Post("/sendbody", func(c *fiber.Ctx) error {
		if err := c.BodyParser(&model.ReqData); err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": model.ReqData})
	})

	// ดึงค่า header
	app.Get("/getheader", func(c *fiber.Ctx) error {
		fmt.Println(c.Get("autorization"))
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "test header"})
	})
}

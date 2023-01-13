package routes

import (
	"log"
	"strconv"
	"test-fiber/database"
	"test-fiber/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RouterMongo(app fiber.Router) {

	app.Get("/find", func(c *fiber.Ctx) error {

		query := bson.M{
			"age": bson.M{"$gt": 20},
		}
		field := bson.M{"_id": 0}

		cursor, err := database.Users.Find(database.DB_Ctx, query, options.Find().SetProjection(field))

		if err != nil {
			panic(err)
		}

		// นำค่า table ทั้งหมด ยัดเข้า bson array
		if err := cursor.All(database.DB_Ctx, &model.ResultsFind); err != nil {
			panic(err)
		}

		cursor.Close(database.DB_Ctx)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "find", "data": model.ResultsFind})
	})

	app.Get("/findone", func(c *fiber.Ctx) error {
		query := bson.M{
			"age": bson.M{"$lt": 23},
		}

		// นำค่า single value ยัดเข้า bson
		if err := database.Users.FindOne(database.DB_Ctx, query).Decode(&model.ResultsFindOne); err != nil {
			if err.Error() == "mongo: no documents in result" {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "no data", "data": fiber.Map{}}) // return empty map
			}
			panic(err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "find one", "data": model.ResultsFindOne})
	})

	app.Post("/insertone", func(c *fiber.Ctx) error {
		if err := c.BodyParser(&model.BodyData); err != nil {
			return err
		}

		// insert แบบรับ data มาจาก body
		insertBodyResult, err := database.Users.InsertOne(database.DB_Ctx, model.BodyData)

		if err != nil {
			panic(err)
		}
		log.Printf("insert success: %v", insertBodyResult)

		// insert แบบเขียน data เข้าไปเอง
		insertResult, err := database.Users.InsertOne(database.DB_Ctx, bson.M{
			"username":      "julladith2",
			"password":      "222222",
			"age":           12,
			"phone":         "0659985548",
			"employee_code": 64216,
		})

		if err != nil {
			panic(err)
		}
		log.Printf("insert success: %v", insertResult)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "insert one", "data": fiber.Map{}})
	})

	app.Post("/insertmany", func(c *fiber.Ctx) error {
		if err := c.BodyParser(&model.BodyData); err != nil {
			return err
		}

		// insert many แบบรับ data มาจาก body และเขียนเอง
		insertManyResult, err := database.Users.InsertMany(database.DB_Ctx, []interface{}{
			model.BodyData,
			bson.M{
				"username":      "julladith4",
				"password":      "222222",
				"age":           12,
				"phone":         "0659985548",
				"employee_code": 64216,
			},
		})

		if err != nil {
			panic(err)
		}
		log.Printf("insert success: %v", insertManyResult)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "insert many", "data": fiber.Map{}})
	})

	app.Put("/updateone", func(c *fiber.Ctx) error {
		username := c.Query("username")
		if err := c.BodyParser(&model.BodyData); err != nil {
			return err
		}

		// update one แบบรับ params และ body
		updateOneResult, err := database.Users.UpdateOne(database.DB_Ctx, bson.M{"username": username},
			bson.M{"$set": model.BodyData},
		)

		if err != nil {
			panic(err)
		}
		log.Printf("update success: %v", updateOneResult)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "update one", "data": fiber.Map{}})
	})

	app.Put("/updatemany", func(c *fiber.Ctx) error {
		ageStr := c.Query("age")
		if err := c.BodyParser(&model.BodyData); err != nil {
			return err
		}

		ageInt, err := strconv.Atoi(ageStr)

		// update many แบบรับ params และ body
		updateManyResult, err := database.Users.UpdateMany(database.DB_Ctx, bson.M{"age": ageInt},
			bson.M{"$set": model.BodyData},
		)

		if err != nil {
			panic(err)
		}
		log.Printf("update success: %v", updateManyResult)
		log.Printf("show value of map : %v", model.BodyData.(map[string]interface{})["password"])

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "update many", "data": fiber.Map{}})
	})

	app.Delete("/deleteone", func(c *fiber.Ctx) error {
		username := c.Query("username")

		// delete one แบบรับ params
		DeleteOneResult, err := database.Users.DeleteOne(database.DB_Ctx, bson.M{"username": username})

		if err != nil {
			panic(err)
		}
		log.Printf("delete success: %v", DeleteOneResult)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "delete one", "data": fiber.Map{}})
	})

	app.Delete("/deletemany", func(c *fiber.Ctx) error {
		ageStr := c.Query("age")
		ageInt, err := strconv.Atoi(ageStr)

		// delete many แบบรับ params
		DeleteManyResult, err := database.Users.DeleteMany(database.DB_Ctx, bson.M{"age": ageInt})

		if err != nil {
			panic(err)
		}
		log.Printf("delete success: %v", DeleteManyResult)

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "delete many", "data": fiber.Map{}})
	})
}

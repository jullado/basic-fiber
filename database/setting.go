package database

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB        *mongo.Database
	DB_Ctx    = context.TODO()
	DB_Client *mongo.Client
	DB_Err    error

	// collectiion name
	Users *mongo.Collection
	Tests *mongo.Collection
	Anime *mongo.Collection
)

func InitialDBMongo() {
	DB_HOST := viper.GetString("app.db.host")
	DB_USERNAME := viper.GetString("app.db.username")
	DB_PASSWORD := viper.GetString("app.db.password")
	DB_NAME := viper.GetString("app.db.name")
	URI := "mongodb://" + DB_USERNAME + ":" + DB_PASSWORD + "@" + DB_HOST + "/?authSource=admin"

	DB_Client, DB_Err = mongo.Connect(DB_Ctx, options.Client().ApplyURI(URI))

	if DB_Err != nil {
		// Disconnect
		if err := DB_Client.Disconnect(DB_Ctx); err != nil {
			panic(err)
		}
		fmt.Println("DB connection error:", URI)
		panic(DB_Err)
	}

	DB = DB_Client.Database(DB_NAME)

	// collectiion name
	Users = DB.Collection("users")
	Tests = DB.Collection("tests")

	fmt.Println("DB connection:", URI)
}

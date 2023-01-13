package model

import "go.mongodb.org/mongo-driver/bson"

var (
	ReqData interface{}

	ResultsFind    []bson.M
	ResultsFindOne bson.M
	BodyData       interface{}
)

type Params4 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

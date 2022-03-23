package api

import (
	"context"
	"encoding/json"
	"first/internal/database"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	ID   int64  `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string `bson:"name,omitempty" json:"name,omitempty"`
	Age  string `bson:"age,omitempty" json:"age,omitempty"`
}

func clientCollection() (context.Context, *mongo.Collection, *mongo.Client) {
	ctx, db, client := database.Connect()
	return ctx, db.Collection("client"), client
}

// POST method - Create a new client
func CreateClientEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("Content-Type", "application/")

	var client Client
	json.NewDecoder(request.Body).Decode(&client)

	ctx, collection, clt := clientCollection()
	defer database.Disconnect(ctx, clt)

	result, err := collection.InsertOne(ctx, client)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

// GET method - Get all clients
func GetClientsEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var clients []Client

	ctx, collection, clt := clientCollection()
	defer database.Disconnect(ctx, clt)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &clients); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(clients)
}

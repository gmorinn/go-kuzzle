package main

import (
	"fmt"
	"kuzzletest/utils"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/protocol/websocket"
)

type TypeCollectionEnum string

const (
	Keyword  TypeCollectionEnum = "keyword"
	Text                        = "text"
	Integer                     = "integer"
	Float                       = "float"
	Date                        = "date"
	GeoPoint                    = "geo_point"
)

type Type struct {
	CollectionType TypeCollectionEnum `json:"type"`
}

type User struct {
	Username    Type `json:"username"`
	Birthday    Type `json:"birthday"`
	Age         Type `json:"age"`
	Description Type `json:"description"`
	GPA         Type `json:"gpa"`
}

type PayloadUser struct {
	Username    string    `json:"username"`
	Birthday    time.Time `json:"birthday"`
	Age         int64     `json:"age"`
	Description string    `json:"description"`
	GPA         float64   `json:"gpa"`
}

type CollectionUser struct {
	Properties User `json:"properties"`
}

func main() {
	// Creates a WebSocket connection.
	c := websocket.NewWebSocket(os.Getenv("HOST"), nil)
	// Instantiates a Kuzzle client
	kuzzle, _ := kuzzle.NewKuzzle(c, nil)
	// Connects to the server.
	err := kuzzle.Connect()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("Connected!")
	api := KuzzleAPI{
		API: kuzzle,
	}
	err = api.createIndex("test")
	var userTable CollectionUser = CollectionUser{
		Properties: User{
			Username:    Type{Keyword},
			Birthday:    Type{Date},
			Age:         Type{Integer},
			Description: Type{Text},
			GPA:         Type{Float},
		},
	}
	res := utils.GetFormatJSON(userTable)
	fmt.Println(string(res))
	err = api.DeleteCollection("test", "users")
	fmt.Println(err)
	err = api.CreateCollection("test", "users", res)

	payload := PayloadUser{
		Username:    "guillaume",
		Birthday:    time.Now(),
		Age:         19,
		Description: "AZERTYUIOP POIUYTREZA",
		GPA:         2.5,
	}
	id := uuid.New().String()
	res, err = api.CreateDocument("test", "users", id, payload)
	fmt.Println(string(res))
	fmt.Println(err)
	kuzzle.Disconnect()
}

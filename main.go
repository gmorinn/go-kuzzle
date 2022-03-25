package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/protocol/websocket"
)

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
	res, _ := api.listIndex()
	for _, v := range res {
		fmt.Println(v)
	}
	kuzzle.Disconnect()
}

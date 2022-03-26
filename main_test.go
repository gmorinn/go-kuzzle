package main

import (
	"log"
	"os"
	"testing"

	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/protocol/websocket"
)

func removeAllIndex(test *KuzzleAPI) {
	indexes, err := test.listIndex()
	if err != nil {
		log.Fatal(err)
	}
	if len(indexes) > 0 {
		err = test.DeleteManyIndex(indexes)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func newTest() *KuzzleAPI {
	c := websocket.NewWebSocket(os.Getenv("HOST"), nil)
	kuzzle, _ := kuzzle.NewKuzzle(c, nil)
	err := kuzzle.Connect()
	if err != nil {
		log.Fatal(err)
	}
	test := &KuzzleAPI{
		API: kuzzle,
	}

	// Remove all indexes
	removeAllIndex(test)
	/////////////////////////////////

	return test
}

var testAPI *KuzzleAPI

func TestMain(m *testing.M) {
	testAPI = newTest()
	os.Exit(m.Run())
}

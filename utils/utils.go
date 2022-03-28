package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"unicode"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func EmptyIndex() error {
	return fmt.Errorf("Index is empty!")
}

func EmptyCollection() error {
	return fmt.Errorf("Collection is empty!")
}

func GetFormatJSON(data interface{}) []byte {
	myIn, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return json.RawMessage(myIn)
}

func IsContainUpper(name string) bool {
	if name == "" {
		return false
	}
	for _, r := range name {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

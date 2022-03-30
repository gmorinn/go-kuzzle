package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"unicode"

	"github.com/google/uuid"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CheckErrorIndex(index string) error {
	if index == "" {
		return fmt.Errorf("Index is empty!")
	}
	if IsContainUpper(index) {
		return fmt.Errorf("Uppercase is not allowed!")
	}
	return nil
}

func CheckErrorID(id string) error {
	if id == "" {
		return fmt.Errorf("id is empty!")
	}
	if IsContainUpper(id) {
		return fmt.Errorf("Uppercase is not allowed!")
	}
	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("Wrong format id!")
	}
	return nil
}

func CheckErrorCollection(collection string) error {
	if collection == "" {
		return fmt.Errorf("Collection is empty!")
	}
	if IsContainUpper(collection) {
		return fmt.Errorf("Uppercase is not allowed!")
	}
	return nil
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

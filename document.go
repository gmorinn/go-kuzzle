package main

import (
	"encoding/json"
	"fmt"
	"kuzzletest/utils"
)

func (k *KuzzleAPI) CreateDocument(index, collection, id string, data interface{}) (json.RawMessage, error) {
	if err := utils.CheckErrorIndex(index); err != nil {
		return nil, err
	}
	if err := utils.CheckErrorCollection(collection); err != nil {
		return nil, err
	}
	if err := utils.CheckErrorID(id); err != nil {
		return nil, err
	}
	if isIndex, _ := k.ExistIndex(index); !isIndex {
		return nil, fmt.Errorf("Index \"%s\" doesn't exist", index)
	}
	if isCollection, _ := k.ExistCollection(index, collection); !isCollection {
		return nil, fmt.Errorf("Index \"%s\" with collection \"%s\" doesn't exist!", index, collection)
	}
	if res, _ := k.GetDocument(index, collection, id); res != nil {
		return nil, fmt.Errorf("ID \"%s\" already exists", id)
	}
	payload := utils.GetFormatJSON(data)
	res, err := k.API.Document.Create(index, collection, id, payload, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (k *KuzzleAPI) DeleteDocument(index, collection, id string) (string, error) {
	if res, _ := k.GetDocument(index, collection, id); res == nil {
		return "", fmt.Errorf("No data for ID \"%s\"!", id)
	}
	res, err := k.API.Document.Delete(index, collection, id, nil)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (k *KuzzleAPI) GetDocument(index, collection, id string) (json.RawMessage, error) {
	if isCollection, _ := k.ExistCollection(index, collection); !isCollection {
		return nil, fmt.Errorf("Index \"%s\" with collection \"%s\" doesn't exist!", index, collection)
	} else {
		res, err := k.API.Document.Get(index, collection, id, nil)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// func (k *KuzzleAPI) CountDocument(index string, collection string, query json.RawMessage) (int, error) {
// }

// func (k *KuzzleAPI) CreateManyDocument(index string, collection string, data interface{}) (json.RawMessage, error) {
// }

// func (k *KuzzleAPI) DeleteManyDocument(index string, collection string, id []string) ([]string, error) {
// }

// func (k *KuzzleAPI) GetManyDocument(index string, collection string, id []string) (json.RawMessage, error) {
// }

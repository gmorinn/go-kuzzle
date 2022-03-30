package main

import (
	"fmt"
	"kuzzletest/utils"

	"github.com/kuzzleio/sdk-go/kuzzle"
)

type KuzzleAPI struct {
	API *kuzzle.Kuzzle `json:"api"`
}

func (k *KuzzleAPI) createIndex(name string) error {
	if err := utils.CheckErrorIndex(name); err != nil {
		return err
	}
	if exist, _ := k.ExistIndex(name); exist {
		return fmt.Errorf("Index \"%s\" already exists!", name)
	}
	if err := k.API.Index.Create(name, nil); err != nil {
		return err
	}
	return nil
}

func (k *KuzzleAPI) DeleteIndex(name string) error {
	if err := utils.CheckErrorIndex(name); err != nil {
		return err
	}
	if err := k.API.Index.Delete(name, nil); err != nil {
		return fmt.Errorf("Error with index \"%s\"!", name)
	}
	return nil
}

func (k *KuzzleAPI) DeleteManyIndex(indexes []string) error {
	if len(indexes) == 0 {
		return fmt.Errorf("Array is empty!")
	}
	_, err := k.API.Index.MDelete(indexes, nil)
	if err != nil {
		return err
	}
	return nil
}

func (k *KuzzleAPI) ExistIndex(name string) (bool, error) {
	if err := utils.CheckErrorIndex(name); err != nil {
		return false, err
	}
	exist, err := k.API.Index.Exists(name, nil)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (k *KuzzleAPI) listIndex() ([]string, error) {
	res, err := k.API.Index.List(nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

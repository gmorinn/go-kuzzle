package main

import (
	"fmt"

	"github.com/kuzzleio/sdk-go/kuzzle"
)

type KuzzleAPI struct {
	API *kuzzle.Kuzzle `json:"api"`
}

func (k *KuzzleAPI) createIndex(name string) error {
	if exist, _ := k.ExistIndex(name); exist {
		return fmt.Errorf("Index \"%s\" already exists!", name)
	}
	if err := k.API.Index.Create(name, nil); err != nil {
		return err
	}
	return nil
}

func (k *KuzzleAPI) DeleteIndex(name string) error {
	if exist, _ := k.ExistIndex(name); exist {
		if err := k.API.Index.Delete(name, nil); err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("Index \"%s\" doesn't exist!", name)
	}
}

func (k *KuzzleAPI) DeleteManyIndex(indexes []string) error {
	_, err := k.API.Index.MDelete(indexes, nil)
	if err != nil {
		return err
	}
	return nil
}

func (k *KuzzleAPI) ExistIndex(name string) (bool, error) {
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

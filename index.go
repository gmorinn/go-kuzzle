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
		fmt.Printf("Index %s already exists!\n", name)
		return nil
	}
	if err := k.API.Index.Create(name, nil); err != nil {
		fmt.Printf("Index %s not created!\n", name)
		return err
	}
	fmt.Printf("Index %s created!\n", name)
	return nil

}

func (k *KuzzleAPI) DeleteIndex(name string) error {
	if exist, _ := k.ExistIndex(name); !exist {
		fmt.Printf("Index %s doesn't exist!\n", name)
		return nil
	}
	if err := k.API.Index.Delete(name, nil); err != nil {
		fmt.Printf("Index %s not deleted!\n", name)
		return err
	}
	fmt.Printf("Index %s deleted!\n", name)
	return nil
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

package main

import (
	"encoding/json"
	"fmt"
)

func (k *KuzzleAPI) CreateCollection(index, collection string) error {
	if exist, _ := k.ExistCollection(index, collection); exist {
		return fmt.Errorf("Index \"%s\" with collection \"%s\" already exists!", index, collection)
	}
	if err := k.API.Collection.Create(index, collection, nil, nil); err != nil {
		return err
	}
	return nil
}

func (k *KuzzleAPI) DeleteCollection(index, collection string) error {
	if exist, _ := k.ExistCollection(index, collection); exist {
		if err := k.API.Collection.Delete(index, collection, nil); err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("Index \"%s\" with collection \"%s\" doesn't exist!", index, collection)
	}
}

func (k *KuzzleAPI) GetCollection(index, collection string) (json.RawMessage, error) {
	if exist, _ := k.ExistCollection(index, collection); !exist {
		return nil, fmt.Errorf("Index \"%s\" with collection \"%s\" doesn't exist!", index, collection)
	}
	res, err := k.API.Collection.GetMapping(index, collection, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (k *KuzzleAPI) GetCollectionSpecification(index, collection string) (json.RawMessage, error) {
	if exist, _ := k.ExistCollection(index, collection); !exist {
		return nil, fmt.Errorf("Index \"%s\" with collection \"%s\" doesn't exist!", index, collection)
	}
	res, err := k.API.Collection.GetSpecifications(index, collection, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (k *KuzzleAPI) ExistCollection(index, collection string) (bool, error) {
	exist, err := k.API.Collection.Exists(index, collection, nil)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (k *KuzzleAPI) ListCollection(index string) (json.RawMessage, error) {
	if exist, _ := k.ExistIndex(index); !exist {
		return nil, fmt.Errorf("Index \"%s\" doesn't exist!", index)
	}
	res, err := k.API.Collection.List(index, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// func (k *KuzzleAPI) DeleteCollectionSpecifications(index, collection string) error {
// }

// func (k *KuzzleAPI) SearchCollectionSpecifications(index, collection string) error {
// }

// func (k *KuzzleAPI) RefreshCollection(index, collection string) error {
// }

// func (k *KuzzleAPI) TruncateCollection(index, collection string) error {
// }

// func (k *KuzzleAPI) UpdateMappingCollection(index, collection string) error {
// }

// func (k *KuzzleAPI) UpdateSpecificationsCollection(index, collection string) error {
// }

// func (k *KuzzleAPI) ValidateSpecificationsCollection(index, collection string) error {
// }

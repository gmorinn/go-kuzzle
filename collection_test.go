package main

import (
	"encoding/json"
	"fmt"
	"kuzzletest/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCollection(t *testing.T) {
	testAPI.CreateCollection("testindex", "testcollection", nil)
	type args struct {
		index      string
		collection string
		data       json.RawMessage
	}
	tests := []struct {
		name     string
		args     args
		response error
	}{
		{
			name: "Create with empty index",
			args: args{
				index:      "",
				collection: utils.RandStringRunes(8),
				data:       nil,
			},
			response: fmt.Errorf("Index is empty!"),
		},
		{
			name: "Create with empty collection",
			args: args{
				index:      utils.RandStringRunes(8),
				collection: "",
				data:       nil,
			},
			response: fmt.Errorf("Collection is empty!"),
		},
		{
			name: "Create with uppercase index",
			args: args{
				index:      "AZERTYUIOP",
				collection: utils.RandStringRunes(8),
				data:       nil,
			},
			response: fmt.Errorf("Uppercase is not allowed!"),
		},
		{
			name: "Create with uppercase collection",
			args: args{
				index:      utils.RandStringRunes(8),
				collection: "AZERTYUIOP",
				data:       nil,
			},
			response: fmt.Errorf("Uppercase is not allowed!"),
		},
		{
			name: "Create with non-existing index and collection",
			args: args{
				index:      utils.RandStringRunes(8),
				collection: utils.RandStringRunes(8),
				data:       nil,
			},
			response: nil,
		},
		{
			name: "Create with existing index and collection",
			args: args{
				index:      "testindex",
				collection: "testcollection",
				data:       nil,
			},
			response: fmt.Errorf("Index \"testindex\" with collection \"testcollection\" already exists!"),
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			err := testAPI.CreateCollection(v.args.index, v.args.collection, v.args.data)
			require.Equal(t, v.response, err)
		})
	}
}

func TestExistCollection(t *testing.T) {
	testAPI.CreateCollection("testindex", "testcollection", nil)
	type args struct {
		index      string
		collection string
	}
	tests := []struct {
		name     string
		args     args
		response struct {
			bool
			error
		}
	}{
		{
			name: "empty index",
			args: args{
				index:      "",
				collection: utils.RandStringRunes(8),
			},
			response: struct {
				bool
				error
			}{false, nil},
		},
		{
			name: "empty collection",
			args: args{
				index:      utils.RandStringRunes(8),
				collection: "",
			},
			response: struct {
				bool
				error
			}{
				false,
				nil,
			},
		},
		{
			name: "uppercase index",
			args: args{
				index:      "AZERTYUIOP",
				collection: utils.RandStringRunes(8),
			},
			response: struct {
				bool
				error
			}{
				false,
				nil,
			},
		},
		{
			name: "uppercase collection",
			args: args{
				index:      utils.RandStringRunes(8),
				collection: "AZERTYUIOP",
			},
			response: struct {
				bool
				error
			}{
				false,
				nil,
			},
		},
		{
			name: "non-existing index and collection",
			args: args{
				index:      utils.RandStringRunes(8),
				collection: utils.RandStringRunes(8),
			},
			response: struct {
				bool
				error
			}{
				false,
				nil,
			},
		},
		{
			name: "existing index and collection",
			args: args{
				index:      "testindex",
				collection: "testcollection",
			},
			response: struct {
				bool
				error
			}{
				true,
				nil,
			},
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res, _ := testAPI.ExistCollection(v.args.index, v.args.collection)
			require.Equal(t, v.response.bool, res)
		})
	}
}

func TestDeleteCollection(t *testing.T) {
	testAPI.CreateCollection("testindex", "testcollection", nil)
	randomIndex := utils.RandStringRunes(8)
	randomCollection := utils.RandStringRunes(8)
	type args struct {
		index      string
		collection string
	}
	tests := []struct {
		name     string
		args     args
		response error
	}{
		{
			name: "non-existing index and collection",
			args: args{
				index:      randomIndex,
				collection: randomCollection,
			},
			response: fmt.Errorf("Index \"%s\" with collection \"%s\" doesn't exist!", randomIndex, randomCollection),
		},
		{
			name: "existing index and collection",
			args: args{
				index:      "testindex",
				collection: "testcollection",
			},
			response: nil,
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			err := testAPI.DeleteCollection(v.args.index, v.args.collection)
			require.Equal(t, v.response, err)
		})
	}
}

func TestGetCollection(t *testing.T) {
	removeAllIndex(testAPI)
	var userTable CollectionUser = CollectionUser{
		Properties: User{
			Username:    Type{Keyword},
			Birthday:    Type{Date},
			Age:         Type{Integer},
			Description: Type{Text},
			GPA:         Type{Float},
		},
	}
	resJson := utils.GetFormatJSON(userTable)
	testAPI.CreateCollection("testindex", "testcollection", resJson)

	randomIndex := utils.RandStringRunes(8)
	randomCollection := utils.RandStringRunes(8)
	type args struct {
		index      string
		collection string
	}
	tests := []struct {
		name     string
		args     args
		response struct {
			json.RawMessage
			error
		}
	}{
		{
			name: "non-existing index and collection",
			args: args{
				index:      randomIndex,
				collection: randomCollection,
			},
			response: struct {
				json.RawMessage
				error
			}{
				nil,
				fmt.Errorf("Index \"%s\" with collection \"%s\" doesn't exist!", randomIndex, randomCollection),
			},
		},
		{
			name: "existing index and collection",
			args: args{
				index:      "testindex",
				collection: "testcollection",
			},
			response: struct {
				json.RawMessage
				error
			}{
				resJson,
				nil,
			}},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			_, err := testAPI.GetCollection(v.args.index, v.args.collection)
			require.Equal(t, v.response.error, err)
			// require.Equal(t, string(v.response.RawMessage), string(res))
		})
	}
}
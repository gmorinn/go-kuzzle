package main

import (
	"fmt"
	"kuzzletest/utils"
	"testing"

	"github.com/kuzzleio/sdk-go/types"
	"github.com/stretchr/testify/require"
)

func TestCreateIndex(t *testing.T) {
	tests := []struct {
		name     string
		args     string
		response error
	}{
		{
			name:     "Create with non-existing index",
			args:     "test",
			response: nil,
		},
		{
			name:     "Create with existing index",
			args:     "test",
			response: fmt.Errorf("Index \"test\" already exists!"),
		},
		{
			name:     "Create with empty index",
			args:     "",
			response: types.KuzzleError(types.KuzzleError{Message: "Index.Create: index required", Stack: "", Status: 400}),
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			err := testAPI.createIndex(v.args)
			require.Equal(t, v.response, err)
		})
	}
}

func TestDeleteIndex(t *testing.T) {
	var randomIndex string = utils.RandStringRunes(8)
	var createNewIndex error = testAPI.createIndex(utils.RandStringRunes(8))
	require.NoError(t, createNewIndex)

	tests := []struct {
		name     string
		args     string
		response error
	}{
		{
			name:     "Delete with non-existing index",
			args:     randomIndex,
			response: fmt.Errorf("Index \"%s\" doesn't exist!", randomIndex),
		},
		{
			name:     "Delete with existing index",
			args:     "test",
			response: nil,
		},
		{
			name:     "Delete with empty index",
			args:     "",
			response: fmt.Errorf("Index \"\" doesn't exist!"),
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			err := testAPI.DeleteIndex(v.args)
			require.Equal(t, v.response, err)
		})
	}
}

func TestDeleteManyIndex(t *testing.T) {
	
}

func TestExistIndex(t *testing.T) {
	var randomIndex string = utils.RandStringRunes(8)
	var createNewIndex error = testAPI.createIndex(randomIndex)
	require.NoError(t, createNewIndex)

	tests := []struct {
		name     string
		args     string
		response struct {
			bool
			error
		}
	}{
		{
			name: "Exist index",
			args: randomIndex,
			response: struct {
				bool
				error
			}{true, nil},
		},
		{
			name: "Not index",
			args: utils.RandStringRunes(8),
			response: struct {
				bool
				error
			}{
				false,
				nil,
			},
		},
		{
			name: "Empty index",
			args: "",
			response: struct {
				bool
				error
			}{
				false,
				types.KuzzleError(types.KuzzleError{Message: "Index.Exists: index required", Stack: "", Status: 400}),
			},
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			exist, err := testAPI.ExistIndex(v.args)
			require.Equal(t, v.response.error, err)
			require.Equal(t, exist, exist)
		})
	}
}

func TestListIndex(t *testing.T) {

}

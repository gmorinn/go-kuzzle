package main

import (
	"fmt"
	"kuzzletest/utils"
	"testing"

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
			response: fmt.Errorf("Index is empty!"),
		},
		{
			name:     "Create index with Uppercase",
			args:     "AZERTYUIO",
			response: fmt.Errorf("Uppercase is not allowed! \"AZERTYUIO\""),
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
	var newIndex string = utils.RandStringRunes(8)
	var createNewIndex error = testAPI.createIndex(newIndex)
	require.NoError(t, createNewIndex)

	tests := []struct {
		name     string
		args     string
		response error
	}{
		{
			name:     "Delete with non-existing index",
			args:     randomIndex,
			response: fmt.Errorf("Error with index \"%s\"!", randomIndex),
		},
		{
			name:     "Delete with existing index",
			args:     newIndex,
			response: nil,
		},
		{
			name:     "Delete with empty index",
			args:     "",
			response: fmt.Errorf("Index is empty!"),
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
	var tab = []string{utils.RandStringRunes(8), utils.RandStringRunes(8), utils.RandStringRunes(8)}
	for _, v := range tab {
		testAPI.createIndex(v)
	}
	tests := []struct {
		name     string
		args     []string
		response error
	}{
		{
			name:     "Delete with non-existing indexes",
			args:     []string{utils.RandStringRunes(8), utils.RandStringRunes(12)},
			response: nil,
		},
		{
			name:     "Delete with existing indexes",
			args:     tab,
			response: nil,
		},
		{
			name:     "Delete no index",
			args:     []string{},
			response: fmt.Errorf("Array is empty!"),
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			err := testAPI.DeleteManyIndex(v.args)
			require.Equal(t, v.response, err)
		})
	}
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
				fmt.Errorf("Index is empty!"),
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
	removeAllIndex(testAPI)
	tests := []struct {
		name     string
		response struct {
			res []string
			err error
		}
	}{
		{
			name: "No indexes",
			response: struct {
				res []string
				err error
			}{
				[]string{},
				nil,
			},
		},
		{
			name: "indexes",
			response: struct {
				res []string
				err error
			}{
				[]string{utils.RandStringRunes(8), utils.RandStringRunes(10)},
				nil,
			},
		},
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if v.name == "indexes" {
				testAPI.createIndex(utils.RandStringRunes(8))
				testAPI.createIndex(utils.RandStringRunes(10))
			}
			res, err := testAPI.listIndex()
			require.GreaterOrEqual(t, len(res), len(v.response.res))
			require.Equal(t, err, v.response.err)
		})
	}
}

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

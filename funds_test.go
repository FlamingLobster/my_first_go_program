package main

import (
	"encoding/json"
	"koho/velocity"
	"testing"
)

func TestRepeatedId(t *testing.T) {
	loadFund := velocity.LoadFund{
		Id: 33,
		Amount: velocity.Dollar{
			Amount: 2352,
		},
	}
	if input, err := json.Marshal(loadFund); err != nil {
		t.Error(err)
	} else {
		velocity.Allowed(string(input))
		if isAllowed, _ := velocity.Allowed(string(input)); isAllowed {
			t.Error("failed to stop duplicate load fund transaction")
		}
	}
}

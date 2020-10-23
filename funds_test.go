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
	response := velocity.Response{
		Id:         loadFund.Id,
		CustomerId: loadFund.CustomerId,
		Accepted:   false,
	}
	input, err := json.Marshal(loadFund)
	if err != nil {
		t.Error(err)
	}
	_, _ = velocity.Allowed(string(input))
	err, output := velocity.Allowed(string(input))
	if err != nil {
		t.Error(err)
	}
	expected, err := json.Marshal(response)
	if err != nil {
		t.Error(err)
	}
	if output != string(expected) {
		t.Error("failed to stop duplicate load fund transaction")
	}
}

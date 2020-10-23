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

func TestSameTransactionIdDifferentCustomerId(t *testing.T) {
	loadFund1 := velocity.LoadFund{
		Id:         33,
		CustomerId: 44,
		Amount: velocity.Dollar{
			Amount: 2352,
		},
	}
	loadFund2 := velocity.LoadFund{
		Id:         loadFund1.Id,
		CustomerId: loadFund1.CustomerId + 1,
		Amount: velocity.Dollar{
			Amount: 2352,
		},
	}
	response := velocity.Response{
		Id:         loadFund2.Id,
		CustomerId: loadFund2.CustomerId,
		Accepted:   true,
	}
	input1, err := json.Marshal(loadFund1)
	if err != nil {
		t.Error(err)
	}
	input2, err := json.Marshal(loadFund2)
	if err != nil {
		t.Error(err)
	}

	_, _ = velocity.Allowed(string(input1))
	err, output := velocity.Allowed(string(input2))
	if err != nil {
		t.Error(err)
	}
	expected, err := json.Marshal(response)
	if err != nil {
		t.Error(err)
	}
	if output != string(expected) {
		t.Error("failed to accepted same transaction id but different customer id")
	}
}

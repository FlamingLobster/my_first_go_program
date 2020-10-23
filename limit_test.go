package main

import (
	"encoding/json"
	"koho/velocity"
	"testing"
	"time"
)

func TestRepeatedId(t *testing.T) {
	velocity.Reset()

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
	velocity.Reset()

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

func TestRejectSingleTransactionGreaterThanDailyLimit(t *testing.T) {
	velocity.Reset()

	loadFund := velocity.LoadFund{
		Id:         33,
		CustomerId: 44,
		Amount: velocity.Dollar{
			Amount: velocity.DailyFundLimit + 1,
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

	err, output := velocity.Allowed(string(input))
	if err != nil {
		t.Error(err)
	}
	expected, err := json.Marshal(response)
	if err != nil {
		t.Error(err)
	}
	if output != string(expected) {
		t.Error("failed reject transaction with amount greater than daily limit")
	}
}

func TestRejectMultipleTransactionAddUpToGreaterThanDailyLimit(t *testing.T) {
	velocity.Reset()

	loadFund1 := velocity.LoadFund{
		Id:         33,
		CustomerId: 44,
		Amount: velocity.Dollar{
			Amount: velocity.DailyFundLimit / 2,
		},
		Timestamp: time.Date(
			2020,
			10,
			10,
			10,
			10,
			10,
			0,
			time.FixedZone("UTC", 0)),
	}
	loadFund2 := velocity.LoadFund{
		Id:         34,
		CustomerId: 44,
		Amount: velocity.Dollar{
			Amount: velocity.DailyFundLimit / 2,
		},
		Timestamp: time.Date(
			2020,
			10,
			10,
			10,
			15,
			10,
			0,
			time.FixedZone("UTC", 0)),
	}
	response2 := velocity.Response{
		Id:         loadFund2.Id,
		CustomerId: loadFund2.CustomerId,
		Accepted:   true,
	}
	loadFund3 := velocity.LoadFund{
		Id:         35,
		CustomerId: 44,
		Amount: velocity.Dollar{
			Amount: velocity.DailyFundLimit / 2,
		},
		Timestamp: time.Date(
			2020,
			10,
			10,
			10,
			20,
			10,
			0,
			time.FixedZone("UTC", 0)),
	}
	response3 := velocity.Response{
		Id:         loadFund3.Id,
		CustomerId: loadFund3.CustomerId,
		Accepted:   false,
	}

	input1, err := json.Marshal(loadFund1)
	if err != nil {
		t.Error(err)
	}
	err, _ = velocity.Allowed(string(input1))
	if err != nil {
		t.Error(err)
	}

	input2, err := json.Marshal(loadFund2)
	if err != nil {
		t.Error(err)
	}
	err, output2 := velocity.Allowed(string(input2))
	if err != nil {
		t.Error(err)
	}
	expected2, err := json.Marshal(response2)
	if err != nil {
		t.Error(err)
	}
	if output2 != string(expected2) {
		t.Error("failed to setup precondition for test")
	}

	input3, err := json.Marshal(loadFund3)
	if err != nil {
		t.Error(err)
	}
	err, output3 := velocity.Allowed(string(input3))
	if err != nil {
		t.Error(err)
	}
	expected3, err := json.Marshal(response3)
	if err != nil {
		t.Error(err)
	}
	if output3 != string(expected3) {
		t.Error("failed reject transaction with amount greater than daily limit")
	}
}

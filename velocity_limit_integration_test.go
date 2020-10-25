package main

import (
	"encoding/json"
	"koho/velocity"
	"testing"
	"time"
)

/**
This came before I understood the typical go style table tests. This can be turned into table tests
*/
func TestRepeatedId(t *testing.T) {
	Reset()

	loadFund := velocity.Funds{
		Id:         33,
		CustomerId: 44,
		Dollar: velocity.Dollar{
			Amount: 2352,
		},
	}

	input, err := json.Marshal(loadFund)
	if err != nil {
		t.Error(err)
	}
	_, _, _ = Allowed(string(input))
	err, action, _ := Allowed(string(input))
	if err != nil {
		t.Error(err)
	}
	if action != velocity.Ignore {
		t.Error("failed to stop duplicate load fund transaction")
	}

}

func TestSameTransactionIdDifferentCustomerId(t *testing.T) {
	Reset()

	loadFund1 := velocity.Funds{
		Id:         33,
		CustomerId: 44,
		Dollar: velocity.Dollar{
			Amount: 2352,
		},
	}
	loadFund2 := velocity.Funds{
		Id:         loadFund1.Id,
		CustomerId: loadFund1.CustomerId + 1,
		Dollar: velocity.Dollar{
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

	_, _, _ = Allowed(string(input1))
	err, _, output := Allowed(string(input2))
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
	Reset()

	loadFund := velocity.Funds{
		Id:         33,
		CustomerId: 44,
		Dollar: velocity.Dollar{
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

	err, _, output := Allowed(string(input))
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
	Reset()

	loadFund1 := velocity.Funds{
		Id:         33,
		CustomerId: 44,
		Dollar: velocity.Dollar{
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
	loadFund2 := velocity.Funds{
		Id:         34,
		CustomerId: 44,
		Dollar: velocity.Dollar{
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
	loadFund3 := velocity.Funds{
		Id:         35,
		CustomerId: 44,
		Dollar: velocity.Dollar{
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
	err, _, _ = Allowed(string(input1))
	if err != nil {
		t.Error(err)
	}

	input2, err := json.Marshal(loadFund2)
	if err != nil {
		t.Error(err)
	}
	err, _, output2 := Allowed(string(input2))
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
	err, _, output3 := Allowed(string(input3))
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

func TestReject4thLoadOfTheDay(t *testing.T) {
	Reset()

	loadFund1 := velocity.Funds{
		Id:         33,
		CustomerId: 44,
		Dollar: velocity.Dollar{
			Amount: velocity.DailyFundLimit / 10,
		},
		Timestamp: time.Date(
			2020,
			10,
			10,
			10,
			10,
			12,
			0,
			time.FixedZone("UTC", 0)),
	}
	loadFund2 := velocity.Funds{
		Id:         34,
		CustomerId: 44,
		Dollar: velocity.Dollar{
			Amount: velocity.DailyFundLimit / 10,
		},
		Timestamp: time.Date(
			2020,
			10,
			10,
			10,
			10,
			13,
			0,
			time.FixedZone("UTC", 0)),
	}
	loadFund3 := velocity.Funds{
		Id:         35,
		CustomerId: 44,
		Dollar: velocity.Dollar{
			Amount: velocity.DailyFundLimit / 10,
		},
		Timestamp: time.Date(
			2020,
			10,
			10,
			10,
			10,
			14,
			0,
			time.FixedZone("UTC", 0)),
	}
	response3 := velocity.Response{
		Id:         loadFund3.Id,
		CustomerId: loadFund3.CustomerId,
		Accepted:   true,
	}
	loadFund4 := velocity.Funds{
		Id:         36,
		CustomerId: 44,
		Dollar: velocity.Dollar{
			Amount: velocity.DailyFundLimit / 10,
		},
		Timestamp: time.Date(
			2020,
			10,
			10,
			10,
			10,
			15,
			0,
			time.FixedZone("UTC", 0)),
	}
	response4 := velocity.Response{
		Id:         loadFund4.Id,
		CustomerId: loadFund4.CustomerId,
		Accepted:   false,
	}

	input1, err := json.Marshal(loadFund1)
	if err != nil {
		t.Error(err)
	}
	err, _, _ = Allowed(string(input1))
	if err != nil {
		t.Error(err)
	}

	input2, err := json.Marshal(loadFund2)
	if err != nil {
		t.Error(err)
	}
	err, _, _ = Allowed(string(input2))
	if err != nil {
		t.Error(err)
	}

	input3, err := json.Marshal(loadFund3)
	if err != nil {
		t.Error(err)
	}
	err, _, output3 := Allowed(string(input3))
	if err != nil {
		t.Error(err)
	}
	expected3, err := json.Marshal(response3)
	if err != nil {
		t.Error(err)
	}
	if output3 != string(expected3) {
		t.Error("failed to setup precondition for test")
	}

	input4, err := json.Marshal(loadFund4)
	if err != nil {
		t.Error(err)
	}
	err, _, output4 := Allowed(string(input4))
	if err != nil {
		t.Error(err)
	}
	expected4, err := json.Marshal(response4)
	if err != nil {
		t.Error(err)
	}
	if output4 != string(expected4) {
		t.Error("failed to reject 4th transaction on same day")
	}
}

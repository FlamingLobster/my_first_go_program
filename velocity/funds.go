package velocity

import (
	"strconv"
	"strings"
	"time"
)

type LoadFund struct {
	Id         int       `json:"id,string"`
	CustomerId int       `json:"customer_id,string"`
	Amount     Dollar    `json:"load_amount"`
	Timestamp  time.Time `json:"time"`
}

type Dollar struct {
	Amount int
}

func (d *Dollar) MarshalJSON() ([]byte, error) {
	var amount = strconv.Itoa(d.Amount)
	if d.Amount >= 100 {
		return []byte("$" + amount[:len(amount)-2] + "." + amount[len(amount)-2:]), nil
	} else if d.Amount >= 10 {
		return []byte("$0." + amount), nil
	} else {
		return []byte("$0.0" + amount), nil
	}
}

func (d *Dollar) UnmarshalJSON(data []byte) error {
	raw := string(data)
	amount := strings.ReplaceAll(raw, "$", "")
	amount = strings.ReplaceAll(amount, ".", "")
	d.Amount, _ = strconv.Atoi(amount)
	return nil
}

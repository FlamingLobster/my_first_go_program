package velocity

import (
	"strconv"
	"time"
)

type LoadFund struct {
	Id         int       `json:"id,string"`
	CustomerId int       `json:"customer_id,string"`
	Amount     Dollar    `json:"load_amount"`
	Timestamp  time.Time `json:"time"`
}

type Dollar struct {
	amount int
}

func (d *Dollar) UnmarshalJSON(data []byte) error {
	raw := string(data)
	var amount = raw[2 : len(raw)-4]
	amount += raw[len(raw)-3 : len(raw)-1]
	d.amount, _ = strconv.Atoi(amount)
	return nil
}

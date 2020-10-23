package velocity

import (
	"time"
)

type Tuple struct {
	id         int
	customerId int
}

func KeyOf(id int, customerId int) Tuple {
	return Tuple{
		id:         id,
		customerId: customerId,
	}
}

type DailyTransactionKey struct {
	customerId int
	datetime   time.Time
}

func TimeKeyOf(customerId int, datetime time.Time) DailyTransactionKey {
	return DailyTransactionKey{
		customerId: customerId,
		datetime:   datetime,
	}
}

type WeeklyTransactionKey struct {
	customerId int
	week       Tuple
}

func WeekKeyOf(customerid int, tuple Tuple) WeeklyTransactionKey {
	return WeeklyTransactionKey{
		customerId: customerid,
		week:       tuple,
	}
}

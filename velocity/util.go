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
	CustomerId int
	datetime   time.Time
}

func TimeKeyOf(customerId int, datetime time.Time) DailyTransactionKey {
	return DailyTransactionKey{
		CustomerId: customerId,
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

func ToStartOfDay(unrounded time.Time) time.Time {
	utcUnrounded := unrounded.UTC()
	return time.Date(unrounded.Year(), unrounded.Month(), unrounded.Day(), 0, 0, 0, 0, utcUnrounded.Location())
}

func ToStartOfWeek(unrounded time.Time) Tuple {
	utcUnrounded := unrounded.UTC()
	year, week := utcUnrounded.ISOWeek()
	return KeyOf(year, week)
}

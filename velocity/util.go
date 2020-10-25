package velocity

import (
	"time"
)

type UniqueTransactionKey struct {
	id         int
	customerId int
}

func KeyOf(id int, customerId int) UniqueTransactionKey {
	return UniqueTransactionKey{
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
	week       UniqueTransactionKey
}

func WeekKeyOf(customerId int, tuple UniqueTransactionKey) WeeklyTransactionKey {
	return WeeklyTransactionKey{
		customerId: customerId,
		week:       tuple,
	}
}

func ToStartOfDay(unrounded time.Time) time.Time {
	utcUnrounded := unrounded.UTC()
	return time.Date(
		unrounded.Year(),
		unrounded.Month(),
		unrounded.Day(),
		0,
		0,
		0,
		0,
		utcUnrounded.Location(),
	)
}

func ToStartOfWeek(unrounded time.Time) UniqueTransactionKey {
	utcUnrounded := unrounded.UTC()
	year, week := utcUnrounded.ISOWeek()
	return KeyOf(year, week)
}

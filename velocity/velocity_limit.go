package velocity

import (
	"encoding/json"
)

var limits = getLimits()

func Reset() {
	limits = getLimits()
}

func Allowed(event string) (error, string) {
	var loadFund LoadFund
	if err := json.Unmarshal([]byte(event), &loadFund); err != nil {
		return err, ""
	} else {
		if limits.allowed(loadFund) {
			if output, err := json.Marshal(Accepted(&loadFund)); err != nil {
				return err, ""
			} else {
				return nil, string(output)
			}
		} else {
			if output, err := json.Marshal(Denied(&loadFund)); err != nil {
				return err, ""
			} else {
				return nil, string(output)
			}
		}
	}
}

const (
	DailyFundLimit  int = 500000
	WeeklyFundLimit int = 2000000
)

type Limits struct {
	userTransactions       map[Tuple]bool
	userDailyTransactions  map[DailyTransactionKey]int
	userWeeklyTransactions map[WeeklyTransactionKey]int
}

func (l Limits) allowed(funds LoadFund) bool {
	if _, present := l.userTransactions[KeyOf(funds.Id, funds.CustomerId)]; present {
		return false
	}
	l.userTransactions[KeyOf(funds.Id, funds.CustomerId)] = true

	isAllowedByDailyLimit := l.allowedByDailyLimit(funds)
	isAllowedByWeeklyLimit := l.allowedByWeeklyLimit(funds)

	if isAllowedByDailyLimit && isAllowedByWeeklyLimit {
		l.update(funds)
		return true
	} else {
		return false
	}
}

func (l Limits) allowedByDailyLimit(funds LoadFund) bool {
	if funds.Amount.Amount > DailyFundLimit {
		return false
	}

	startOfDay := ToStartOfDay(funds.Timestamp)
	if balance, present := l.userDailyTransactions[TimeKeyOf(funds.CustomerId, startOfDay)]; present {
		if balance+funds.Amount.Amount > DailyFundLimit {
			return false
		}
	}
	return true
}

func (l Limits) allowedByWeeklyLimit(funds LoadFund) bool {
	if funds.Amount.Amount > WeeklyFundLimit {
		return false
	}

	week := WeekKeyOf(funds.CustomerId, ToStartOfWeek(funds.Timestamp))
	if balance, present := l.userWeeklyTransactions[week]; present {
		if balance+funds.Amount.Amount > WeeklyFundLimit {
			return false
		}
	}
	return true
}

func (l Limits) update(funds LoadFund) {
	startOfDay := ToStartOfDay(funds.Timestamp)
	if balance, present := l.userDailyTransactions[TimeKeyOf(funds.CustomerId, startOfDay)]; present {
		l.userDailyTransactions[TimeKeyOf(funds.CustomerId, startOfDay)] = balance + funds.Amount.Amount
	} else {
		l.userDailyTransactions[TimeKeyOf(funds.CustomerId, startOfDay)] = funds.Amount.Amount
	}

	week := WeekKeyOf(funds.CustomerId, ToStartOfWeek(funds.Timestamp))
	if balance, present := l.userWeeklyTransactions[week]; present {
		l.userWeeklyTransactions[week] = balance + funds.Amount.Amount
	} else {
		l.userWeeklyTransactions[week] = funds.Amount.Amount
	}
}

func getLimits() *Limits {
	limits := Limits{
		userTransactions:       make(map[Tuple]bool),
		userDailyTransactions:  make(map[DailyTransactionKey]int),
		userWeeklyTransactions: make(map[WeeklyTransactionKey]int),
	}
	return &limits
}

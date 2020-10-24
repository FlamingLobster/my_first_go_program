package velocity

import (
	"encoding/json"
	"errors"
)

var limits = GetLimits()

func Reset() {
	limits = GetLimits()
}

func Allowed(event string) (error, int, string) {
	var loadFund LoadFund
	if err := json.Unmarshal([]byte(event), &loadFund); err != nil {
		return err, Ignore, ""
	} else {
		action := limits.allowed(loadFund)
		switch action {
		case Accept:
			if output, err := json.Marshal(Accepted(&loadFund)); err != nil {
				return err, Ignore, ""
			} else {
				return nil, Accept, string(output)
			}
		case Deny:
			if output, err := json.Marshal(Denied(&loadFund)); err != nil {
				return err, Ignore, ""
			} else {
				return nil, Deny, string(output)
			}
		case Ignore:
			return nil, Ignore, ""
		}
	}
	return errors.New("should not get here"), Ignore, ""
}

const (
	DailyDistinctLimit int = 3
	DailyFundLimit     int = 500000
	WeeklyFundLimit    int = 2000000
)

type Limits struct {
	userTransactions       map[UniqueTransactionKey]bool
	userDailyTransactions  map[DailyTransactionKey]*BalanceAndCount
	userWeeklyTransactions map[WeeklyTransactionKey]int
}

type BalanceAndCount struct {
	balance int
	count   int
}

func (b *BalanceAndCount) addBalance(amount int) *BalanceAndCount {
	b.balance += amount
	b.increment()
	return b
}

func (b *BalanceAndCount) increment() {
	b.count = b.count + 1
}

func (l Limits) allowed(funds LoadFund) int {
	if _, present := l.userTransactions[KeyOf(funds.Id, funds.CustomerId)]; present {
		return Ignore
	}
	l.userTransactions[KeyOf(funds.Id, funds.CustomerId)] = true

	isAllowedByDailyLimit := l.allowedByDailyLimit(funds)
	isAllowedByWeeklyLimit := l.allowedByWeeklyLimit(funds)

	if isAllowedByDailyLimit && isAllowedByWeeklyLimit {
		l.update(funds)
		return Accept
	} else {
		return Deny
	}
}

func (l Limits) allowedByDailyLimit(funds LoadFund) bool {
	if funds.Amount.Amount > DailyFundLimit {
		return false
	}

	startOfDay := ToStartOfDay(funds.Timestamp)
	if balance, present := l.userDailyTransactions[TimeKeyOf(funds.CustomerId, startOfDay)]; present {
		if balance.balance+funds.Amount.Amount > DailyFundLimit || balance.count == DailyDistinctLimit {
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
		l.userDailyTransactions[TimeKeyOf(funds.CustomerId, startOfDay)] = balance.addBalance(funds.Amount.Amount)
	} else {
		l.userDailyTransactions[TimeKeyOf(funds.CustomerId, startOfDay)] = &BalanceAndCount{funds.Amount.Amount, 1}
	}

	week := WeekKeyOf(funds.CustomerId, ToStartOfWeek(funds.Timestamp))
	if balance, present := l.userWeeklyTransactions[week]; present {
		l.userWeeklyTransactions[week] = balance + funds.Amount.Amount
	} else {
		l.userWeeklyTransactions[week] = funds.Amount.Amount
	}
}

func GetLimits() *Limits {
	limits := Limits{
		userTransactions:       make(map[UniqueTransactionKey]bool),
		userDailyTransactions:  make(map[DailyTransactionKey]*BalanceAndCount),
		userWeeklyTransactions: make(map[WeeklyTransactionKey]int),
	}
	return &limits
}

const (
	Accept = iota
	Deny
	Ignore
)

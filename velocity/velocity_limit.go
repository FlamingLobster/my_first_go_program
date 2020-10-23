package velocity

import (
	"encoding/json"
)

var limits = getLimits()

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

type Limits struct {
	userTransactions map[Tuple]bool
}

func (l Limits) allowed(funds LoadFund) bool {
	if _, present := l.userTransactions[KeyOf(funds.Id, funds.CustomerId)]; present {
		return false
	}
	l.userTransactions[KeyOf(funds.Id, funds.CustomerId)] = true
	return true
}

func getLimits() *Limits {
	limits := Limits{userTransactions: make(map[Tuple]bool)}
	return &limits
}

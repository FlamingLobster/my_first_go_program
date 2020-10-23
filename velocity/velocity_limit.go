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
	uniqueTransaction map[int]bool
}

func (l Limits) allowed(funds LoadFund) bool {
	if _, present := l.uniqueTransaction[funds.Id]; present {
		return false
	}
	l.uniqueTransaction[funds.Id] = true
	return true
}

func getLimits() *Limits {
	limits := Limits{uniqueTransaction: make(map[int]bool)}
	return &limits
}

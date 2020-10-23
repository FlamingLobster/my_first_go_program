package velocity

import "encoding/json"

var limits = getLimits()

func Allowed(event string) (bool, string) {
	var loadFund LoadFund
	if err := json.Unmarshal([]byte(event), &loadFund); err != nil {
		return false, ""
	} else {
		if limits.allowed(loadFund) {
			return true, respond(loadFund)
		} else {
			return false, ""
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

func respond(_ LoadFund) string {
	return "yes"
}

func getLimits() *Limits {
	limits := Limits{uniqueTransaction: make(map[int]bool)}
	return &limits
}

package velocity

import "encoding/json"

func Process(event string) (bool, string) {
	var loadFund LoadFund
	if err := json.Unmarshal([]byte(event), &loadFund); err != nil {
		return false, ""
	} else {
		if Allowed(loadFund) {
			return true, Respond(loadFund)
		} else {
			return false, ""
		}
	}
}

type Limits struct {
}

func Respond(_ LoadFund) string {
	return ""
}

func Allowed(_ LoadFund) bool {
	return true
}

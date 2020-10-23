package velocity

type Response struct {
	Id         int  `json:"id,string"`
	CustomerId int  `json:"customer_id,string"`
	Accepted   bool `json:"accepted"`
}

func Accepted(loadFund *LoadFund) *Response {
	r := Response{
		Id:         loadFund.Id,
		CustomerId: loadFund.CustomerId,
		Accepted:   true,
	}
	return &r
}

func Denied(loadFund *LoadFund) *Response {
	r := Response{
		Id:         loadFund.Id,
		CustomerId: loadFund.CustomerId,
		Accepted:   false,
	}
	return &r
}

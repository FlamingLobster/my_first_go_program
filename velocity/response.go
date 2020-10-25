package velocity

type Response struct {
	Id         int  `json:"id,string"`
	CustomerId int  `json:"customer_id,string"`
	Accepted   bool `json:"accepted"`
}

func Accepted(funds *Funds) *Response {
	r := Response{
		Id:         funds.Id,
		CustomerId: funds.CustomerId,
		Accepted:   true,
	}
	return &r
}

func Denied(funds *Funds) *Response {
	r := Response{
		Id:         funds.Id,
		CustomerId: funds.CustomerId,
		Accepted:   false,
	}
	return &r
}

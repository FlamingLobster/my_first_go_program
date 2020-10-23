package velocity

type Response struct {
	Id         int  `json:"id"`
	CustomerId int  `json:"customer_id"`
	Accepted   bool `json:"accepted"`
}

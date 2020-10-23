package velocity

type Response struct {
	Id         int  `json:"id,string"`
	CustomerId int  `json:"customer_id,string"`
	Accepted   bool `json:"accepted"`
}

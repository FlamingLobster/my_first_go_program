package velocity

type Tuple struct {
	id         int
	customerId int
}

func KeyOf(id int, customerId int) Tuple {
	return Tuple{
		id:         id,
		customerId: customerId,
	}
}

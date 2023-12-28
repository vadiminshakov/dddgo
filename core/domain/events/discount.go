package events

type BasketDiscountAdded struct {
	Discount float64
}

func (e *BasketDiscountAdded) Bytes() []byte {
	// TODO: implement

	return []byte{}
}

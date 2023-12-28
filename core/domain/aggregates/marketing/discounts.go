package marketing

const (
	bigPrice      = 1000
	bigItemsCount = 10
)

func DiscountForBigOrders(totalPrice float64, itemsCount int64) float64 {
	if totalPrice > bigPrice && itemsCount > bigItemsCount {
		return 0.9 * totalPrice
	}

	return 1
}

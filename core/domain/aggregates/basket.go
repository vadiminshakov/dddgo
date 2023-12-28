package aggregates

import (
	"errors"
	"examplegood/core/domain/aggregates/marketing"
	"examplegood/core/domain/events"
	"examplegood/core/domain/vos"
)

var (
	ErrFullBasket = errors.New("basket is full")
)

const (
	maxBasketItems  = 30
	maxBasketWeight = 10000 // 10 кг
)

type event interface {
	Bytes() []byte
}

type Basket struct {
	ID          int64
	Items       []vos.BasketItem
	TotalWeight int64

	Events []event
}

func NewBasket(id int64, items []vos.BasketItem) *Basket {
	return &Basket{ID: id, Items: items}
}

func (b *Basket) AddItem(item vos.BasketItem) error {
	if len(b.Items) == maxBasketItems {
		return ErrFullBasket
	}

	if b.TotalWeight+item.Weight > maxBasketWeight {
		return ErrFullBasket
	}

	for i, itemFromBasket := range b.Items {
		if itemFromBasket.GoodID == item.GoodID {
			b.Items[i].Quantity += item.Quantity
		}
	}

	b.Items = append(b.Items, vos.BasketItem{GoodID: item.GoodID, Quantity: item.Quantity})
	b.TotalWeight += item.Weight

	return nil
}

func (b *Basket) RemoveItem(goodID, quantity int64) {
	for i, item := range b.Items {
		if item.GoodID == goodID {
			if item.Quantity <= quantity {
				b.Items = append(b.Items[:i], b.Items[i+1:]...)
				return
			}
			b.Items[i].Quantity -= quantity
			return
		}
	}
}

func (b *Basket) TotalPrice() float64 {
	var totalPrice float64
	for _, item := range b.Items {
		totalPrice += item.Price * float64(item.Quantity)
	}

	newPrice := marketing.DiscountForBigOrders(totalPrice, int64(len(b.Items)))
	b.Events = append(b.Events, &events.BasketDiscountAdded{Discount: totalPrice - newPrice})
	totalPrice = newPrice

	return totalPrice
}

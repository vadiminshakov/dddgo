package repository

import (
	"examplegood/core/domain/vos"
)

type ItemsRepository struct {
}

func (r *ItemsRepository) GetByBasketID(id int64) (*vos.BasketItem, error) {
	// TODO: implement

	return &vos.BasketItem{}, nil
}

func (r *ItemsRepository) Save(basket *vos.BasketItem) error {
	// TODO: implement

	return nil
}

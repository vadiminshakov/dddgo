package repository

import "examplegood/core/domain/aggregates"

type BasketRepository struct {
}

func (r *BasketRepository) GetByID(id int64) (*aggregates.Basket, error) {
	// TODO: implement

	return &aggregates.Basket{}, nil
}

func (r *BasketRepository) Save(basket *aggregates.Basket) error {
	// TODO: implement

	return nil
}

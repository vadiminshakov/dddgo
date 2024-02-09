package services

import (
	"context"
	"testing"
	"time"

	"github.com/vadiminshakov/dddgo/repository"

	"github.com/vadiminshakov/dddgo/core/domain/aggregates"

	"github.com/stretchr/testify/require"
	"github.com/vadiminshakov/dddgo/core/domain/vos"

	"github.com/stretchr/testify/mock"
	"github.com/vadiminshakov/dddgo/pkg/mocks"
)

// TestBasketService_AddItemWithTx is a unit test for BasketService.AddItemWithTx.
// It demonstrates how to test a service with transaction method.
func TestBasketService_AddItemWithTx(t *testing.T) {
	repoBasket := &mocks.BasketRepository{}
	repoBasket.On("Save", mock.Anything).Return(nil)
	repoBasket.On("GetByID", mock.Anything).Return(&aggregates.Basket{
		ID:          1,
		Items:       nil,
		TotalWeight: 0,
		CreatedAt:   time.Time{},
		Events:      nil,
	}, nil)

	outbox := &mocks.Outbox{}
	outbox.On("Save", mock.Anything, mock.Anything).Return(nil)

	repoRegistry := &mocks.RepositoryRegistry{}
	repoRegistry.On("Basket").Return(repoBasket)
	repoRegistry.On("Outbox").Return(outbox)
	repoRegistry.On("Transaction", mock.Anything, mock.AnythingOfType(callBackFnTmpl)).
		Run(testCallback(repoRegistry)).
		Return(nil).Once()

	svc := NewBasketService(repoRegistry)
	err := svc.AddItemWithTx(context.Background(), 1, vos.BasketItem{
		BasketID: 1,
		GoodID:   1,
		Quantity: 1,
		Price:    1,
	})
	require.NoError(t, err)

	repoRegistry.AssertExpectations(t)
	repoBasket.AssertExpectations(t)
	outbox.AssertExpectations(t)
}

const callBackFnTmpl = "func(repository.RepositoryRegistry) error"

func testCallback(registry repository.RepositoryRegistry) func(args mock.Arguments) {
	return func(args mock.Arguments) {
		callback := args.Get(1).(func(repository.RepositoryRegistry) error)
		callback(registry)
	}
}

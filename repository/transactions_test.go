package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/vadiminshakov/dddgo/core/domain/aggregates"
	"github.com/vadiminshakov/dddgo/core/domain/vos"
	testing2 "github.com/vadiminshakov/dddgo/pkg/testing"
	"github.com/vadiminshakov/dddgo/repository/queries"

	"github.com/stretchr/testify/require"
)

func TestTxs(t *testing.T) {
	db, terminate, err := testing2.StartTestPostgres(context.Background())
	require.NoError(t, err)
	defer terminate(context.Background())

	_, err = db.Exec(queries.BasketCreate)
	require.NoError(t, err)

	_, err = db.Exec(queries.ItemsCreate)
	require.NoError(t, err)

	repo, err := New(db, nil)
	require.NoError(t, err)

	items := []*vos.BasketItem{
		{
			BasketID: 1,
			GoodID:   1,
			Quantity: 1,
			Price:    1,
			Weight:   1,
		},
	}

	basket := &aggregates.Basket{
		ID:    1,
		Items: items,
	}

	err = repo.Transaction(context.Background(), func(repo *RepoRegistry) error {
		if err := repo.Basket().Save(basket); err != nil {
			return err
		}
		if err := repo.Items().Save(items[0]); err != nil {
			return err
		}

		return nil
	})
	require.NoError(t, err)

	basketFromDb := &aggregates.Basket{}
	row := db.QueryRow("SELECT * FROM basket WHERE id = $1", basket.ID)
	if row.Err() != nil {
		t.Fatal(err)
	}

	err = row.Scan(&basketFromDb.ID, &basketFromDb.TotalWeight, &basketFromDb.CreatedAt)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, basket.ID, basketFromDb.ID)
	require.Equal(t, basket.TotalWeight, basketFromDb.TotalWeight)

	itemFromDb := &vos.BasketItem{}
	row = db.QueryRow("SELECT * FROM items WHERE basket_id = $1", basket.ID)
	if row.Err() != nil {
		t.Fatal(err)
	}

	err = row.Scan(&itemFromDb.BasketID, &itemFromDb.GoodID, &itemFromDb.Quantity, &itemFromDb.Price, &itemFromDb.Weight)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, items[0].BasketID, itemFromDb.BasketID)
	require.Equal(t, items[0].GoodID, itemFromDb.GoodID)
	require.Equal(t, items[0].Quantity, itemFromDb.Quantity)
	require.Equal(t, items[0].Price, itemFromDb.Price)
	require.Equal(t, items[0].Weight, itemFromDb.Weight)
}

func TestTxRollback(t *testing.T) {
	db, terminate, err := testing2.StartTestPostgres(context.Background())
	require.NoError(t, err)
	defer terminate(context.Background())

	_, err = db.Exec(queries.BasketCreate)
	require.NoError(t, err)

	_, err = db.Exec(queries.ItemsCreate)
	require.NoError(t, err)

	repo, err := New(db, nil)
	require.NoError(t, err)

	items := []*vos.BasketItem{
		{
			BasketID: 1,
			GoodID:   1,
			Quantity: 1,
			Price:    1,
			Weight:   1,
		},
	}

	basket := &aggregates.Basket{
		ID:    1,
		Items: items,
	}

	errRollback := errors.New("rollback")
	err = repo.Transaction(context.Background(), func(repo *RepoRegistry) error {
		if err := repo.Basket().Save(basket); err != nil {
			return err
		}
		if err := repo.Items().Save(items[0]); err != nil {
			return err
		}

		return errRollback
	})
	require.ErrorIs(t, err, errRollback)

	basketFromDb := &aggregates.Basket{}
	row := db.QueryRow("SELECT * FROM basket WHERE id = $1", basket.ID)
	if row.Err() != nil {
		t.Fatal(err)
	}

	err = row.Scan(&basketFromDb.ID, &basketFromDb.TotalWeight, &basketFromDb.CreatedAt)
	require.ErrorIs(t, err, sql.ErrNoRows)

	itemFromDb := &vos.BasketItem{}
	row = db.QueryRow("SELECT * FROM items WHERE basket_id = $1", basket.ID)
	if row.Err() != nil {
		t.Fatal(err)
	}

	err = row.Scan(&itemFromDb.BasketID, &itemFromDb.GoodID, &itemFromDb.Quantity, &itemFromDb.Price, &itemFromDb.Weight)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

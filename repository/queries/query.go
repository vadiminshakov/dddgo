package queries

const (
	BasketCreate = "CREATE TABLE IF NOT EXISTS basket (" +
		"id SERIAL PRIMARY KEY, total_weight INTEGER, created_at TIMESTAMP NOT NULL DEFAULT NOW())"
	BasketSave    = "INSERT INTO basket (id, total_weight) VALUES ($1, $2) ON CONFLICT (id) DO UPDATE SET total_weight = $2"
	BasketGetByID = "SELECT id, total_weight, created_at FROM basket WHERE id = $1"

	ItemsCreate = "CREATE TABLE IF NOT EXISTS items " +
		"(basket_id INTEGER REFERENCES basket (id) ON DELETE CASCADE, " +
		"good_id INTEGER, " +
		"quantity INTEGER, " +
		"price NUMERIC(10,2), " +
		"weight INTEGER, " +
		"PRIMARY KEY (basket_id, good_id))"
	ItemsSave = "INSERT INTO items (basket_id, good_id, quantity, price, weight) VALUES ($1, $2, $3, $4, $5) " +
		"ON CONFLICT (basket_id, good_id) DO UPDATE SET quantity = $3, price = $4, weight = $5"
	ItemsGetByBasketID = "SELECT basket_id, good_id, quantity, price, weight FROM items WHERE basket_id = $1"
)

package repository

import (
	"database/sql"
	"pretest-golang-tdi/config"
	"pretest-golang-tdi/model"
)

func AddItemToCart(userID, productID, quantity int) error {
	tx, err := config.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var cartID int
	err = tx.QueryRow("SELECT cart_id FROM carts WHERE user_id = $1", userID).Scan(&cartID)
	if err == sql.ErrNoRows {
		err = tx.QueryRow("INSERT INTO carts (user_id) VALUES ($1) RETURNING cart_id", userID).Scan(&cartID)
	}
	if err != nil {
		return err
	}

	query := `
		INSERT INTO cart_items (cart_id, product_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (cart_id, product_id)
		DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity;
	`
	_, err = tx.Exec(query, cartID, productID, quantity)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func GetUserCart(userID int) (model.Cart, error) {
	var cart model.Cart

	err := config.DB.QueryRow("SELECT cart_id, user_id, created_at FROM carts WHERE user_id = $1", userID).Scan(&cart.CartID, &cart.UserID, &cart.CreatedAt)
	if err != nil {
		return cart, err
	}

	query := `
		SELECT ci.cart_item_id, ci.quantity, p.product_id, p.name, p.description, p.price, p.stock
		FROM cart_items ci
		JOIN products p ON ci.product_id = p.product_id
		WHERE ci.cart_id = $1
	`
	rows, err := config.DB.Query(query, cart.CartID)
	if err != nil {
		return cart, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.CartItem
		if err := rows.Scan(&item.CartItemID, &item.Quantity, &item.Product.ProductID, &item.Product.Name, &item.Product.Description, &item.Product.Price, &item.Product.Stock); err != nil {
			return cart, err
		}
		item.CartID = cart.CartID
		item.ProductID = item.Product.ProductID
		cart.CartItems = append(cart.CartItems, item)
	}

	return cart, nil
}

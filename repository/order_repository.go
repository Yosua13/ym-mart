package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"pretest-golang-tdi/config"
	"pretest-golang-tdi/model"
	"time"

	"github.com/lib/pq"
)

func Checkout(userID int) (model.Order, error) {
	var order model.Order

	tx, err := config.DB.Begin()
	if err != nil {
		return order, err
	}
	defer tx.Rollback()

	// 1. Ambil keranjang dan item-itemnya
	cart, err := GetUserCart(userID)
	if err != nil {
		return order, errors.New("keranjang tidak ditemukan")
	}
	if len(cart.CartItems) == 0 {
		return order, errors.New("keranjang kosong")
	}

	// 2. Cek stok produk (dengan row locking untuk keamanan)
	totalAmount := 0.0
	for _, item := range cart.CartItems {
		var currentStock int
		err := tx.QueryRow("SELECT stock FROM products WHERE product_id = $1 FOR UPDATE", item.ProductID).Scan(&currentStock)
		if err != nil {
			return order, err
		}
		if currentStock < item.Quantity {
			return order, fmt.Errorf("stok untuk produk %s tidak mencukupi", item.Product.Name)
		}
		totalAmount += item.Product.Price * float64(item.Quantity)
	}

	// 3. Buat order baru
	order.UserID = userID
	order.InvoiceNumber = fmt.Sprintf("INV/%d/%d", userID, time.Now().Unix())
	order.Status = "Menunggu Pembayaran"
	order.TotalAmount = totalAmount

	orderQuery := `INSERT INTO orders (user_id, invoice_number, total_amount, status) VALUES ($1, $2, $3, $4) RETURNING order_id, created_at`
	err = tx.QueryRow(orderQuery, order.UserID, order.InvoiceNumber, order.TotalAmount, order.Status).Scan(&order.OrderID, &order.CreatedAt)
	if err != nil {
		return order, err
	}

	// 4. Masukkan item ke order_items dan kurangi stok
	for _, item := range cart.CartItems {
		itemQuery := `INSERT INTO order_items (order_id, product_id, product_name, price_at_purchase, quantity) VALUES ($1, $2, $3, $4, $5)`
		_, err = tx.Exec(itemQuery, order.OrderID, item.ProductID, item.Product.Name, item.Product.Price, item.Quantity)
		if err != nil {
			return order, err
		}

		updateStockQuery := `UPDATE products SET stock = stock - $1 WHERE product_id = $2`
		_, err = tx.Exec(updateStockQuery, item.Quantity, item.ProductID)
		if err != nil {
			return order, err
		}
	}

	// 5. Kosongkan keranjang
	_, err = tx.Exec("DELETE FROM cart_items WHERE cart_id = $1", cart.CartID)
	if err != nil {
		return order, err
	}

	return order, tx.Commit()
}

func GetUserOrders(userID int) ([]model.Order, error) {
	var orders []model.Order
	// 1. Ambil semua order utama milik user
	query := `SELECT order_id, user_id, invoice_number, total_amount, status, created_at FROM orders WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := config.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orderMap := make(map[int]*model.Order)
	var orderIDs []int

	for rows.Next() {
		var o model.Order
		if err := rows.Scan(&o.OrderID, &o.UserID, &o.InvoiceNumber, &o.TotalAmount, &o.Status, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
		orderIDs = append(orderIDs, o.OrderID)
		orderMap[o.OrderID] = &orders[len(orders)-1]
	}

	if len(orderIDs) == 0 {
		return orders, nil
	}

	// 2. Ambil semua item dari order-order tersebut dalam satu query
	itemQuery := `SELECT order_item_id, order_id, product_id, product_name, price_at_purchase, quantity FROM order_items WHERE order_id = ANY($1)`
	itemRows, err := config.DB.Query(itemQuery, pq.Array(orderIDs))
	if err != nil {
		return nil, err
	}
	defer itemRows.Close()

	// 3. Masukkan item ke dalam order yang sesuai
	for itemRows.Next() {
		var item model.OrderItem
		if err := itemRows.Scan(&item.OrderItemID, &item.OrderID, &item.ProductID, &item.ProductName, &item.PriceAtPurchase, &item.Quantity); err != nil {
			return nil, err
		}
		if order, ok := orderMap[item.OrderID]; ok {
			order.OrderItems = append(order.OrderItems, item)
		}
	}

	return orders, nil
}

func GetOrderByID(orderID int) (model.Order, error) {
	var order model.Order

	// Ambil data order utama
	err := config.DB.QueryRow("SELECT order_id, user_id, invoice_number, total_amount, status, created_at FROM orders WHERE order_id = $1", orderID).Scan(&order.OrderID, &order.UserID, &order.InvoiceNumber, &order.TotalAmount, &order.Status, &order.CreatedAt)
	if err != nil {
		return order, err
	}

	// Ambil item-item terkait
	itemQuery := `SELECT order_item_id, order_id, product_id, product_name, price_at_purchase, quantity FROM order_items WHERE order_id = $1`
	rows, err := config.DB.Query(itemQuery, orderID)
	if err != nil {
		return order, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.OrderItem
		if err := rows.Scan(&item.OrderItemID, &item.OrderID, &item.ProductID, &item.ProductName, &item.PriceAtPurchase, &item.Quantity); err != nil {
			return order, err
		}
		order.OrderItems = append(order.OrderItems, item)
	}

	return order, nil
}

func UpdateOrderStatus(orderID int, newStatus string) error {
	query := `UPDATE orders SET status = $1 WHERE order_id = $2`
	result, err := config.DB.Exec(query, newStatus, orderID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

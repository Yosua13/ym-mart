package repository

import (
	"pretest-golang-tdi/config"
	"pretest-golang-tdi/model"
	"strings"
)

func CreateProduct(product *model.Product) error {
	query := `INSERT INTO products (store_id, name, description, price, stock) VALUES ($1, $2, $3, $4, $5) RETURNING product_id, created_at`
	err := config.DB.QueryRow(query, product.StoreID, product.Name, product.Description, product.Price, product.Stock).Scan(&product.ProductID, &product.CreatedAt)
	return err
}

func GetProductByID(id int) (model.Product, error) {
	var product model.Product
	query := `SELECT product_id, store_id, name, description, price, stock, created_at FROM products WHERE product_id = $1`
	err := config.DB.QueryRow(query, id).Scan(&product.ProductID, &product.StoreID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.CreatedAt)
	return product, err
}

func GetAllProducts(search, sort string) ([]model.Product, error) {
	var products []model.Product

	baseQuery := "SELECT p.product_id, p.store_id, p.name, p.description, p.price, p.stock, p.created_at FROM products p"
	args := []interface{}{}
	whereClauses := []string{}

	if search != "" {
		whereClauses = append(whereClauses, "p.name ILIKE $1")
		args = append(args, "%"+search+"%")
	}

	if len(whereClauses) > 0 {
		baseQuery += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	switch sort {
	case "terbaru":
		baseQuery += " ORDER BY p.created_at DESC"
	case "terlama":
		baseQuery += " ORDER BY p.created_at ASC"
		// case "terlaris":
		// 	baseQuery = `
		// 		SELECT p.product_id, p.store_id, p.name, p.description, p.price, p.stock, p.created_at
		// 		FROM products p
		// 		LEFT JOIN (
		// 			SELECT product_id, SUM(quantity) as total_sold
		// 			FROM order_items
		// 			GROUP BY product_id
		// 		) sales ON p.product_id = sales.product_id
		// 	` + baseQuery[strings.Index(baseQuery, " WHERE"):] + ` ORDER BY COALESCE(sales.total_sold, 0) DESC`
	}

	rows, err := config.DB.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ProductID, &p.StoreID, &p.Name, &p.Description, &p.Price, &p.Stock, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

package repository

import (
	"pretest-golang-tdi/config"
	"pretest-golang-tdi/model"
	"strings"
)

func CreateStore(store *model.Store, userID int) error {
	query := `INSERT INTO stores (user_id, name, city) VALUES ($1, $2, $3) RETURNING store_id, created_at`
	err := config.DB.QueryRow(query, userID, store.Name, store.City).Scan(&store.StoreID, &store.CreatedAt)
	return err
}

func GetStoreByID(id int) (model.Store, error) {
	var store model.Store
	query := `SELECT store_id, name, city, created_at FROM stores WHERE store_id = $1`
	err := config.DB.QueryRow(query, id).Scan(&store.StoreID, &store.Name, &store.City, &store.CreatedAt)
	return store, err
}

func GetAllStores(search string) ([]model.Store, error) {
	var stores []model.Store

	baseQuery := "SELECT s.store_id, s.name, s.city, s.created_at FROM stores s"
	args := []interface{}{}
	whereClauses := []string{}

	if search != "" {
		whereClauses = append(whereClauses, "s.name ILIKE $1")
		args = append(args, "%"+search+"%")
	}

	if len(whereClauses) > 0 {
		baseQuery += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	rows, err := config.DB.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s model.Store
		if err := rows.Scan(&s.StoreID, &s.Name, &s.City, &s.CreatedAt); err != nil {
			return nil, err
		}
		stores = append(stores, s)
	}

	return stores, nil
}

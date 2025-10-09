package repository

import (
	"pretest-golang-tdi/config"
	"pretest-golang-tdi/model"
)

func CreateUser(user *model.User) error {
	query := `INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3) RETURNING user_id, created_at`
	err := config.DB.QueryRow(query, user.Username, user.PasswordHash, user.Role).Scan(&user.UserID, &user.CreatedAt)
	return err
}

func GetUserByUsername(username string) (model.User, error) {
	var user model.User
	query := `SELECT user_id, username, password_hash, role, created_at FROM users WHERE username = $1`
	err := config.DB.QueryRow(query, username).Scan(&user.UserID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt)
	return user, err
}

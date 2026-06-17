package repositories

import (
	"golangpkg/config"
	"golangpkg/models"
)

func CreateUser(user models.User) error {
	query := `INSERT INTO users(name, email, password) VALUES ($1, $2, $3)`
	_, err := config.DB.Exec(query, user.Name, user.Email, user.Password)
	return err
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email, password FROM users WHERE email=$1`
	err := config.DB.QueryRow(query, email).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
	)
	return user, err
}

func GetUserByID(id int) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email FROM users WHERE id=$1`
	err := config.DB.QueryRow(query, id).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
	)
	return user, err
}
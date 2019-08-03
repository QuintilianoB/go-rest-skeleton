package repository

import (
	"database/sql"
	"go-rest-skeleton/drivers"
	"go-rest-skeleton/models"
)

type UserRepository struct{}

func (u UserRepository) CreateUser(user *models.User) error {
	db := drivers.ConnectDB()
	defer db.Close()

	stmt := "select * from user where user = ?"
	rows, err := db.Query(stmt, user.User)

	if err != nil {
		return err
	}

	if rows.Next() {
		return models.UserAlreadyExist
	}

	stmt = "insert into user (user, password) value (? ,?)"
	_, err = db.Query(stmt, user.User, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) FindUser(user *models.User) error {
	db := drivers.ConnectDB()
	defer db.Close()

	stmt := "select * from user where user = ?"
	err := db.QueryRow(stmt, user.User).Scan(&user.ID, &user.User, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.UserNotFound
		} else {
			return err
		}
	}

	return nil
}

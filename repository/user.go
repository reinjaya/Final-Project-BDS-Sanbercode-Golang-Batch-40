package repository

import (
	"database/sql"
	"rein/final/structs"
)

func GetAllUser(db *sql.DB) (result []structs.User, err error) {
	sql := "SELECT * FROM users"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.UserName, &user.Password)
		if err != nil {
			panic(err)
		}

		result = append(result, user)
	}
	return
}

func GetUserById(db *sql.DB, userId structs.User) (result structs.User, err error) {
	sql := "SELECT * FROM users WHERE id = $1"

	rows, err := db.Query(sql, userId.ID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.UserName, &user.Password)
		if err != nil {
			panic(err)
		}

		result = user
	}
	return
}

func GetUserByUsername(db *sql.DB, userId structs.User) (result structs.User, err error) {
	sql := "SELECT * FROM users WHERE username = $1"

	rows, err := db.Query(sql, userId.UserName)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.UserName, &user.Password)
		if err != nil {
			panic(err)
		}

		result = user
	}
	return
}

func InsertUser(db *sql.DB, user structs.User) (err error) {
	sql := "INSERT INTO users (id, name, email, username, password) VALUES ($1,$2,$3,$4,$5)"

	errs := db.QueryRow(sql, user.ID, user.Name, user.Email, user.UserName, user.Password)
	return errs.Err()
}

func UpdateUser(db *sql.DB, user structs.User) (err error) {
	sql := "UPDATE users SET name = $1, email = $2, username = $3, password = $4 WHERE id = $5"

	errs := db.QueryRow(sql, user.Name, user.Email, user.UserName, user.Password, user.ID)
	return errs.Err()
}

func DeleteUser(db *sql.DB, user structs.User) (err error) {
	sql := "DELETE FROM users WHERE id = $1"

	errs := db.QueryRow(sql, user.ID)
	return errs.Err()
}

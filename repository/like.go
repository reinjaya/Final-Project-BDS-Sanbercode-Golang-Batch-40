package repository

import (
	"database/sql"
	"rein/final/structs"
)

func GetAllLike(db *sql.DB) (result []structs.Like, err error) {
	sql := "SELECT * FROM likes"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var like = structs.Like{}

		err = rows.Scan(&like.ID, &like.IDUser, &like.IDArticle, &like.Respon)
		if err != nil {
			panic(err)
		}

		result = append(result, like)
	}
	return
}

func GetLikeByIdUser(db *sql.DB, likeId structs.Like) (result structs.Like, err error) {
	sql := "SELECT * FROM likes WHERE id_user = $1"

	rows, err := db.Query(sql, likeId.IDUser)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var like = structs.Like{}

		err = rows.Scan(&like.ID, &like.IDUser, &like.IDArticle, &like.Respon)
		if err != nil {
			panic(err)
		}

		result = like
	}
	return
}

func GetLikeById(db *sql.DB, likeId structs.Like) (result structs.Like, err error) {
	sql := "SELECT * FROM likes WHERE id = $1"

	rows, err := db.Query(sql, likeId.ID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var like = structs.Like{}

		err = rows.Scan(&like.ID, &like.IDUser, &like.IDArticle, &like.Respon)
		if err != nil {
			panic(err)
		}

		result = like
	}
	return
}

func GetLikeByIdArticle(db *sql.DB, likeId structs.Like) (result structs.Like, err error) {
	sql := "SELECT * FROM likes WHERE id_article = $1"

	rows, err := db.Query(sql, likeId.IDArticle)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var like = structs.Like{}

		err = rows.Scan(&like.ID, &like.IDUser, &like.IDArticle, &like.Respon)
		if err != nil {
			panic(err)
		}

		result = like
	}
	return
}

func InsertLike(db *sql.DB, like structs.Like) (err error) {
	sql := "INSERT INTO likes (id, id_user, id_article, respon) VALUES ($1,$2,$3,$4)"

	errs := db.QueryRow(sql, like.ID, like.IDUser, like.IDArticle, like.Respon)
	return errs.Err()
}

func UpdateLike(db *sql.DB, like structs.Like) (err error) {
	sql := "UPDATE likes SET id_user = $1, id_article = $2, respon = $3 WHERE id = $4"

	errs := db.QueryRow(sql, like.IDUser, like.IDArticle, like.Respon, like.ID)
	return errs.Err()
}

func DeleteLike(db *sql.DB, like structs.Like) (err error) {
	sql := "DELETE FROM likes WHERE id = $1"

	errs := db.QueryRow(sql, like.ID)
	return errs.Err()
}

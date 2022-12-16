package repository

import (
	"database/sql"
	"rein/final/structs"
)

func GetAllComment(db *sql.DB) (result []structs.Comment, err error) {
	sql := "SELECT * FROM comments"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var comment = structs.Comment{}

		err = rows.Scan(&comment.ID, &comment.IDUser, &comment.IDArticle, &comment.IDReply, &comment.Comment, &comment.Image)
		if err != nil {
			panic(err)
		}

		result = append(result, comment)
	}
	return
}

func GetCommentById(db *sql.DB, commentId structs.Comment) (result structs.Comment, err error) {
	sql := "SELECT * FROM comments WHERE id = $1"

	rows, err := db.Query(sql, commentId.ID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var comment = structs.Comment{}

		err = rows.Scan(&comment.ID, &comment.IDUser, &comment.IDArticle, &comment.IDReply, &comment.Comment, &comment.Image)
		if err != nil {
			panic(err)
		}

		result = comment
	}
	return
}

func GetCommentByIdArticle(db *sql.DB, commentId structs.Comment) (result structs.Comment, err error) {
	sql := "SELECT * FROM comments WHERE id_article = $1"

	rows, err := db.Query(sql, commentId.IDArticle)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var comment = structs.Comment{}

		err = rows.Scan(&comment.ID, &comment.IDUser, &comment.IDArticle, &comment.IDReply, &comment.Comment, &comment.Image)
		if err != nil {
			panic(err)
		}

		result = comment
	}
	return
}

func InsertComment(db *sql.DB, comment structs.Comment) (err error) {
	sql := "INSERT INTO comments (id, id_user, id_article, id_reply_comment, comment, image) VALUES ($1,$2,$3,$4,$5,$6)"

	errs := db.QueryRow(sql, comment.ID, comment.IDUser, comment.IDArticle, comment.IDReply, comment.Comment, comment.Image)
	return errs.Err()
}

func UpdateComment(db *sql.DB, comment structs.Comment) (err error) {
	sql := "UPDATE comments SET id_user = $1, id_article = $2, id_reply_comment = $3, comment = $4, image = $5 WHERE id = $6"

	errs := db.QueryRow(sql, comment.IDUser, comment.IDArticle, comment.IDReply, comment.Comment, comment.Image, comment.ID)
	return errs.Err()
}

func DeleteComment(db *sql.DB, comment structs.Comment) (err error) {
	sql := "DELETE FROM comments WHERE id = $1"

	errs := db.QueryRow(sql, comment.ID)
	return errs.Err()
}

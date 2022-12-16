package repository

import (
	"database/sql"
	"rein/final/structs"
)

func GetAllArticle(db *sql.DB) (result []structs.Article, err error) {
	sql := "SELECT * FROM articles"

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var article = structs.Article{}

		err = rows.Scan(&article.ID, &article.IDUser, &article.Title, &article.Image, &article.Content)
		if err != nil {
			panic(err)
		}

		result = append(result, article)
	}
	return
}

func GetArticleById(db *sql.DB, articleId structs.Article) (result structs.Article, err error) {
	sql := "SELECT * FROM articles WHERE id = $1"

	rows, err := db.Query(sql, articleId.ID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var article = structs.Article{}

		err = rows.Scan(&article.ID, &article.IDUser, &article.Title, &article.Image, &article.Content)
		if err != nil {
			panic(err)
		}

		result = article
	}
	return
}

func InsertArticle(db *sql.DB, article structs.Article) (err error) {
	sql := "INSERT INTO articles (id, id_user, title, image, content) VALUES ($1,$2,$3,$4,$5)"

	errs := db.QueryRow(sql, article.ID, article.IDUser, article.Title, article.Image, article.Content)
	return errs.Err()
}

func UpdateArticle(db *sql.DB, article structs.Article) (err error) {
	sql := "UPDATE articles SET id_user = $1, title = $2, image = $3, content = $4 WHERE id = $5"

	errs := db.QueryRow(sql, article.IDUser, article.Title, article.Image, article.Content, article.ID)
	return errs.Err()
}

func DeleteArticle(db *sql.DB, article structs.Article) (err error) {
	sql := "DELETE FROM articles WHERE id = $1"

	errs := db.QueryRow(sql, article.ID)
	return errs.Err()
}

package repository

import (
	"database/sql"
	"quiz-3/structs"
	"time"
)

func GetAllBook(db *sql.DB) (err error, results []structs.Book) {
	sql := `SELECT * FROM books`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book structs.Book

		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.CategoryID,
		)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}
	
	return
}

func GetBooksByCategoryID(db *sql.DB, category_id int) (err error, results []structs.Book) {
	sql := `SELECT id, title, description, image_url, release_year, price, total_page, thickness, created_at, updated_at, category_id
		FROM books WHERE category_id = $1`
	
	rows, err := db.Query(sql, category_id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book structs.Book

		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CreatedAt,
			&book.UpdatedAt,
			&book.CategoryID,
		)
		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	return
}

func AddBook(db *sql.DB, book structs.Book) (err error) {
	sql := `INSERT INTO books (
		id,
		title,
		description,
		image_url,
		release_year,
		price,
		total_page,
		thickness,
		created_at,
		updated_at,
		category_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	
	created_time := time.Now().Local()
	updated_time := time.Now().Local()

	var thickness string
	if book.TotalPage <= 100 {
		thickness = "tipis"
	} else if book.TotalPage >= 101 && book.TotalPage <= 200 {
		thickness = "sedang"
	} else {
		thickness = "tebal"
	}

	errs := db.QueryRow(
		sql,
		book.ID,
		book.Title,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		thickness,
		created_time,
		updated_time,
		book.CategoryID,
	)
	
	return errs.Err()
}

func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	sql := `UPDATE books SET
		title = $1,
		description = $2,
		image_url = $3,
		release_year = $4,
		price = $5,
		total_page = $6,
		thickness = $7,
		updated_at = $8,
		category_id = $9 WHERE id = $10`

	updated_time := time.Now().Local()

	var thickness string
	if book.TotalPage <= 100 {
		thickness = "tipis"
	} else if book.TotalPage >= 101 && book.TotalPage <= 200 {
		thickness = "sedang"
	} else {
		thickness = "tebal"
	}

	errs := db.QueryRow(
		sql,
		book.Title,
		book.Description,
		book.ImageUrl,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		thickness,
		updated_time,
		book.CategoryID,
		book.ID,
	)
	
	return errs.Err()
}

func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	sql := `DELETE FROM books WHERE id = $1`

	errs := db.QueryRow(sql, book.ID)
	
	return errs.Err()
}
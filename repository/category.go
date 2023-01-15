package repository

import (
	"database/sql"
	"quiz-3/structs"
	"time"
)

func GetAllCategory(db *sql.DB) (err error, results []structs.Category) {
	sql := `SELECT * FROM categories`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var category structs.Category

		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, category)
	}
	
	return
}

func AddCategory(db *sql.DB, category structs.Category) (err error) {
	sql := `INSERT INTO categories (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)`
	
	created_time := time.Now().Local()
	updated_time := time.Now().Local()

	errs := db.QueryRow(sql, category.ID, category.Name, created_time, updated_time)
	
	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := `UPDATE categories SET name = $1, updated_at = $2 WHERE id = $3`
	updated_time := time.Now().Local()

	errs := db.QueryRow(sql, category.Name, updated_time, category.ID)
	
	return errs.Err()
}

func DeleteCategory(db *sql.DB, person structs.Category) (err error) {
	sql := `DELETE FROM categories WHERE id = $1`

	errs := db.QueryRow(sql, person.ID)
	
	return errs.Err()
}
package services

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/models"
	"time"
)

func GetAllProject(db *sql.DB) (err error, projects []models.Project) {
	sql := `SELECT * FROM projects`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var project models.Project

		err = rows.Scan(
			&project.ID, 
			&project.Name, 
			&project.Description, 
			&project.Status,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}

		projects = append(projects, project)
	}
	
	return
}

func CreateProject(db *sql.DB, project models.Project) (err error) {
	sql := `INSERT INTO projects (
		id,
		name,
		description,
		status,
		created_at,
		updated_at
		) VALUES ($1, $2, $3, $4, $5, $6)`

	created_time := time.Now().Local()
	updated_time := time.Now().Local()

	err = db.QueryRow(
		sql,
		project.ID,
		project.Name,
		project.Description,
		project.Status,
		created_time,
		updated_time,
	).Err()
	if err != nil {
		panic(err)
	}
	
	return
}

func UpdateProject(db *sql.DB, project models.Project) (err error) {
	sql := `UPDATE projects SET
		name = $1,
		description = $2,
		status = $3,
		updated_at = $4
		WHERE id = $5`

	updated_time := time.Now().Local()

	err = db.QueryRow(
		sql,
		project.Name,
		project.Description,
		project.Status,
		updated_time,
		project.ID,
	).Err()
	if err != nil {
		panic(err)
	}

	return
}

func DeleteProject(db *sql.DB, project models.Project) (err error) {
	sql := `DELETE FROM projects WHERE id = $1`

	err = db.QueryRow(sql, project.ID).Err()
	if err != nil {
		panic(err)
	}

	return
}
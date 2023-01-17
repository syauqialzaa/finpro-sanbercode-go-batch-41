package services

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/models"
	"time"
)

func GetAllTask(db *sql.DB) (err error, tasks []models.Task) {
	sql := `SELECT * FROM tasks`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task

		err = rows.Scan(
			&task.ID,
			&task.ProjectID,
			&task.TeamID,
			&task.Name,
			&task.Description,
			&task.AssignedMemberID,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}

		tasks = append(tasks, task)
	}
	
	return
}

func GetTaskByProjectID(db *sql.DB, projectID int) (err error, tasks []models.Task) {
	sql := `SELECT
		id,
		project_id,
		team_id,
		name,
		description,
		assigned_member_id,
		status,
		created_at,
		updated_at
		FROM tasks WHERE project_id = $1`

	rows, err := db.Query(sql, projectID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task

		err = rows.Scan(
			&task.ID,
			&task.ProjectID,
			&task.TeamID,
			&task.Name,
			&task.Description,
			&task.AssignedMemberID,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}

		tasks = append(tasks, task)
	}

	return
}

func CreateTask(db *sql.DB, task models.Task) (err error) {
	sql := `INSERT INTO tasks (
		id,
		project_id,
		team_id,
		name,
		description,
		assigned_member_id,
		status,
		created_at,
		updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	created_time := time.Now().Local()
	updated_time := time.Now().Local()

	err = db.QueryRow(
		sql,
		task.ID,
		task.ProjectID,
		task.TeamID,
		task.Name,
		task.Description,
		task.AssignedMemberID,
		task.Status,
		created_time,
		updated_time,
	).Err()
	if err != nil {
		panic(err)
	}
	
	return
}

func UpdateTask(db *sql.DB, task models.Task) (err error) {
	sql := `UPDATE tasks SET
		project_id = $1,
		team_id = $2,
		name = $3,
		description = $4,
		assigned_member_id = $5,
		status = $6,
		updated_at = $7
		WHERE id = $8`

	updated_time := time.Now().Local()

	err = db.QueryRow(
		sql,
		task.ProjectID,
		task.TeamID,
		task.Name,
		task.Description,
		task.AssignedMemberID,
		task.Status,
		updated_time,
		task.ID,
	).Err()
	if err != nil {
		panic(err)
	}

	return
}

func DeleteTask(db *sql.DB, task models.Task) (err error) {
	sql := `DELETE FROM tasks WHERE id = $1`

	err = db.QueryRow(sql, task.ID).Err()
	if err != nil {
		panic(err)
	}

	return
}
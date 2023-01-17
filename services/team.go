package services

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/models"
)

func GetAllTeam(db *sql.DB) (err error, teams []models.Team) {
	sql := `SELECT * FROM teams`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var team models.Team

		err = rows.Scan(&team.ID, &team.Name)
		if err != nil {
			panic(err)
		}

		teams = append(teams, team)
	}
	
	return
}

func CreateTeam(db *sql.DB, team models.Team) (err error) {
	sql := `INSERT INTO teams (id, name) VALUES ($1, $2)`

	err = db.QueryRow(sql, team.ID, team.Name).Err()
	if err != nil {
		panic(err)
	}
	
	return
}

func UpdateTeam(db *sql.DB, team models.Team) (err error) {
	sql := `UPDATE teams SET name = $1 WHERE id = $2`

	err = db.QueryRow(sql, team.Name, team.ID).Err()
	if err != nil {
		panic(err)
	}

	return
}

func DeleteTeam(db *sql.DB, team models.Team) (err error) {
	sql := `DELETE FROM teams WHERE id = $1`

	err = db.QueryRow(sql, team.ID).Err()
	if err != nil {
		panic(err)
	}

	return
}
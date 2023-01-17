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
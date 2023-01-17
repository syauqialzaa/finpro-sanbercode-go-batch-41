package services

import (
	"database/sql"
	"final-project-sanbercode-go-batch-41/models"
)

func GetAllMember(db *sql.DB) (err error, members []models.Member) {
	sql := `SELECT * FROM members`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var member models.Member

		err = rows.Scan(&member.ID, &member.Name, &member.Email, &member.TeamID)
		if err != nil {
			panic(err)
		}

		members = append(members, member)
	}
	
	return
}

func GetMemberByTeamID(db *sql.DB, teamID int) (err error, members []models.Member) {
	sql := `SELECT id, name, email, team_id FROM members WHERE team_id = $1`

	rows, err := db.Query(sql, teamID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var member models.Member

		err = rows.Scan(&member.ID, &member.Name, &member.Email, &member.TeamID)
		if err != nil {
			panic(err)
		}

		members = append(members, member)
	}

	return
}

func AddMember(db *sql.DB, member models.Member) (err error) {
	sql := `INSERT INTO members (id, name, email, team_id) VALUES ($1, $2, $3, $4)`

	err = db.QueryRow(sql, member.ID, member.Name, member.Email, member.TeamID).Err()
	if err != nil {
		panic(err)
	}
	
	return
}

func UpdateMember(db *sql.DB, member models.Member) (err error) {
	sql := `UPDATE members SET name = $1, email = $2, team_id = $3 WHERE id = $4`

	err = db.QueryRow(sql, member.Name, member.Email, member.TeamID, member.ID).Err()
	if err != nil {
		panic(err)
	}

	return
}

func DeleteMember(db *sql.DB, member models.Member) (err error) {
	sql := `DELETE FROM members WHERE id = $1`

	err = db.QueryRow(sql, member.ID).Err()
	if err != nil {
		panic(err)
	}

	return
}
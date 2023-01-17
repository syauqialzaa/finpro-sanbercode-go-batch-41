package models

type Member struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	TeamID   int       `json:"team_id"`
}
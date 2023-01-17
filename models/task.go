package models

import "time"

type Task struct {
	ID                 int        `json:"id"`
	ProjectID          int        `json:"project_id"`
	TeamID             int        `json:"team_id"`
	Name               string     `json:"name"`
	Description        string     `json:"description"`
	AssignedMemberID   int        `json:"assigned_member_id"`
	Status             string     `json:"status"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time	`json:"updated_at"`
}
package models

import "time"

type Project struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time	`json:"updated_at"`
}
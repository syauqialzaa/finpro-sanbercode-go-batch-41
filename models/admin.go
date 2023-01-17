package models

import "time"

type Admin struct {
	ID           int          `json:"id"`
	Username     string       `json:"username"`
	Password     string       `json:"password"`
	Token        string       `json:"token"`
	TokenExp     time.Time	  `json:"token_exp"`
}
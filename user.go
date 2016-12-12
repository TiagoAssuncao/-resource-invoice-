package main

import "time"

type User struct {
	Id               int       `json:"id"`
	FirstName        string    `json:"firstname"`
	LastName         string    `json:"lastname"`
	Email            string    `json:"email"`
	IsActive         bool      `json:"isactive"`
	CreateAt         time.Time `json:"create"`
	DeactiveAt       time.Time `json:"deactive"`
}

type Users []User

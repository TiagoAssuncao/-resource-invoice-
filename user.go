package main

import "time"

type User struct {
	Id               int       `json:"id"`
	FirstName        string    `json:"name"`
	LastName         string    `json:"name"`
	Email            string    `json:"name"`
	IsActive         bool      `json:"completed"`
	CreateAt         time.Time `json:"due"`
	DeactiveAt       time.Time `json:"due"`
}

type Users []User

package main

import "time"

type Invoice struct {
	Id        int       `json:"id"`
	ReferenceMonth int       `json:"referencemonth"`
	ReferenceYear int       `json:"referenceyear"`
	Amount int       `json:"amount"`
	Document string    `json:"document"`
	Description string    `json:"description"`
	IsActive bool      `json:"isactive"`
	CreateAt time.Time `json:"createat"`
	DeactiveAt time.Time `json:"deactiveat"`
}

type Invoices []Invoice

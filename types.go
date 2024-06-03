package main

import (
	"math/rand"
	"time"
)

type CreateAccReq struct {
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
}

type Account struct {
	ID         int       `json:"id"`
	FirstName  string    `json:"firstName"`
	SecondName string    `json:"secondName"`
	Number     int       `json:"number"`
	Balance    int       `json:"balance"`
	CreatedAt  time.Time `json:"createdAt"`
}

func newAccount(firstName, secondName string) *Account {
	return &Account{
		FirstName:  firstName,
		SecondName: secondName,
		Number:     rand.Intn(100),
		Balance:    rand.Intn(100000),
		CreatedAt:  time.Now().UTC(),
	}
}

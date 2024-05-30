package main

import "math/rand"

type Account struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Number     int    `json:"number"`
	Balance    int    `json:"balance"`
}

func newAccount(firstName, secondName string) *Account {
	return &Account{
		ID:         rand.Intn(100),
		FirstName:  firstName,
		SecondName: secondName,
		Number:     rand.Intn(100),
		Balance:    rand.Intn(100000),
	}
}

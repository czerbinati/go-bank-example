package main

import "math/rand"

// Account Type
type Account struct {
	ID        int32  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

// Constructor
func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        int32(rand.Intn(10000)),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Int63n(1000000),
	}
}

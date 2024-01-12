package bank

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// Customer ...
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// Account ...
type Account struct {
	Customer
	Number  int32
	Balance float64
}

func Hello() string {
	return "Hey! I'm working!"
}

// Deposit ...
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

// Withdraw ...
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw should be greater than the account's balance")
	}

	a.Balance -= amount
	return nil
}

type Statementer interface {
	Statement() string
}

type CustomAccount struct {
	*Account
}

// Statement ...
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

func (a *CustomAccount) Statement() string {
	b, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func Statement(s Statementer) string {
	return s.Statement()
}

func (a *Account) Transform(to *Account, amount float64) {
	a.Withdraw(amount)
	to.Deposit(amount)
}

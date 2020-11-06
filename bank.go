package bank

import (
	"errors"
	"strconv"
)

type Account struct {
	Amount    int
	Statement string
}

// type AccountErr string

var (
	NotEnoughMoneyInAccount = errors.New("You do not have enough money in your account.")
)

func (a *Account) Balance() int {
	return a.Amount
}

func (a *Account) Add(deposit int) {
	a.Amount += deposit

	a.Statement += "You added " + strconv.Itoa(deposit) + ". "
	a.Statement += "Balance: " + strconv.Itoa(a.Amount) + ". "
}

func (a *Account) Withdraw(value int) error {
	if value > a.Amount {
		return NotEnoughMoneyInAccount
	} else {
		a.Amount -= value

		a.Statement += "You withdrew " + strconv.Itoa(value) + ". "
		a.Statement += "Balance: " + strconv.Itoa(a.Amount) + ". "
		return nil
	}

}

func (a *Account) BankStatement() string {
	return a.Statement
}

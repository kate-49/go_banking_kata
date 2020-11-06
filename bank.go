package bank

import (
	"errors"
	"strconv"
)

type Account struct {
	Amount int
}

var m = make(map[string]int)

// type AccountErr string

var (
	NotEnoughMoneyInAccount = errors.New("You do not have enough money in your account.")
)

func (a *Account) Balance() int {
	return a.Amount
}

func (a *Account) Add(deposit int) {
	a.Amount += deposit
	total := a.Amount
	comment := ""
	comment += "You added " + strconv.Itoa(deposit) + ". "
	comment += "Balance: " + strconv.Itoa(a.Amount) + ". \n"

	m[comment] = total
}

// func (a *Account) Withdraw(value int) error {
// 	if value > a.Amount {
// 		return NotEnoughMoneyInAccount
// 	} else {
// 		a.Amount -= value

// 		a.Statement += "You withdrew " + strconv.Itoa(value) + ". "
// 		a.Statement += "Balance: " + strconv.Itoa(a.Amount) + ". \n"
// 		return nil
// 	}

// }

func (a *Account) BankStatement() string {
	returnStatement := ""

	for comment, _ := range m {
		returnStatement += comment
	}

	return returnStatement
}

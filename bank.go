package bank

import (
	"errors"
	"strconv"
)

type Account struct {
	Amount int
}

var m = make(map[string]int)

var (
	NotEnoughMoneyInAccount = errors.New("You do not have enough money in your account.")
)

func (a *Account) Balance() int {
	return a.Amount
}

func (a *Account) Add(deposit int) {
	a.Amount += deposit
	total := a.Amount
	addComment := ""
	addComment += "You added " + strconv.Itoa(deposit) + ". "
	addComment += "Balance: " + strconv.Itoa(a.Amount) + ". \n"

	m[addComment] = total
}

func (a *Account) Withdraw(value int) error {
	if value > a.Amount {
		return NotEnoughMoneyInAccount
	} else {
		a.Amount -= value
		total := a.Amount

		withdrawComment := ""
		withdrawComment += "You withdrew " + strconv.Itoa(value) + ". "
		withdrawComment += "Balance: " + strconv.Itoa(a.Amount) + ". \n"

		m[withdrawComment] = total
		return nil
	}

}

func (a *Account) BankStatement() string {
	returnStatement := ""

	for comment, _ := range m {
		returnStatement += comment
	}

	return returnStatement
}

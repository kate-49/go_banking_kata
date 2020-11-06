package bank

import (
	"errors"
	"strconv"
	"time"
)

type Account struct {
	Amount    int
	Statement *UserStatement
}

type UserStatement struct {
	Update string
	Date   time.Time
}

var (
	NotEnoughMoneyInAccount = errors.New("You do not have enough money in your account.")
)

func (a *Account) Balance() int {
	return a.Amount
}

func (a *Account) Add(deposit int) {
	a.Amount += deposit

	addComment := ""
	addComment += "You added " + strconv.Itoa(deposit) + ". "
	addComment += "Balance: " + strconv.Itoa(a.Amount) + ". \n"

	a.Statement[addComment] = time.Now()
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

		a.Statement[withdrawComment] = total
		return nil
	}

}

func (a *Account) BankStatement() string {
	returnStatement := ""

	for comment, _ := range Account.m {
		returnStatement += comment
	}

	return returnStatement
}

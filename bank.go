package bank

import "strconv"

type Account struct {
	Amount    int
	Statement string
}

func (a *Account) Balance() int {
	return a.Amount
}

func (a *Account) Add(deposit int) {
	a.Amount += deposit
	newStatement := "You added "
	newStatement += strconv.Itoa(deposit)
	a.Statement += newStatement
}

func (a *Account) BankStatement() string {
	return a.Statement
}

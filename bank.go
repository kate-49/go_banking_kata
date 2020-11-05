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

	a.Statement += "You added " + strconv.Itoa(deposit) + ". "
	a.Statement += "Balance: " + strconv.Itoa(a.Amount) + ". "
}

func (a *Account) Withdraw(value int) {
	a.Amount -= value
}

func (a *Account) BankStatement() string {
	return a.Statement
}

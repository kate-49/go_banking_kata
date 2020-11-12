package bank

import "errors"

type Account struct {
	Amount    int
	Statement []string
}

var (
	NotEnoughMoneyInAccount = errors.New("You do not have enough money in your account.")
)

func (a *Account) Balance() int {
	return a.Amount
}

func (a *Account) BankStatement() []string {
	return a.Statement
}

func (a *Account) Add(deposit int) {
	a.Amount += deposit
	a.Statement = append(a.Statement, "test statement")
}

func (a *Account) Withdraw(value int) error {
	if value > a.Amount {
		return NotEnoughMoneyInAccount
	} else {
		a.Amount -= value
		a.Statement = append(a.Statement, "test statement")
		return nil
	}

}

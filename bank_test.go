package bank

import "testing"

func TestCanCheckBalance(t *testing.T) {
	account := Account{10}
	got := account.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCanAddMoney(t *testing.T) {
	account := Account{100}
	account.Add(10)
	got := account.Balance()
	want := 110

	if got != want {
		t.Errorf("Got %d want %d", got, want)
	}
}

// func TestBankStatementPrintsInteractions(t *testing.T) {
// 	account := Account{100}
// 	got := account.BankStatement
// 	want := ""
// }

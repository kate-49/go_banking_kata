package bank

import "testing"

func TestCanCheckBalance(t *testing.T) {
	account := Account{10, ""}
	got := account.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCanAddMoney(t *testing.T) {
	account := Account{100, ""}
	account.Add(10)
	got := account.Balance()
	want := 110

	if got != want {
		t.Errorf("Got %d want %d", got, want)
	}
}

func TestBankStatementPrints1Interactions(t *testing.T) {
	account := Account{50, ""}
	account.Add(10)
	got := account.BankStatement()
	want := "You added 10. Balance is now 60. "

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func TestBankStatementPrintsMultipleInteractions(t *testing.T) {
	account := Account{20, ""}
	account.Add(1)
	account.Add(5)
	got := account.BankStatement()
	want := "You added 1. Balance is now 21. You added 5. Balance is now 26. "

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

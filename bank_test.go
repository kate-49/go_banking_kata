package bank

import (
	"testing"
)

func assertValues(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestCanCheckBalance(t *testing.T) {
	account := Account{10, ""}
	got := account.Balance()
	want := 10

	assertValues(t, got, want)
}

func TestCanAddMoney(t *testing.T) {
	account := Account{100, ""}

	t.Run("one addition", func(t *testing.T) {
		account.Add(10)
		got := account.Balance()
		want := 110

		assertValues(t, got, want)
	})

	t.Run("multiple additions", func(t *testing.T) {
		account.Add(20)
		account.Add(12)
		account.Add(8)

		got := account.Balance()
		want := 150

		assertValues(t, got, want)
	})
}

func TestCanWithdrawMoney(t *testing.T) {
	account := Account{90, ""}

	t.Run("single withdrawl", func(t *testing.T) {
		account.Withdraw(20)
		got := account.Balance()
		want := 70

		assertValues(t, got, want)
	})

	t.Run("multiple withdawl", func(t *testing.T) {
		account.Withdraw(10)
		account.Withdraw(1)
		account.Withdraw(2)

		got := account.Balance()
		want := 57

		assertValues(t, got, want)
	})
}

func TestBankStatementPrints1Interactions(t *testing.T) {
	t.Run("Single statement", func(t *testing.T) {
		account := Account{50, ""}
		account.Add(10)
		got := account.BankStatement()
		want := "You added 10. Balance: 60. "

		assertStrings(t, got, want)
	})

	t.Run("multiple statement", func(t *testing.T) {
		account := Account{20, ""}
		account.Add(1)
		account.Add(5)
		got := account.BankStatement()
		want := "You added 1. Balance: 21. You added 5. Balance: 26. "

		assertStrings(t, got, want)
	})
}

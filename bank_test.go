package bank

import (
	"reflect"
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

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error")
	}
}

func TestCanCheckBalance(t *testing.T) {
	account1 := Account{10, []string{""}}
	got := account1.Balance()
	want := 10

	assertValues(t, got, want)
}

func TestCanAddMoney(t *testing.T) {

	t.Run("multiple additions", func(t *testing.T) {
		account2 := Account{100, []string{""}}
		account2.Add(20)
		account2.Add(12)
		account2.Add(8)

		got := account2.Balance()
		want := 140

		assertValues(t, got, want)
	})
}

func TestCanWithdrawMoney(t *testing.T) {

	t.Run("multiple withdawl", func(t *testing.T) {
		account3 := Account{650, []string{" "}}
		account3.Withdraw(10)
		account3.Withdraw(1)
		account3.Withdraw(2)

		got := account3.Balance()
		want := 637

		assertValues(t, got, want)
	})

	t.Run("can't withdraw money if not enough in account", func(t *testing.T) {
		account4 := Account{20, []string{""}}
		err := account4.Withdraw(40)

		assertError(t, err, NotEnoughMoneyInAccount)
	})
}

func TestBankStatementPrints1Interactions(t *testing.T) {
	t.Run("Single statement", func(t *testing.T) {
		account6 := Account{50, []string{""}}
		account6.Add(10)
		got := account6.BankStatement()
		want := []string{"", "You added 10. Balance: 60."}

		if (reflect.DeepEqual(got, want)) != true {
			t.Errorf("Got %v, want %v", got, want)
		}

	})

	t.Run("multiple statement", func(t *testing.T) {
		account := Account{20, []string{""}}
		account.Add(1)
		account.Withdraw(2)
		account.Add(5)
		account.Withdraw(1)
		got := account.BankStatement()
		want := []string{
			"",
			"You added 1. Balance: 21.",
			"You withdrew 2. Balance: 19.",
			"You added 5. Balance: 24.",
			"You withdrew 1. Balance: 23.",
		}

		if (reflect.DeepEqual(got, want)) != true {
			t.Errorf("Got %v, want %v", got, want)
		}
	})
}

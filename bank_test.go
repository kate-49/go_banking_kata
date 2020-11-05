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

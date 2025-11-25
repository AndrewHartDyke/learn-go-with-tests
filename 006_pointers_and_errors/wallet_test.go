package wallet

import "testing"

func TestWallet(t *testing.T) {

	t.Run("bitcoin stringer", func(t *testing.T) {
		coin := Bitcoin(20)
		got := coin.String()
		want := "20 BTC"
		assertEqual(t, got, want)
	})
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{20}
		err := wallet.Withdraw(Bitcoin(5))
		assertNoError(t, err)
		assertBalance(t, wallet, 15)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(5)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(10))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(startingBalance))
	})
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q want %q", got, want)
	}
}

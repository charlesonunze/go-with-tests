package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin struct
type Bitcoin int

// Wallet struct
type Wallet struct {
	balance Bitcoin
}

// Deposit function
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("Cannot withdraw, insufficient funds")

// Withdraw function
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance function
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

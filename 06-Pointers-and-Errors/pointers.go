package pointers

import "fmt"

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

// Withdraw function
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

// Balance function
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

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
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

// Balance function
func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	return w.balance
}

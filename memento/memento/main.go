package main

import "fmt"

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	return &Memento{Balance: b.balance}
}

func (b *BankAccount) Restore(snapshot *Memento) {
	b.balance = snapshot.Balance
}

func NewBankAccount(balance int) (*BankAccount, *Memento) {
	return &BankAccount{balance: balance}, &Memento{Balance: balance}
}

func main() {
	account, snap0 := NewBankAccount(100)
	snap1 := account.Deposit(50)
	snap2 := account.Deposit(25)
	fmt.Println(account)

	account.Restore(snap1)
	fmt.Println(account)

	account.Restore(snap2)
	fmt.Println(account)

	account.Restore(snap0)
	fmt.Println(account)
}

package main

import "fmt"

type Memento struct {
	Balance int
}

type BankAccount struct {
	balance int
	changes []*Memento
	current int
}

func (b *BankAccount) String() string {
	return fmt.Sprintf("Balance = $%d, current = %d", b.balance, b.current)
}

func (b *BankAccount) Deposit(amount int) *Memento {
	b.balance += amount
	snap := &Memento{Balance: b.balance}
	b.changes = append(b.changes, snap)
	b.current++
	fmt.Printf("Deposited %d, balance is now %d\n", amount, b.balance)
	return snap
}

func (b *BankAccount) Restore(snap *Memento) {
	if snap != nil {
		b.balance = snap.Balance
		b.changes = append(b.changes, snap)
		b.current = len(b.changes) - 1
	}
}

func (b *BankAccount) Undo() *Memento {
	if b.current > 0 {
		b.current--
		snap := b.changes[b.current]
		b.balance = snap.Balance
		return snap
	}
	return nil
}

func (b *BankAccount) Redo() *Memento {
	if b.current+1 < len(b.changes) {
		b.current++
		snap := b.changes[b.current]
		b.balance = snap.Balance
		return snap
	}
	return nil
}

func NewBankAccount(balance int) *BankAccount {
	account := &BankAccount{balance: balance}
	account.changes = append(account.changes, &Memento{Balance: balance})
	return account
}

func main() {
	account := NewBankAccount(100)
	account.Deposit(50)
	account.Deposit(25)
	fmt.Println(account)

	account.Undo()
	fmt.Println("Undo 1:", account)

	account.Undo()
	fmt.Println("Undo 2:", account)

	account.Redo()
	fmt.Println("Redo:", account)
}

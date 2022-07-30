package main

import "fmt"

const (
	Deposit Action = iota
	Withdraw
)

var overdraftLimit = -500

type Action int

type Command interface {
	Call()
	Undo()
}

type BankAccountCommand struct {
	account   *BankAccount
	action    Action
	amount    int
	succeeded bool
}

func (b *BankAccountCommand) Call() {
	switch b.action {
	case Deposit:
		b.account.Deposit(b.amount)
		b.succeeded = true
	case Withdraw:
		b.succeeded = b.account.Withdraw(b.amount)
	}
}

func (b *BankAccountCommand) Undo() {
	if !b.succeeded {
		return
	}
	switch b.action {
	case Deposit:
		b.account.Withdraw(b.amount)
	case Withdraw:
		b.account.Deposit(b.amount)
	}
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{
		account: account,
		action:  action,
		amount:  amount,
	}
}

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount, "\b, balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount, "\b, balance is now", b.balance)
		return true
	}
	return false
}

func main() {
	account := BankAccount{}
	depositCommand := NewBankAccountCommand(&account, Deposit, 100)
	depositCommand.Call()
	fmt.Println(account)

	withdrawCommand := NewBankAccountCommand(&account, Withdraw, 25)
	withdrawCommand.Call()
	fmt.Println(account)

	withdrawCommand.Undo()
	fmt.Println(account)
}

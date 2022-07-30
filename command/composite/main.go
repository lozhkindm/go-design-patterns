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
	GetSucceeded() bool
	SetSucceeded(v bool)
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

func (b *BankAccountCommand) GetSucceeded() bool {
	return b.succeeded
}

func (b *BankAccountCommand) SetSucceeded(v bool) {
	b.succeeded = v
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{
		account: account,
		action:  action,
		amount:  amount,
	}
}

type CompositeBankAccountCommand struct {
	commands []Command
}

func (c *CompositeBankAccountCommand) Call() {
	for _, cmd := range c.commands {
		cmd.Call()
	}
}

func (c *CompositeBankAccountCommand) Undo() {
	for i := range c.commands {
		c.commands[len(c.commands)-i-1].Undo()
	}
}

func (c *CompositeBankAccountCommand) GetSucceeded() bool {
	for _, cmd := range c.commands {
		if !cmd.GetSucceeded() {
			return false
		}
	}
	return true
}

func (c *CompositeBankAccountCommand) SetSucceeded(v bool) {
	for _, cmd := range c.commands {
		cmd.SetSucceeded(v)
	}
}

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount   int
}

func (m *MoneyTransferCommand) Call() {
	ok := true
	for _, cmd := range m.commands {
		if ok {
			cmd.Call()
			ok = cmd.GetSucceeded()
		} else {
			cmd.SetSucceeded(false)
		}
	}
}

func NewMoneyTransferCommand(from *BankAccount, to *BankAccount, amount int) *MoneyTransferCommand {
	cmd := &MoneyTransferCommand{
		from:   from,
		to:     to,
		amount: amount,
	}
	cmd.commands = append(cmd.commands, NewBankAccountCommand(from, Withdraw, amount))
	cmd.commands = append(cmd.commands, NewBankAccountCommand(to, Deposit, amount))
	return cmd
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
	sender := &BankAccount{balance: 100}
	receiver := &BankAccount{balance: 0}

	transfer := NewMoneyTransferCommand(sender, receiver, 25)
	transfer.Call()
	fmt.Println(sender, receiver)

	transfer.Undo()
	fmt.Println(sender, receiver)
}

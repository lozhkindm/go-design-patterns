package main

import "fmt"

type BankAccount struct {
	Balance int
}

func Deposit(account *BankAccount, amount int) {
	fmt.Println("Depositing", amount)
	account.Balance += amount
}

func Withdraw(account *BankAccount, amount int) {
	if account.Balance >= amount {
		fmt.Println("Withdrawing", amount)
		account.Balance -= amount
	}
}

func main() {
	account := &BankAccount{Balance: 0}
	var commands []func()
	commands = append(commands, func() {
		Deposit(account, 100)
	})
	commands = append(commands, func() {
		Withdraw(account, 25)
	})

	for _, cmd := range commands {
		cmd()
	}
	fmt.Println(account)
}

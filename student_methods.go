package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
}

func (a *BankAccount) Withdraw(amount float64) bool {
	if a.Balance >= amount {
		a.Balance -= amount
		return true
	}
	return false
}

func (a BankAccount) GetBalance() float64 {
	return a.Balance
}

func main() {
	account := BankAccount{Owner: "Ирина"}

	account.Deposit(1000)
	fmt.Printf("Баланс после пополнения: %.2f\n", account.GetBalance())

	if account.Withdraw(300) {
		fmt.Println("Снятие успешно")
	} else {
		fmt.Println("Недостаточно средств")
	}

	fmt.Printf("Текущий баланс: %.2f\n", account.GetBalance())
}

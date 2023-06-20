package main

import (
	"fmt"
)

type Customer struct {
	accountNumber int
	name          string
	balance       float64
}

type Bank struct {
	customers []Customer
}

func (bank *Bank) addCustomer(accountNumber int, name string, balance float64) {
	newCustomer := Customer{
		accountNumber: accountNumber,
		name:          name,
		balance:       balance,
	}
	bank.customers = append(bank.customers, newCustomer)
}

func (bank *Bank) totalCustomers() int {
	return len(bank.customers)
}

func (bank *Bank) printAllCustomers() {
	for _, customer := range bank.customers {
		fmt.Printf("Account Number: %d, Name: %s, Balance: %.2f\n", customer.accountNumber, customer.name, customer.balance)
	}
}

func (customer *Customer) withdraw(amount float64) {
	if customer.balance >= amount {
		customer.balance -= amount
		fmt.Printf("Withdrawal successful. New balance: %.2f\n", customer.balance)
	} else {
		fmt.Println("Insufficient funds.")
	}
}

func (customer *Customer) deposit(amount float64) {
	customer.balance += amount
	fmt.Printf("Deposit successful. New balance: %.2f\n", customer.balance)
}

func (customer *Customer) printBalance() {
	fmt.Printf("Account Number: %d, Name: %s, Balance: %.2f\n", customer.accountNumber, customer.name, customer.balance)
}

func (bank *Bank) closeAccount(accountNumber int) {
	for i, customer := range bank.customers {
		if customer.accountNumber == accountNumber {
			bank.customers = append(bank.customers[:i], bank.customers[i+1:]...)
			fmt.Println("Account closed successfully.")
			return
		}
	}
	fmt.Println("Account not found.")
}

func main() {
	bank := Bank{}

	// Adding customers
	bank.addCustomer(1, "John Doe", 1000.0)
	bank.addCustomer(2, "Jane Smith", 5000.0)

	// Print total number of customers
	fmt.Println("Total Customers:", bank.totalCustomers())

	// Print all customer data
	bank.printAllCustomers()

	// Perform operations on individual accounts
	accountNumber := 1
	customer := bank.customers[accountNumber-1]

	fmt.Println("----- Account Operations -----")
	customer.printBalance()

	customer.withdraw(500.0)
	customer.deposit(1000.0)
	customer.printBalance()

	bank.closeAccount(accountNumber)
	bank.printAllCustomers()
}

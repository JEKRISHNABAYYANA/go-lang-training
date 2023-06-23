package main

import (
	"fmt"
)

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

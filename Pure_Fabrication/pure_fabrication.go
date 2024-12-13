//Explanation:
//Transaction Struct: Represents a bank transaction, focusing on domain-specific data.
//LogService (Pure Fabrication): Provides generic logging functionality that could be used by any system component. It does not map to a domain concept.
//BankAccount Struct: Manages transactions and utilizes the LogService to log actions, offloading logging responsibilities to a non-domain service.
//Client Code: Initializes BankAccount with a LogService and demonstrates how actions are logged separately from the business logic.

package main

import (
	"fmt"
	"time"
)

// Transaction represents a basic bank transaction
type Transaction struct {
	ID     string
	Amount float64
}

// LogService is a pure fabrication responsible for logging
type LogService struct{}

// Log logs messages with timestamps
func (ls *LogService) Log(message string) {
	fmt.Printf("%s: %s\n", time.Now().Format(time.RFC3339), message)
}

// BankAccount manages transactions and uses LogService to log actions
type BankAccount struct {
	ID           string
	transactions []Transaction
	logService   *LogService
}

// NewBankAccount creates a new bank account with the given ID
func NewBankAccount(id string, logService *LogService) *BankAccount {
	return &BankAccount{ID: id, logService: logService}
}

// AddTransaction adds a transaction to the bank account
func (ba *BankAccount) AddTransaction(transaction Transaction) {
	ba.transactions = append(ba.transactions, transaction)
	ba.logService.Log(fmt.Sprintf("Added transaction ID: %s, Amount: %.2f", transaction.ID, transaction.Amount))
}

// PrintTransactions prints all transactions in the bank account
func (ba *BankAccount) PrintTransactions() {
	ba.logService.Log("Printing all transactions")
	for _, transaction := range ba.transactions {
		fmt.Printf("Transaction ID: %s, Amount: %.2f\n", transaction.ID, transaction.Amount)
	}
}

func main() {
	logService := &LogService{}
	account := NewBankAccount("123456", logService)

	transaction1 := Transaction{ID: "T1", Amount: 100.0}
	transaction2 := Transaction{ID: "T2", Amount: 200.0}

	account.AddTransaction(transaction1)
	account.AddTransaction(transaction2)

	account.PrintTransactions()
}

// OUTPUT
//
//2024-12-13T22:38:34+02:00: Added transaction ID: T1, Amount: 100.00
//2024-12-13T22:38:35+02:00: Added transaction ID: T2, Amount: 200.00
//2024-12-13T22:38:35+02:00: Printing all transactions
//Transaction ID: T1, Amount: 100.00
//Transaction ID: T2, Amount: 200.00

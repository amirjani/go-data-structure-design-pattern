// Facade is a structural design pattern that provides a simplified interface to a library, a framework, or any other complex set of classes.
// Facade is used to abstract subsystem interfaces with a helper.
// The facade design pattern is used in scenarios when the number of interfaces increases and the system gets complicated.
// Facade is an entry point to different subsystems, and it simplifies the dependencies between the systems
// A loosely coupled principle can be realized with a facade pattern

// * The facade delegates the requests from the client to the module classes
// * Module classes implement the behaviors and functionalities of the module subsystem
// * The client invokes the facade method.

package main

import "fmt"

// ! => Account part
type Account struct {
	id          string
	accountType string
}

func (account *Account) create(accountType string) *Account {
	fmt.Println("Account creation with type")
	account.accountType = accountType
	return account
}

func (account *Account) getById(id string) *Account {
	fmt.Println("getting account by Id")
	return account
}

func (account *Account) deletedById(id string) {
	fmt.Println("delete account by id")
}

// ! => Customer Part
type Customer struct {
	id   string
	name string
}

func (customer *Customer) create(name string) *Customer {
	fmt.Println("creating customer")
	customer.name = name
	return customer
}

// ! => Transaction Part
type Transaction struct {
	id            string
	amount        float32
	srcAccountId  string
	destAccountId string
}

func (transaction *Transaction) create(srcAccountId string, destAccountId string, amount float32) *Transaction {
	fmt.Println("creating transaction")
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
	return transaction
}

type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

func (facade *BranchManagerFacade) createCustomerAccount(customerName string, accountType string) (*Customer, *Account) {
	var customer = facade.customer.create(customerName)
	var account = facade.account.create(accountType)
	return customer, account
}

func (facade *BranchManagerFacade) createTransaction(srcAccountId string,
	destAccountId string, amount float32) *Transaction {
	var transaction = facade.transaction.create(srcAccountId, destAccountId, amount)
	return transaction
}

func main() {
	var facade = NewBranchManagerFacade()
	var customer *Customer
	var account *Account

	customer, account = facade.createCustomerAccount("Amir Jani", "Saving")
	fmt.Println(customer.name)
	fmt.Println(account.accountType)
	var transaction = facade.createTransaction("111111", "222222", 100000)
	fmt.Println(transaction.amount)
}

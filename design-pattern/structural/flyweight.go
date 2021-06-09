// Flyweight is a structural design pattern that lets you fit more objects into the available amount of
// RAM by sharing common parts of state between multiple objects instead of keeping all of the data in each object.

// Flyweight is used to manage the state of an object with high variation
// Variable object data is referred to as extrinsic state, and the rest of the object state is intrinsic
// Flyweight pattern helps reduce the overall memory usage and the object initializing overhead

// * The FlyWeight interface has a method through which flyweights can get and act on the extrinsic state
// * ConcreteFlyWeight implements the FlyWeight interface to represent flyweight objects.
// * FlyweightFactory is used to create and manage flyweight objects
// * Client classes
package main

import "fmt"

type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	var dto = factory.pool[dtoType]
	if dto == nil {
		fmt.Println("new DTO of dtoType: " + dtoType)
		switch dtoType {
		case "customer":
			factory.pool[dtoType] = Customer{id: "1"}
		case "employee":
			factory.pool[dtoType] = Employee{id: "2"}
		case "manager":
			factory.pool[dtoType] = Manager{id: "3"}
		case "address":
			factory.pool[dtoType] = Address{id: "4"}
		}
		dto = factory.pool[dtoType]
	}
	return dto
}

// DataTransferObject interface
type DataTransferObject interface {
	getId() string
}

//Customer struct
type Customer struct {
	id   string //sequence generator
	name string
	ssn  string
}

// Customer class method getId
func (customer Customer) getId() string {
	fmt.Println("getting customer Id")
	return customer.id
}

//Employee struct
type Employee struct {
	id   string
	name string
}

//Employee class method getId
func (employee Employee) getId() string {
	return employee.id
}

//Manager struct
type Manager struct {
	id   string
	name string
	dept string
}

func (manager Manager) getId() string {
	return manager.id
}

//Address struct
type Address struct {
	id          string
	streetLine1 string
	streetLine2 string
	state       string
	city        string
}

//Address class method getId
func (address Address) getId() string {
	return address.id
}

//main method
func main() {
	var factory = DataTransferObjectFactory{make(map[string]DataTransferObject)}
	var customer DataTransferObject = factory.getDataTransferObject("customer")
	fmt.Println("Customer ", customer.getId())
	var employee DataTransferObject = factory.getDataTransferObject("employee")
	fmt.Println("Employee ", employee.getId())
	var manager DataTransferObject = factory.getDataTransferObject("manager")
	fmt.Println("Manager", manager.getId())
	var address DataTransferObject = factory.getDataTransferObject("address")
	fmt.Println("Address", address.getId())
}

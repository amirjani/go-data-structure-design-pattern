// Decorator is a structural design pattern that lets you attach new behaviors to objects by placing
// these objects inside special wrapper objects that contain the behavior

// The decorator pattern helps with subclassing when modifying functionality, instead of static inheritance
// An object can have multiple decorators and run-time decorators
// The single responsibility principle can be achieved using a decorator

// * The concrete component implements the component interface.
// * The decorator class implements the component interface and provides additional functionality in the same method or additional methods
package main

import "fmt"

type IProcess interface {
	process()
}

type ProcessClass struct{}

func (process *ProcessClass) process() {
	fmt.Println("ProcessClass process")
}

type ProcessDecorator struct {
	processInstance *ProcessClass
}

func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Printf("ProcessDecorator process and ")
		decorator.processInstance.process()
	}
}

func main() {
	var process = &ProcessClass{}
	var decorator = &ProcessDecorator{}

	decorator.process()
	decorator.processInstance = process
	decorator.process()
}

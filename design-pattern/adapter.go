// The adapter pattern provides a wrapper with an interface required by the API client to link incompatible types and
// act as a translator between the two types.

// The dependency inversion principle can be adhered to by using the adapter pattern,
// when a class defines its own interface to the next level module interface implemented by an adapter class.

// Multiple formats handling source-to-destination transformations are the scenarios where the adapter pattern is applied

// * Target is the interface that the client calls and invokes methods on the adapter and adaptee
// * The client wants the incompatible interface implemented by the adapter.
// * The adapter translates the incompatible interface of the adaptee into an interface that the client wants.
package main

import "fmt"

type IProcess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}

type Adaptee struct {
	adapterType int
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

func (adapter Adapter) process() {
	fmt.Println("Adapter Process")
	adapter.adaptee.convert()
}

func main() {
	var processor IProcess = Adapter{}
	processor.process()
}

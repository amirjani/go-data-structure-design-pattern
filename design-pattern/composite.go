// A composite is a group of similar objects in a single object.
// Objects are stored in a tree form to persist the whole hierarchy
// New types of objects can be added without changing the interface and the client code.

// * The component interface defines the default behavior of all objects and behaviors for accessing the components of the composite.
// * The composite and component classes implement the component interface.
// * The client interacts with the component interface to invoke methods in the composite.

package main

import "fmt"

type IComposite interface {
	perform()
}

type Leaflet struct {
	name string
}

func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet " + leaf.name)
}

type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

func (branch *Branch) perform() {
	fmt.Println("Branch: " + branch.name)

	for _, leaf := range branch.leafs {
		leaf.perform()
	}

	for _, branch := range branch.branches {
		branch.perform()
	}
}

func (branch *Branch) add(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

func main() {
	var branchOne = &Branch{name: "Branch One"}
	var branchTwo = &Branch{name: "Branch Two"}

	var leafOne = &Leaflet{name: "Leaf One"}
	var leafTwo = &Leaflet{name: "Leaf Two"}

	branchOne.add(*leafOne)
	branchOne.add(*leafTwo)
	branchOne.addBranch(*branchTwo)
	branchOne.perform()
}

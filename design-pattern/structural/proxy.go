// The proxy pattern forwards to a real object and acts as an interface to others.
// The proxy pattern controls access to an object and provides additional functionality.
// Remote, smart, virtual, and protection proxies are the scenarios where this pattern is applied.

// * Subject is an interface for the RealObject and Proxy class
// * The RealSubject object is created and maintained as a reference in the Proxy class.
// * VirtualProxy is used to access RealObject and invoke the performAction method.
package main

import "fmt"

type IRealObject interface {
	performAction()
}

type RealObject struct{}

func (realObject *RealObject) performAction() {
	fmt.Println("RealObject performAction()")
}

type VirtualProxy struct {
	realObject *RealObject
}

func (virtualProxy *VirtualProxy) performAction() {
	if virtualProxy.realObject == nil {
		virtualProxy.realObject = &RealObject{}
	}
	fmt.Println("Virtual Proxy performAction()")
	virtualProxy.realObject.performAction()
}

func main() {
	var object VirtualProxy = VirtualProxy{}
	object.performAction()
}

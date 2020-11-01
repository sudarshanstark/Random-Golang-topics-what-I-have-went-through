// Empty Interfaces : Interface with no method.
// Uses : Empty interfaces are used by code that handles values of unknown type.
// Any Go type satisfies this interface and that means it can represent any value.
// An Empty Interface may hold values of any type.
// You cannot use directly an empty interface in opoerations.
// You can pass an emty interface type as a function parameter of any type.
/* Interfaces stores two values
1) Dynamic Type => To access a dynamic value we have to do an assertion.
2) Dynamic Concrete Value
*/

package main

import "fmt"

type EmptyInterface interface {
}

type person struct {
	info interface{} // Demonstrating Empty Interface here.
	// Here info field is of type Empty Interface. That means you can store a any values into this info field.
}

func main() {
	var empty interface{}
	empty = 5
	fmt.Println(empty)
	empty = "go"
	fmt.Println(empty)
	empty = []int{1, 2, 3}
	fmt.Println(empty)
	fmt.Println(len(empty.([]int))) // Using type assertion here. The type you have to assert is enclosed in paranthesis.
	// .([]int) is the asserted type here
	// Here len() function accepts "empty" array as an argument of dynamic type.
	// So you have to use type assertion.

	you := person{}
	you.info = "Your name"
	you.info = 40
	you.info = []float64{5.6, 2., 2.6}
	fmt.Println(you.info)
}

// This is the power of empty interfaces.
// Everythig that is so powerful comes with a cost.
// Empty interfaces can cause programs to become hard to maintain.
// The empty interface type is increasingly being misused as a connvinient way to bypass go-compilers type checking.
// The one principle of Go is that it allows to write type safe code.
// There are places where empty interface is used in place where explicit interface could have been used.
// The problem with runtime type checking is that you will never know there is a problem untill it is run.

package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
)


// variables
var (
	/*name, course string
	module float64*/
	//name, course, module = "Nigel", "Docker Deep Dive", 3.2
	name = "Nigel"
	course = "Docker Deep Dive"
	module = 3.2
)

// hello
/*
	hello
*/
func main() {
	fmt.Println("Hello from", runtime.GOOS)

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("course is", course, "and is of type", reflect.TypeOf(course))
	fmt.Println("MOdule is", module, "and is of type", reflect.TypeOf(module))

	// only works in functions and you have to declare them
	a := 10.000
	b := 3
	// need to use these variables or it won't compile

	fmt.Println("Value of a is", a, "and is of type", reflect.TypeOf(a))
	fmt.Println("Value of b is", b, "and is of type", reflect.TypeOf(b))

	c := int(a) + b //cannot add1 float and int
	fmt.Println("Value of c is", c, "and is of type", reflect.TypeOf(c))
	fmt.Println("Value of c is", c, "and its memory address is", &c)

	// using pointers
	ptr := &module
	fmt.Println("Memory address of *c* variable is", ptr, "and the value of *c* is", *ptr)

	changeString(&course)
	fmt.Println("new course value =", course)

	const someConstant = "constantThing"


	// environment vars
	fmt.Println(os.Environ())

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	ps1 := os.Getenv("PS1")
	fmt.Println(ps1)

}

// pass by reference
func changeString(str *string) string {
	*str = "Some other string from changeString(...)"
	return *str
}
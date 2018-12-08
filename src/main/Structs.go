package main

import "fmt"

// custom types

func main() {

	type User struct {
		name string
		age int
	}

	blankUser := new(User) //pointer
	fmt.Println(blankUser)

	// don't have to name the fields if you put them
	// in the same order
	Anthony := User{
		name: "Anthony",
		age: 30,
	}

	fmt.Println(Anthony)

	fmt.Println(Anthony.name)
	fmt.Println(Anthony.age)

	Anthony.age = 31
	fmt.Println(Anthony)

}

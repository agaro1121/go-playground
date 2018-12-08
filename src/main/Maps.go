package main

import "fmt"

func main() {

	// pass by Ref


	stuffs := make(map[int]string)
	stuffs[0] = "zero"
	stuffs[1] = "one"

	fmt.Println(stuffs)

	stuffs2 := map[int]string {
		2: "two",
		3: "three",
	}


	fmt.Println(stuffs2[2])
	fmt.Println(stuffs2)

	for k, v := range stuffs2 {
		fmt.Printf("\nkey= %v value= %v", k, v)
	}

	stuffs2[2] = "twoz"
	fmt.Println(stuffs2)

	//upsert
	stuffs2[4] = "four"
	fmt.Println(stuffs2)

	// delete
	delete(stuffs2, 4)
	fmt.Println(stuffs2)

}
package main

// only 'for' keyword
// you can emulate while loops with this

func main() {

	for i := 0; i < 10; i++ {
		println(i)
		if i == 7 {
			break
		}
	}

	numbers := []int{1,2,3,4,5}

	for index, value := range numbers {
		println("index=", index, "value=", value)
	}

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		println("only printing odd=",i)
	}

	// can use labels to break out of nested loops
	outerloop:
		for i := 0; i < 10; i++ {
			if i % 2 == 0 {
				continue
			}
			println("only printing odd=",i)
			if i == 9 {
				println("Running loop again internally and breaking out")
				for i := 0; i < 5; i++ {
					println(i)
				}
				break outerloop
			}
		}

}

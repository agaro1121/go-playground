package main

func main() {
	println("Hello World")
	variadiableArgs(1,2,3,4,4)

	// inline functions
	println("value =", func (stuff ...int) int {
		for _, i := range stuff {
		println(i)
	}
		return stuff[0]
	}(1,2,3))

	println(add1(1))

	// HO function with function literal
	// no lambda expression syntactic sugar
	println(doMath(1,2, func(x, y int) int {
		return x + y
	}))

	numTerms, result := add(1,2,3,4)
	println("num of terms=", numTerms, "result=", result)

	numTerms2, result2 := add2(1,2,3,4)
	println("num of terms=", numTerms2, "result=", result2)

}

func variadiableArgs(stuff ...int) int {
	for _, i := range stuff {
		println(i)
	}
	return stuff[0]
}

//method
type Count int
func (c Count) Incr() int {
	c = c + 1
	return int(c)
}

//first class functions
var add1 = func(n int) int {
	return n + 1
}

//higher order functions
func doMath(x, y int, calc func (int, int) int) int {
	return calc(x, y)
}

// return multiple values
func add(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}
	return len(terms), result
}

// return multiple mixed named values
func add2(terms ...int) (int, result int) {
	for _, term := range terms {
		result += term
	}
	return len(terms), result
}


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
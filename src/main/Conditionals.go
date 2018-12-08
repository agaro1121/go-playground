package main

import (
	"math/rand"
	"time"
)

func main() {

	if 3 < 2 {
		println("true")
	} else {
		println("false")
	}

	// simple statements
	if x, y := 1,2 ; x < y {
		println("true")
	} else {
		println("false")
	}

	// switch
	// default can be put anywhere. Doesn't have to be last
	// no fall through. Other languages need 'break'. 'break' is implicit
	switch "docker" {

	case "linux":
		println("linux")

	case "docker":
		println("docker")

	case "windows":
		println("windows")

	default:
		println("default")

	}

	snitches(2)
	fallthroughs(2)

	multipleCases()
}

func snitches(x int) {
	switch x {
	case 1:
		println("1")
	case 2:
		println("2")
	case 3:
		println("3")
	default:
		println("default")
	}
}

/*
	Will run the statement after 2
*/
func fallthroughs(x int) {
	switch x {
	case 1:
		println("1")
	case 2:
		println("2")
		fallthrough
	case 3:
		println("3")
		//fallthrough // will print 2, 3, default if this were uncommented
	default:
		println("default")
	}
}

func multipleCases() {
	switch tmpNum := random(); tmpNum {
		case 0,2,4,6,8:
			println("even=", tmpNum)
		case 1,3,5,7,9:
			println("odd=", tmpNum)
		default:
			println("unexpected=", tmpNum)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
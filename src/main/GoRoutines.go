package main

import "time"

func main() {
	go func(){
		println("Saluton")
	}()

	go func(){
		println("Mondo")
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}

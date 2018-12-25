package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(4)

	mutex := make(chan bool, 1)

	f, _ := os.Create("./log.txt")
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <- logCh
			if ok {
				f, e := os.OpenFile("./log.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				println("error=", e)

				logTime := time.Now().Format(time.RFC3339)
				r, err := f.WriteString(logTime + " - " + msg)
				fmt.Println("result=", r, "error=", err)
				f.Close()
			} else {
				fmt.Println("not ok")
				break
			}
		}
	}()

	for i := 0; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex <- true
			go func() {
				msg := fmt.Sprintf("%d + %d = %d\n", i,j,i+j)
				logCh <- msg
				fmt.Println(msg)
				<- mutex
			}()
		}
	}

	fmt.Scanln()

}
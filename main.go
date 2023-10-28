package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	for {
		password := ""
		fmt.Println("Welcome to Guess The Password Game!")
		fmt.Scanln(&password)
		if password != "BlackOps1337" {
			fmt.Println("Wrong")
		} else {
			fmt.Println("Congrats on winning")
			//fmt.Scanln()
			time.Sleep(3 * time.Second)
			os.Exit(0)
		}
	}

}

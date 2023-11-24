package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Please select a question #: ")

	// read input using go std lib bufio, getting input from cmdline
	reader := bufio.NewReader(os.Stdin)

	// read in the user's selected question number
	questionNumber, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	
	message := fmt.Sprintf("You selected question %v", questionNumber)
	fmt.Println(message)
}

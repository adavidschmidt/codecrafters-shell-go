package main

import (
	"bufio"
	//Uncomment this block to pass the first stage
	"fmt"
	"os"
)

func main() {
	//Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	bufio.NewReader(os.Stdin).ReadString('\n')
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Fprint(os.Stdout, input[:len(command)-1]+": command not found\n")
}

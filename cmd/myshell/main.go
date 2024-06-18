package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func main() {
	//Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Printf("%s: command not found\n", strings.TrimRight(command, "\n"))
}

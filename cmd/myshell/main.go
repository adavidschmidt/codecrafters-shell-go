package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
)

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func runEcho(a []string) {
	b := strings.Join(a, " ")
	fmt.Fprintf("%s\n", b)
}

func runExit() {
	os.Exit(0)
}

func runType(a string, b []string) {
	if stringInSlice(a, b) {
		fmt.Fprintf("%s is a shell builtin\n", a)
	} else {
		paths := strings.Split(os.Getenv("PATH"), ":")
		for _, path := range paths {
			dir := path + "/" + a
			if _, err := os.Stat(dir); err == nil {
				fmt.Fprintf("%s is %v\n", a, dir)
				return
			}
		}
		fmt.Fprintf("%s: not found\n", a)
	} 	
}

func runCode(a string, b string) {
	paths := strings.Split(os.Getenv("PATH"), ":")
		for _, path := range paths {
			dir := path + "/" + a
			if _, err := os.Stat(dir); err == nil {
				c := strings.Split(a, "_")
				fmt.Fprintf("Hello %s! The secret code is %v\n", b, c[1])
				return
			}
		}
		fmt.Fprintf("%s: not found\n", a)
}

func main() {


	// Wait for user input
	reader := bufio.NewReader(os.Stdin)
	listCommands := []string{"echo", "type", "exit"}
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, _ := reader.ReadString('\n')
		tokens := strings.Fields(strings.TrimSpace(command))
		command = strings.Join(tokens, " ")
		switch {
		case tokens[0] == "exit" && tokens[1] == "0":
			runExit()
			
		case tokens[0] == "echo":
			runEcho(tokens[1:])
		
		case tokens[0] == "type":
			runType(tokens[1], listCommands)

		case strings.Contains(tokens[0], "_"):
			runCode(tokens[0], tokens[1])

		default:
			fmt.Fprintf("%s: command not found\n", command)
		}
	}
}

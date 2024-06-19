package main

import (
	"bufio"
	"strings"
	"fmt"
	"os"
	"os/exec"
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
	fmt.Printf("%s\n", b)
}

func runExit() {
	os.Exit(0)
}

func runType(a string, b []string) {
	if stringInSlice(a, b) {
		fmt.Printf("%s is a shell builtin\n", a)
	} else {
		paths := strings.Split(os.Getenv("PATH"), ":")
		for _, path := range paths {
			dir := path + "/" + a
			if _, err := os.Stat(dir); err == nil {
				fmt.Printf("%s is %s\n", a, dir)
				return
			}
		}
		fmt.Printf("%s: not found\n", a)
	} 	
}

func runCode(a string) bool {
	tokens := strings.Split(a, " ")
	if len(tokens) > 1 {
		paths := strings.Split(os.Getenv("PATH"), ":")
			for _, path := range paths {
				dir := path + "/" + tokens[0]
				if _, err := os.Stat(dir); err == nil {
					cmd := exec.Command(tokens[0], tokens[1])
					cmd.Stdout = os.Stdout
					cmd.Stdin = os.Stdin
					if err := cmd.Run(); err != nil {
						fmt.Printf("Error: %s", err)
					}
					return true
				}
			}
			
			return false
		} else {
			return false
		}
}

func pwdGet() {
	directory, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(directory)
}

func changeDir(a string) {
	tokens := strings.Split(a, " ")
	if len(tokens) > 1 {
		err := os.Chdir(tokens[1])
		if err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", tokens[1])
		}
	}
}

func main() {


	// Wait for user input
	reader := bufio.NewReader(os.Stdin)
	listCommands := []string{"echo", "type", "exit", "pwd", "cd"}
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

		case tokens[0] == "pwd":
			pwdGet()

		case tokens[0] == "cd"
			changeDir(command)
			
		default:
			if !runCode(command) {
				fmt.Printf("%s: command not found\n", command)
			}
		}	
	}
}

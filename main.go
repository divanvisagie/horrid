package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func runCommand(command string) string {
	commandStructure := strings.Split(command, " ")

	args := commandStructure[1:]
	fmt.Println(args)

	spacedArgs := strings.Join(args, " ")

	commandOutput, err := exec.Command(commandStructure[0], spacedArgs).Output()
	if err != nil {
		return err.Error()
	}
	return string(commandOutput)
}

func main() {

	output := runCommand("cat main.go")

	fmt.Println(output)
}

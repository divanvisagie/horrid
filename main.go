package main

import (
	"fmt"
	"os/exec"
)

func runCommand(command string) string {
	commandOutput, err := exec.Command(command).Output()
	if err != nil {
		return err.Error()
	}
	return string(commandOutput)
}

func main() {

	output := runCommand("ls")

	fmt.Println(output)
}

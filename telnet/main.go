package main

import (
	"fmt"
	"os/exec"
)

func main() {
	app := "telnet"

	cmd := exec.Command(app, "postal.pamachef.com 25")
	stdout, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	// Print the output
	fmt.Println(string(stdout))
}

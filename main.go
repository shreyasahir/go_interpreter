package main

import (
	"fmt"
	"go_interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome %s to Monkey programming version 1.0 ", user.Username)
	fmt.Printf("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}

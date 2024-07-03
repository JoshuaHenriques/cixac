package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/joshuahenriques/cixac/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Cixac programming language!\n",
		user.Username)
	fmt.Printf("Type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

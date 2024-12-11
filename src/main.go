package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/singlaanish56/Interpreter-In-Go/repl"
)

func main() {

	user, err := user.Current()
	if err != nil{
		panic(err)
	}

	fmt.Printf("Welcome to the monkey business : %s\n", user.Username)

	fmt.Println("go ahead type something")

	repl.Start(os.Stdin, os.Stdout)
}
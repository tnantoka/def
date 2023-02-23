package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"def/repl"
)

var engine = flag.String("engine", "vm", "use 'vm' or 'eval'")

func main() {
	flag.Parse()

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Def programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")

	fmt.Printf("Engine: %s\n", *engine)

	if *engine == "vm" {
		repl.StartVM(os.Stdin, os.Stdout)
	} else {
		repl.StartEval(os.Stdin, os.Stdout)
	}
}

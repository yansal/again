package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fail := flag.Bool("fail", false, "run until command fails")
	flag.Parse()
	name := flag.Arg(0)
	if name == "" {
		fmt.Println("usage:", os.Args[0], "name [args...]")
		os.Exit(1)
	}
	for {
		cmd := exec.Command(name, flag.Args()[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil && !*fail {
			log.Print(err)
		} else if err != nil || !*fail {
			break
		}
	}
}

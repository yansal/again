package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
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

		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		err = cmd.Wait()
		if err != nil {
			log.Print(err)
			time.Sleep(time.Second)
		} else {
			break
		}
	}
}

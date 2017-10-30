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
	log.SetFlags(log.Lshortfile)
	log.SetPrefix("again: ")

	fail := flag.Bool("fail", false, "run until command fails")
	sleep := flag.Duration("sleep", 0, "how long to sleep before running again?")
	flag.Parse()
	cmdname := flag.Arg(0)
	if cmdname == "" {
		fmt.Println("usage:", "again cmdname [cmdargs...]")
		os.Exit(1)
	}
	for {
		cmd := exec.Command(cmdname, flag.Args()[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil && !*fail {
			log.Print(err)
		} else if err != nil || !*fail {
			break
		}

		if *sleep != 0 {
			time.Sleep(*sleep)
		}
	}
}

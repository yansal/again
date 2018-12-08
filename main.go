package main

import (
	"bytes"
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
	silent := flag.Bool("silent", false, "don't print failed attempts error message")
	flag.Parse()

	cmdname := flag.Arg(0)
	if cmdname == "" {
		fmt.Println("usage:", "again [options] cmdname [cmdargs...]")
		os.Exit(1)
	}

	errbuf := new(bytes.Buffer)

	for {
		cmd := exec.Command(cmdname, flag.Args()[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = errbuf

		err := cmd.Run()
		if err != nil && !*fail {
			if !*silent {
				log.Print(err)
				fmt.Print(errbuf.String())
				errbuf.Reset()
			}
		} else if err != nil || !*fail {
			break
		}

		if *sleep != 0 {
			time.Sleep(*sleep)
		}
	}
}

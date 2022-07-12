package main

import (
	"os"
	"strconv"

	"github.com/alessio-pareto/process"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	p, err := os.FindProcess(pid)
	if err != nil {
		os.Exit(1)
	}

	err = process.StopProcess(p)
	if err != nil {
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"
)

type Config struct {
	Tools map[string]string
}

func main() {
	var path string
	if len(os.Args) < 2 {
		path = ".mise.toml"
	} else {
		path = os.Args[1]
	}
	raw, err := os.ReadFile(path)
	if err != nil && !os.IsNotExist(err) {
		die(err)
	}
}

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

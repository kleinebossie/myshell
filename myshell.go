package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		/*
			host, err := os.Hostname()
			if err != nil {
				fmt.Fprintln(os.Stderr)
			}

			fmt.Printf("\n---%v---\n", host)
		*/

		pathArr := strings.Split(path, "/")

		fmt.Print("~")

		for i := 2; i < len(pathArr); i++ {
			if i == 2 || i == len(pathArr)-1 {
				fmt.Printf("/%v", pathArr[i])
				continue
			}
			fmt.Printf("/%v", (pathArr[i])[0:1])
		}

		fmt.Print(" > ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

var ErrNoPath = errors.New("path required")

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return ErrNoPath
		}

		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

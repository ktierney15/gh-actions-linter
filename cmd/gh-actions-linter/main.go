package main

import (
	"os"
	"fmt"

	"github.com/ktierney15/gh-actions-linter/internal/lint"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage: gha-lint <command> [options]")
		os.Exit(1)
	}

	fileName := args[0]

	// err := lint.Run(fileName)
	lint.Run(fileName)

	// if err != nil {
	// 	fmt.Println("Error Linting File:", err)
	// 	os.Exit(1)
	// }

	fmt.Println(fileName, "Linted Sucessfully")


}
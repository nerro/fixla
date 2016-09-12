package main

import (
	"fmt"
	"os"
)

func init() {
	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}
	// check for '--version' and '--help' options
	for _, arg := range os.Args {
		if arg == "--version" || arg == "-v" {
			fmt.Println(FormattedVersion())
			os.Exit(0)
		}
		if arg == "--help" || arg == "-h" {
			printHelp()
			os.Exit(0)
		}
	}
}

func main() {
	fmt.Println("to be implemented...")
}

// printHelp provides a pretty usage message including a list of the correct
// commands, arguments and options acceptable to the program.
func printHelp() {
	fmt.Println("usage: fixla [--version] [--help] <command> [<args>]")
	fmt.Println()
	fmt.Println("Available commands are:")
	fmt.Println("   trailing-whitespaces   Handle trailing whitespaces")
	fmt.Println()
}

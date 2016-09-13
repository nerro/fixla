package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// define available commands
	trailingWhitespacesCommand := flag.NewFlagSet("trailing-whitespaces", flag.ExitOnError)
	fileFlag := trailingWhitespacesCommand.String("file", "", "file to check")

	if len(os.Args) == 1 {
		printHelp()
		os.Exit(0)
	}
	// check for using '--version' and '--help' options
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

	switch os.Args[1] {
	case "trailing-whitespaces":
		trailingWhitespacesCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("fixla: '%s' is not a fixla command. See 'fixla --help'\n", os.Args[1])
		os.Exit(1)
	}

	if trailingWhitespacesCommand.Parsed() {
		if *fileFlag == "" {
			fmt.Println("Please define the file using -path option.")
		} else {
			fmt.Println("//TODO: logic for trailing whatespaces processing here")
		}
	}
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

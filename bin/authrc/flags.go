package main

import (
	"fmt"
	"os"

	"github.com/gotamer/authrc/bin/authrc/build"
)

var (
	arg_Name, arg_Path    string
	arg_Env, arg_Template bool
	arg_Debug, arg_All    bool
)

func flags() {
	if len(os.Args) == 1 {
		print(help)
		os.Exit(1)
	} else {

		for no, arg := range os.Args {
			switch arg {
			case "-h", "--help":
				print(help)
				os.Exit(0)
			case "--version":
				fmt.Printf("%s\n", build.Info())
				os.Exit(0)
			case "-v":
				fmt.Printf("%s\n", build.Info())
				os.Exit(0)
			case "-d", "--debug", "--verbose":
				arg_Debug = true
			case "-e", "--environment":
				arg_Env = true
			case "-t", "--template":
				arg_Template = true
			case "-a", "--all":
				arg_All = true
			case "-n", "--name":
				arg_Name = os.Args[no+1]
			case "-p", "--path":
				arg_Path = os.Args[no+1]
			}
		}
	}
}

const help = `
authrc - The friendly authentication runcom

SYNOPSIS
	authrc [OPTIONS]

DESCRIPTION
authrc is a command line tool to read .authrc files.
The .authrc file contains login and initialization information used for auto login.

OPTIONS

	-n or --name
           Service or Alias

	-p or --path  '~/.config/authrc'
	       Path to the '.authrc file'

	-e or --environment [bool]
	       Will set env variables.
	       AUTHRC_SERVIVE, AUTHRC_USERNAME, AUTHRC_PASSWORD, AUTHRC_COMMAND

	-h or --help [bool]
	       Report usage information and exit.

	-v or --version [bool]
	       Print version and exit.

	-t or --template [bool]
	       Print template authrc line and exit

	--verbose [bool]
	       Alias for --debug

	-d or --debug [bool]
	       Print verbose information about the process.

COPYRIGHT
	MIT License © 2023 TaMeR
`

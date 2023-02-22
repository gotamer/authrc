// authrc - The friendly authentication runcom command
// authrc is a command line tool to read .authrc files.
// The .authrc file contains login and initialization information used for auto login.
package main

import (
	"flag"
	"fmt"
	"os"

	"git.hansaray.pw/go/authrc"
)

const help = `
authrc - The friendly authentication runcom

SYNOPSIS
	authrc [OPTIONS]

DESCRIPTION
authrc is a command line tool to read .authrc files.
The .authrc file contains login and initialization information used for auto login.

OPTIONS

	-name  Service or Alias

	-path  '~/.config/authrc'
	       Path to the '.authrc file'

	-env [bool]
	       Will set env variables.
	       AUTHRC_SERVIVE, AUTHRC_USERNAME, AUTHRC_PASSWORD, AUTHRC_COMMAND

	-h or --help [bool]
	       Report usage information and exit.

	-v or --version [bool]
	       Print version and exit.

	-t [bool]
	       Print template authrc line and exit

	--verbose [bool]
	       Alias for --debug=files

	-d [bool]
	       Print verbose information about the process.

REPORTING BUGS
	https://git.hansaray.pw/go/authrc

COPYRIGHT
	MIT License © 2023 TaMeR
`

var (
	path = flag.String("path", "", "path to the '.authrc file'")
	name = flag.String("name", "", "name of Service or Alias")
	env = flag.Bool("env", false, "set env variables")
	all = flag.Bool("all", false, "print all")
	debug = flag.Bool("d", false, "print debug info")
	template = flag.Bool("t", false, "print template line and exit")
)

func init() {
	if len(os.Args) == 1 {
		print(help)
		os.Exit(1)
	} else {
		switch os.Args[1] {
		case "-h", "--help", "help":
			print(help)
			os.Exit(0)
		case "--version", "version":
			fmt.Printf("Version: %s Date: %s\n", VerTag, VerModDate)
			os.Exit(0)
		case "-v":
			print(VerTag)
			os.Exit(0)
		}
	}
	flag.Parse()
}

func main() {
	var err error

	if *template {
		var a string
		if *name != "" {
			a = *name
		}
		fmt.Printf("ALIAS: %s, SERVICE: , USERNAME: , PASSWORD: , COMMAND: ", a)
		os.Exit(0)
	}

	if *name == "" {
		print(help)
		os.Exit(1)
	}

	err = authrc.SetFilePath(*path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	s, u, p, c := authrc.GetAll(*name)
	if *env {
 		os.Setenv("AUTHRC_SERVIVE", s)
		os.Setenv("AUTHRC_USERNAME", u)
		os.Setenv("AUTHRC_PASSWORD", p)
		os.Setenv("AUTHRC_COMMAND", c)
	} else if *all {
		fmt.Printf("%s %s %s %s\n", s, u, p, c)
	}else{
		fmt.Printf("%s\n", p)
	}

	if *debug {
		fmt.Printf("ALIAS: %s, SERVICE: %s, USERNAME: %s, PASSWORD: %s, COMMAND: %s", *name, s, u, p, c)
	}
}

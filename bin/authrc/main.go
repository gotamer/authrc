// authrc project main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"git.hansaray.pw/go/authrc"
)

const help = `
authrc - The friendly authentication runcom

Usage:
	authrc [OPTIONS] [alias, service]

DESCRIPTION
authrc is a command line tool to read .authrc files.
The .authrc file contains login and initialization information used for auto login.

The following options are available:

-name  Service or Alias

-path  '~/.config/authrc'
       Path to the '.authrc file'

-env   Set env variables.
       AUTHRC_SERVIVE, AUTHRC_USERNAME, AUTHRC_PASSWORD

-h or --help
       Print this help and exit.

-v or --version
       Print version and exit.

-t     Print template authrc line and exit

-d     Print verbose information about the process.

More info at: https://git.hansaray.pw/go/authrc

`

var (
	path = flag.String("path", "", "path to the '.authrc file'")
	name = flag.String("name", "", "name of Service or Alias")
	env = flag.Bool("env", false, "set env variables")
	all = flag.Bool("all", false, "print all")
	debug = flag.Bool("d", false, "print all")
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
		case "-v", "--version", "version":
			print("not implemented, comming soon")
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
		fmt.Printf("ALIAS: %s, SERVICE: , USERNAME: , PASSWORD: ", a)
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

	s, u, p := authrc.GetAll(*name)
	if *env {
 		os.Setenv("AUTHRC_SERVIVE", s)
		os.Setenv("AUTHRC_USERNAME", u)
		os.Setenv("AUTHRC_PASSWORD", p)
	} else if *all {
		fmt.Printf("%s %s %s\n", s, u, p)
	}else{
		fmt.Printf("%s\n", p)
	}

	if *debug {
		fmt.Printf("ALIAS: %s, SERVICE: %s, USERNAME: %s, PASSWORD: %s", *name, s, u, p)
	}
}
// authrc - The friendly authentication runcom command
// authrc is a command line tool to read .authrc files.
// The .authrc file contains login and initialization information used for auto login.
package main

import (
	"fmt"
	"os"

	"github.com/gotamer/authrc"
)

func main() {
	flags()
	var err error

	if arg_Template {
		fmt.Printf("ALIAS: %s, SERVICE: , USERNAME: , PASSWORD: , COMMAND: ", arg_Name)
		os.Exit(0)
	}

	if arg_Name == "" {
		print(help)
		os.Exit(1)
	}

	if err = authrc.SetFilePath(arg_Path); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	s, u, p, c := authrc.GetAll(arg_Name)
	if arg_Env {
		os.Setenv("AUTHRC_SERVIVE", s)
		os.Setenv("AUTHRC_USERNAME", u)
		os.Setenv("AUTHRC_PASSWORD", p)
		os.Setenv("AUTHRC_COMMAND", c)
	} else if arg_All {
		fmt.Printf("%s %s %s %s\n", s, u, p, c)
	} else {
		fmt.Printf("%s\n", p)
	}

	if arg_Debug {
		fmt.Printf("ALIAS: %s, SERVICE: %s, USERNAME: %s, PASSWORD: %s, COMMAND: %s", arg_Name, s, u, p, c)
	}
}

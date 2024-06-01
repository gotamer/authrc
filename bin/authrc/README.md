[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/gotamer/authrc/bin/authrc?status.svg)](https://godoc.org/github.com/gotamer/authrc/bin/authrc)

## authrc - The friendly authentication runcom

```bash
authrc - The friendly authentication runcom

SYNOPSIS
	authrc [OPTIONS]

DESCRIPTION
authrc is a command line tool to read .authrc files.
The .authrc file contains login and initialization information used for auto login.

OPTIONS

	-n or --name
           Service or Alias

	-l or --list
           List all Aliases

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

```

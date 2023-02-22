[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/git.hansaray.pw/go/authrc/bin/authrc?status.svg)](https://godoc.org/git.hansaray.pw/go/authrc/bin/authrc)

## authrc - The friendly authentication runcom

```bash
authrc - The friendly authentication runcom

Usage:
	authrc [OPTIONS]

DESCRIPTION
authrc is a command line tool to read .authrc files.
The .authrc file contains login and initialization information used for auto login.

The following options are available:

-name  Service or Alias

-path  '~/.config/authrc'
       Path to the .authrc file

-env   Set env variables.
       AUTHRC_SERVIVE, AUTHRC_USERNAME, AUTHRC_PASSWORD, AUTHRC_COMMAND

-h or --help
       Print this help and exit.

-v or --version
       Print version and exit.

-t     Print template authrc line and exit.

-d     Print verbose information about the process.

More info at: https://git.hansaray.pw/go/authrc

```

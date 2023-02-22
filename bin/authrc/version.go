package main

import (
	"fmt"
	"os"
)

func init() {
	if len(os.Args) > 1 {
		var arg = os.Args[1]
		if arg == "-v" || arg == "--version" || arg == "version" {
			fmt.Printf("Version Date: %s\n", VerModDate)
			fmt.Printf("Version Tag: %s\n", VerTag)
			fmt.Printf("Version Git: %s\n", VerGit)
			os.Exit(0)
		}
	}
}

//VerModTime is a UTC Unix time stamp
const VerModTime = 1677084911

//VerModDate is a UTC Date time stamp
const VerModDate = "2023-02-22T16:55:11Z"

//VerLong is the full version from Git command output
const VerLong = "v0.0.6-0-g7860c4d"

//VerDirty means app was build with a git dir that contained modifications which had not been committed.
const VerDirty = false

//VerGit is the 7 hexadecimal digits version from Git.
const VerGit = "g7860c4d"

//VerTag is the Tag version from Git.
const VerTag = "v0.0.6"

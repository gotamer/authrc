// authrc - The friendly authentication runcom library
// The .authrc file contains login and initialization information used for auto login.
package authrc

import (
	"io/ioutil"
	"strings"
	"os"
	"path/filepath"
)

var (
	FILE_AUTHRC string
	database []*authrcLine
)

type authrcLine struct {
	alias  string
	service    string
	username    string
	password string
	command string
}

func SetFilePath(path string) (err error) {
	if path == "" {
		var dir_user_cfg string
		if dir_user_cfg, err = os.UserConfigDir(); err != nil {
			return err
		}
		FILE_AUTHRC = filepath.Join(dir_user_cfg, "authrc")
	}else{
		FILE_AUTHRC = path
	}
	var data []byte
	if data, err = load(); err != nil {
		return err
	}
	parse(string(data))
	return
}

// Returns service, username, password and command
func GetAll(name string) (string, string, string, string) {
	authrc := get(name)
	if authrc != nil {
		return authrc.service, authrc.username, authrc.password, authrc.command
	}
	return "", "", "", ""
}

// Returns both username and password
func Get(name string) (string, string) {
	authrc := get(name)
	if authrc != nil {
		return authrc.username, authrc.password
	}
	return "", ""
}

func GetService(name string) string {
	authrc := get(name)
	if authrc != nil {
		return authrc.service
	}
	return ""
}

func GetPassword(name string) string {
	authrc := get(name)
	if authrc != nil {
		return authrc.password
	}
	return ""
}

func GetUsername(name string) string {
	authrc := get(name)
	if authrc != nil {
		return authrc.username
	}
	return ""
}

func get(name string) (item *authrcLine) {
	for _,item = range database {
		if item.alias == name || item.service == name {
			return
		}
	}
	item = new(authrcLine)
	return
}

func load() (data []byte, err error){

	if data, err = ioutil.ReadFile(FILE_AUTHRC); err != nil {
		if os.IsNotExist(err) {
			//log.Fatalf("no netrc file found at %s: %v\n", file_user_authrc, err)
			return
		}
		return
	}
	return
}

func parse(data string) () {

	for _, line := range strings.Split(data, "\n") {

		if len(line) == 0 {
			continue
		}

		var authrc = new(authrcLine)

		for _, fields := range strings.Split(line, ",") {

			fields = strings.TrimSpace(fields)
			field := strings.Fields(fields)

			i := 0
			for ; i < len(field)-1; i += 2 {

				switch field[i] {
				case "ALIAS:":
					authrc.alias = field[i+1]
				case "SERVICE:":
					authrc.service = field[i+1]
				case "USERNAME:":
					authrc.username = field[i+1]
				case "PASSWORD:":
					authrc.password = field[i+1]
				case "COMMAND:":
					authrc.command = field[i+1]
				}
			}
		}
		database = append(database, authrc)
	}
	return
}

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/git.hansaray.pw/go/authrc?status.svg)](https://godoc.org/git.hansaray.pw/go/authrc)

# authrc - auth runcom
### Welcome to the friendly authentication runcom library and command


# authrc - runcom command

`authrc` is a command line tool to read .authrc files.

See 'bin' folder

# .authrc - runcom file

The .autrc file contains login and initialization information used for auto login. It generally resides in the user’s config directory, but a location outside of the config directory can be set.

- ALIAS<br>
    Alias of a service, or remote host

- SERVICE<br>
    Identify a service or remote host name. The auto login process searches the .authrc file for a SERVICE token that matches the service or remote machine specified on the command line argument. Once a match is made, the subsequent .authrc tokens are processed, stopping when the end of line is reached

- USERNAME<br>
    Identify a user of the service, or on the remote machine.

- PASSWORD<br>
    Supply a password. If this token is present, the auto-login process will supply the specified string if the service or remote server requires a password as part of the login process.

.authrc single line example:
```bash
ALIAS: blog, SERVICE: pasts.example.net, USERNAME: bloguser@example.net, PASSWORD: 12345abcd
```


[![License](https://img.shields.io/badge/license-MIT-blue)](./LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4)](https://pkg.go.dev/github.com/android-project-46group/sgi-cli)
[![Release](https://img.shields.io/github/release/android-project-46group/sgi-cli.svg?style=flat-square)](https://github.com/android-project-46group/sgi-cli/releases)

# sgi-cli

CLI tool for Sakamichi Group Information (SGI).

## INSTALLATION

### Get binary from releases

Built binaries are available from GitHub Releases.  
https://github.com/android-project-46group/sgi-cli/releases

**MacOS, Linux**

```sh
# CAUTION: this script adds binary to /usr/local/bin folder
$ curl -Lsf https://github.com/android-project-46group/sgi-cli/main/_tools/scripts/installer.sh | bash
```

#### Uninstall

**MacOS, Linux**

```sh
# delete the binary
$ rm /usr/local/bin/sgi

# delete the related files
$ rm -r ~/.sgi
```

### Build from source

1. Install [Go](https://go.dev/doc/install) version 1.18 or later
2. Update your go related environment variables as described in the Go documentation

> The install directory is controlled by the GOPATH and GOBIN
> environment variables. If GOBIN is set, binaries are installed
> to that directory. If GOPATH is set, binaries are installed to
> the bin subdirectory of the first directory in the GOPATH list.
> Otherwise, binaries are installed to the bin subdirectory of
> the default GOPATH ($HOME/go or %USERPROFILE%\go).

3. Then you can build from source

```sh
$ go install github.com/android-project-46group/sgi-cli@latest

# Rename the binary if you want.
$ mv $GOPATH/bin/sgi-cli $GOPATH/bin/sgi
```

#### Uninstall

Jst delete the binary.

```sh
$ rm $GOPATH/bin/sgi-cli

# delete the related files
$ rm -r ~/.sgi
```

### How to set account information

Put the following format json

```json
// ~/.sgi/account.json
{
  "baseURL": "https://uri/to/api_server",
  "apiKey": "your_api_key"
}
```

Or log-in with `login` subcommand,

```sh
$ sgi login
Please provide your login information.

 BaseURL:
 ...
```

## COMMANDS

### _member_

About members information.  
For now, only **_ls_** subcommand is available

```sh
$ sgi member ls -h
fetch all information

Usage:
  sgi member ls [flags]

Flags:
  -d, --data           print all data
  -g, --group string   group name to fetch
  -h, --help           help for ls
  -j, --json           print as a json format

# list nogizaka members
$ sgi member ls -g nogizaka
```

### _group_

About groups information.  
For now, only **_ls_** subcommand is available

```sh
$ sgi group ls -h
fetch all information

Usage:
  sgi group ls [flags]

Flags:
  -h, --help   help for ls
  -j, --json   print as a json format

# list nogizaka members
$ sgi group ls
```

## LICENSE

under [MIT License](./LICENSE).

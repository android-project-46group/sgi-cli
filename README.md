[![License](https://img.shields.io/badge/license-MIT-blue)](./LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4)](https://pkg.go.dev/github.com/android-project-46group/sgi-cli)
[![Release](https://img.shields.io/github/release/android-project-46group/sgi-cli.svg?style=flat-square)](https://github.com/android-project-46group/sgi-cli/releases)

# sgi-cli

CLI tool for Sakamichi Group Information (SGI).

## Installation

Built binaries are available from GitHub Releases.  
https://github.com/android-project-46group/sgi-cli/releases

**MacOS, Linux**

```sh
# CAUTION: this script adds binary to /usr/local/bin folder
$ curl -Lsf https://github.com/android-project-46group/sgi-cli/main/_tools/scripts/installer.sh | bash
```

### How to set account information

```json
// ~/.sgi/account.json
{
  "baseURL": "https://uri/to/api_server",
  "apiKey": "your_api_key"
}
```

### Uninstall

**MacOS, Linux**

```sh
# delete the binary
$ rm /usr/local/bin/sgi

# delete the related files
$ rm -r ~/.sgi
```

## COMMANDS

### _member_

About members information.  
For now, only **_ls_** subcommand is available

```sh
$ sgi member ls -h
fetch all information

Usage:
  sgi-cli member ls [flags]

Flags:
  -d, --data           print all data
  -g, --group string   group name to fetch
  -h, --help           help for ls
  -j, --json           print as a json format

# list nogizaka members
$ sgi member ls -g nogizaka
```

## LICENSE

under [MIT License](./LICENSE).

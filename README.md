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
## LICENSE

under [MIT License](./LICENSE).

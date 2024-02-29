# Grole :cloud:

![Release](https://img.shields.io/github/v/release/Cloud-Technology-Solutions/grole?style=social)
![Go Version](https://img.shields.io/github/go-mod/go-version/Cloud-Technology-Solutions/grole?style=plastic)
[![Pipeline Status](https://github.com/Cloud-Technology-Solutions/grole/actions/workflows/release.yaml/badge.svg)](https://github.com/Cloud-Technology-Solutions/grole/actions/workflows/release.yaml)

![Maintained](https://img.shields.io/maintenance/yes/2023)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

## Description :rocket:

Grole is a tool used for establishing which Google
Cloud Roles contain a given permission,
and also which permissions a given Role has.

If you have ever found yourself wondering:

- "Which google cloud permissions does a given role have?"
- "Which google cloud role has x permissions included?"

Then `grole` is the tool for you!

## Pre-Requisites :warning:

- [Gcloud](https://cloud.google.com/sdk/docs/install)
- Run `gcloud auth application-default login`

## Installing & Updating :package:

If you already have Go installed, you can simply install `grole` using:

```bash
go install github.com/jay-kinder/grole@latest
```

This will also allow you to update to the latest version.

If you don't have Go installed, you can install and update `grole` manually:

### Linux

```bash
GROLE_VERSION=v1.0.0 \

sudo rm -rf /usr/local/bin/grole \
wget -c https://github.com/jay-kinder/grole/releases/download/"${GROLE_VERSION}"/grole-"${GROLE_VERSION}"-linux-"$(dpkg --print-architecture)".tar.gz \
sudo tar -C /usr/local/bin -xzf grole-"${GROLE_VERSION}"-linux-"$(dpkg --print-architecture)".tar.gz \
rm -f grole-"${GROLE_VERSION}"-linux-"$(dpkg --print-architecture)".tar.gz
```

### Mac

```bash
GROLE_VERSION=v1.0.0 \

sudo rm -rf /usr/local/bin/grole \
wget -c https://github.com/jay-kinder/grole/releases/download/"${GROLE_VERSION}"/grole-"${GROLE_VERSION}"-"$(dpkg --print-architecture)".tar.gz \
sudo tar -C /usr/local/bin -xzf grole-"${GROLE_VERSION}"-"$(dpkg --print-architecture)".tar.gz \
rm -f grole-"${GROLE_VERSION}"-"$(dpkg --print-architecture)".tar.gz
```

### Windows

Go to the [Releases](https://github.com/Cloud-Technology-Solutions/grole/releases)
page and download the zip file you require.

Once unzipped, you will find a `.exe` you can run to install `grole`.

## Usage :white_check_mark:

```bash
-r, --role: Provide a role name to see all the permissions it has
-p, --perm: Provide permission(s) and see which role(s) contain both this permission(s) and the smallest number of other permissions (helps to follow the principle of least privilege)
```

```bash
grole [-p|--perm] <value> # optional: --all

# you can pass multiple permissions by using multiple -p flags
# grole will provide role(s) which contain all permissions provided

grole [-r|--role] <value>
```

### Example :eyes:

```bash
grole [-p|--perm] resourcemanager.organizations.get
grole [-p|--perm] resourcemanager.projects.get [-p|--perm] compute.vpnGateways.list

grole [-r|--role] compute.admin
```

### Additional Flags :checkered_flag:

```txt
-h, --help: Print help information
--all: Provide permission(s) with [-p | --perm ] and see all role(s) that contain this permission(s)
-v, --version: Get currently installed version of grole
```

## Issues :beetle:

If you find any bugs, or think of any issues you would like to raise, please raise
them [here](https://github.com/Cloud-Technology-Solutions/grole/issues).

It will be a big help if you could please try to choose the most appropriate label
when raising an issue for `grole`.

## Contributing

See the [`CONTRIBUTING.md`](CONTRIBUTING.md) for the project's contributing
guidelines.

## Code of Conduct

:hugs: Be nice.

## License :balance_scale:

Grole is released under the GNU GENERAL PUBLIC license.
See [`LICENSE`](LICENSE) for details.

## Uninstalling :wave:

To uninstall `grole`, you simply need to delete the `grole` directory from the install
location.

This is likely to either be `$(go env GOPATH)/bin` or `/usr/local/bin/grole`
for Linux/Mac, and your chosen install location for Windows.

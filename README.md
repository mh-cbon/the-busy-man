# the-busy-man

[![travis Status](https://travis-ci.org//mh-cbon/the-busy-man.svg?branch=master)](https://travis-ci.org//mh-cbon/the-busy-man) [![Appveyor Status](https://ci.appveyor.com/api/projects/status//github/mh-cbon/the-busy-man?branch=master&svg=true)](https://ci.appveyor.com/projects//mh-cbon/the-busy-man) [![Go Report Card](https://goreportcard.com/badge/github.com/mh-cbon/the-busy-man)](https://goreportcard.com/report/github.com/mh-cbon/the-busy-man) [![GoDoc](https://godoc.org/github.com/mh-cbon/the-busy-man?status.svg)](http://godoc.org/github.com/mh-cbon/the-busy-man) [![MIT License](http://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Package The busy man is a cli tool to initialize a project.


s/The busy man/[l'homme pressé](https://www.youtube.com/watch?v=Wkxe1kQiuGU/)/

# TOC
- [Install](#install)
  - [go](#go)
- [Usage](#usage)
  - [$ the-busy-man -help](#-the-busy-man--help)
  - [$ the-busy-man -l](#-the-busy-man--l)
  - [$ the-busy-man -h changelog](#-the-busy-man--h-changelog)
  - [$ the-busy-man -h emd](#-the-busy-man--h-emd)
  - [$ the-busy-man -h git](#-the-busy-man--h-git)
  - [$ the-busy-man -h glide](#-the-busy-man--h-glide)
  - [$ the-busy-man -h go](#-the-busy-man--h-go)
  - [$ the-busy-man -h gump](#-the-busy-man--h-gump)
  - [$ the-busy-man -h license](#-the-busy-man--h-license)
- [Cli examples](#cli-examples)
- [Your own plugin](#your-own-plugin)
- [Recipes](#recipes)
  - [Release the project](#release-the-project)
- [History](#history)

# Install

#### go
```sh
go get github.com/mh-cbon/the-busy-man
```

## Usage

#### $ the-busy-man -help
```sh
the-busy-man 0.0.0

Usage

	the-busy-man [-w directory] [plugins intents]

	  -w:              The directory to initialize.
	  plugins intents: A list of plugin wiht their intents
	                   such as plugin:intent1+intent2+'intent 3'.

Options
	-w:               The directory to initialize.
	-l:               List all plugins.
	-h|help [plugin]: Show help [of a plugin].
	-v|-version:      The directory to initialize.
```
#### $ the-busy-man -l
```sh
- changelog: Initialize a changelog file
- emd: Initialize a README emd file
- git: Initialize a git repository
- glide: Initialize a glide package
- go: Initialize a golang project
- golang: Initialize a golang project
- gump: Initialize a release script
- license: Initialize a license file
```
#### $ the-busy-man -h changelog
```sh
the-busy-man 0.0.0

	changelog: Initialize an changes.log file.
```
#### $ the-busy-man -h emd
```sh
the-busy-man 0.0.0

	emd:init: Intialize the README with the default template.
	emd:user/repo: Download a README.e.md file from the repo github.com/user/repo/README.e.md.
```
#### $ the-busy-man -h git
```sh
the-busy-man 0.0.0

	git: Run git init.
	git:commit: Run git add -A && git commit -am with a default message.
	git:amend: Add --amend to the commit.
```
#### $ the-busy-man -h glide
```sh
the-busy-man 0.0.0

	glide: Run glide init.
```
#### $ the-busy-man -h go
```sh
the-busy-man 0.0.0

	golang: Initialize a default main.go.
```
#### $ the-busy-man -h gump
```sh
the-busy-man 0.0.0

	gump: Initialize an empty .version.sh file.
	gump:user/repo: Initialize a .version.sh file downloaded from github.com/user/repo/.version.sh.
```
#### $ the-busy-man -h license
```sh
the-busy-man 0.0.0

	license:name: Initialize a LICENSE file matching name.
```

## Cli examples

```sh
the-busy-man git:init license:mit emd:mh-cbon/emd golang gump:mh-cbon/gump git:commit changelog git:commit+amend
```

I recommend you create an alias:
```sh
$ cat <<EOT >> ~/.bashrc
alias tbm="the-busy-man git:init license:mit emd:mh-cbon/emd golang gump:mh-cbon/gump git:commit changelog git:commit+amend"
EOT
source ~/.bashrc
```



# Your own plugin

For simplicity, and given current status of `go plugin`,
i propose you PR that repo with the plugin you d like to add.

A plugin is a struct that implements `plugin` interface:
```go
type plugin interface {
	Name() string
	Help()
	Description() string
	Handle(w wish.Wishes) error
}
```

# Recipes

#### Release the project

```sh
gump patch -d # check
gump patch # bump
```

# History

[CHANGELOG](CHANGELOG.md)

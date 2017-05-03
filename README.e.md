---
License: MIT
LicenseFile: LICENSE
LicenseColor: yellow
---
# {{.Name}}

{{template "badge/travis" .}} {{template "badge/appveyor" .}} {{template "badge/goreport" .}} {{template "badge/godoc" .}} {{template "license/shields" .}}

{{pkgdoc}}

s/The busy man/[l'homme pressé](https://www.youtube.com/watch?v=Wkxe1kQiuGU/)

# {{toc 5}}

# Install

#### go
{{template "go/install" .}}

## Usage

#### $ {{exec "the-busy-man" "-help" | color "sh"}}

#### $ {{exec "the-busy-man" "-l" | color "sh"}}

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

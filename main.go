// Package The busy man is a cli tool to initialize a project.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mh-cbon/the-busy-man/changelog"
	"github.com/mh-cbon/the-busy-man/emd"
	"github.com/mh-cbon/the-busy-man/git"
	"github.com/mh-cbon/the-busy-man/golang"
	"github.com/mh-cbon/the-busy-man/gump"
	"github.com/mh-cbon/the-busy-man/license"
	"github.com/mh-cbon/the-busy-man/utils"
	"github.com/mh-cbon/the-busy-man/wish"
)

var name = "the-busy-man"
var version = "0.0.0"

func main() {

	wd, err := os.Getwd()
	if err != nil {
		panic(err) //not so good.
	}

	var w string
	var l bool
	var h bool
	var help bool
	var v bool
	var ver bool
	flag.StringVar(&w, "w", wd, "Working directory.")
	flag.BoolVar(&l, "l", false, "Plugins list.")
	flag.BoolVar(&h, "h", false, "Show help.")
	flag.BoolVar(&help, "help", false, "Show help.")
	flag.BoolVar(&v, "v", false, "Show version.")
	flag.BoolVar(&ver, "version", false, "Show version.")
	flag.Parse()

	args := flag.Args()

	if ver || v {
		showVer()
		return
	}
	if help || h {
		if len(args) > 0 {
			showPluginHelp(args)
		} else {
			showHelp()
		}
		return
	}
	if l {
		showPlugins()
		return
	}

	wishes := wish.NewWishes()

	for _, arg := range args {
		w, err := wish.Parse(arg)
		if err != nil {
			panic(err)
		}
		wishes.Push(w)
	}
}

func showVer() {
	fmt.Printf("%v %v\n", name, version)
}

//go:generate lister -p utils utils/string_slice.go string:StringSlice

func showPluginHelp(p []string) {
	showVer()
	fmt.Println()
	x := utils.NewStringSlice().Push(p...)
	for _, plugin := range getPlugins() {
		if x.Index(plugin.Name()) > -1 {
			plugin.Help()
		}
	}
}

func showHelp() {
	showVer()
	fmt.Println()
	fmt.Println("Usage")
	fmt.Println()
	fmt.Printf("	%v [-w directory] [plugins intents]\n\n", name)
	fmt.Printf("	  -w:              The directory to initialize.\n")
	fmt.Printf("	  plugins intents: A list of plugin wiht their intents\n")
	fmt.Printf("	                   such as plugin:intent1+intent2+'intent 3'.\n")
	fmt.Println()
	fmt.Println("Options")
	fmt.Printf("	-w:               The directory to initialize.\n")
	fmt.Printf("	-l:               List all plugins.\n")
	fmt.Printf("	-h|help [plugin]: Show help [of a plugin].\n")
	fmt.Printf("	-v|-version:      The directory to initialize.\n")
	fmt.Println()
}

func showPlugins() {
	for _, p := range getPlugins() {
		fmt.Printf("- %v: %v\n", p.Name(), p.Description())
	}
}

type plugin interface {
	Name() string
	Help()
	Description() string
	Handle(w wish.Wishes) error
}

func getPlugins() map[string]plugin {
	ret := map[string]plugin{}
	ret["changelog"] = &changelog.Plugin{}
	ret["emd"] = &emd.Plugin{}
	ret["git"] = &git.Plugin{}
	ret["golang"] = &golang.Plugin{}
	ret["gump"] = &gump.Plugin{}
	ret["license"] = &license.Plugin{}
	return ret
}

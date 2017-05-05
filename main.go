// Package The busy man is a cli tool to initialize a project.
package main

//go:generate lister -p utils utils/string_slice.go string:StringSlice

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/mh-cbon/the-busy-man/changelog"
	"github.com/mh-cbon/the-busy-man/emd"
	"github.com/mh-cbon/the-busy-man/git"
	"github.com/mh-cbon/the-busy-man/glide"
	"github.com/mh-cbon/the-busy-man/golang"
	"github.com/mh-cbon/the-busy-man/gump"
	"github.com/mh-cbon/the-busy-man/license"
	"github.com/mh-cbon/the-busy-man/plugin"
	"github.com/mh-cbon/the-busy-man/wish"
)

var name = "the-busy-man"
var version = "0.0.0"

var verbose = false

func main() {

	verbose = os.Getenv("VERBOSE") != ""

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
	if len(args) == 0 && (help || h) {
		showHelp()
		return
	}

	common := &plugin.Plugin{}
	common.SetVerbose(os.Getenv("VERBOSE") != "")

	plugins := getPlugins(common)

	if len(args) > 0 && (help || h) {
		showPluginHelp(plugins, args)
		return
	}
	if l {
		showPlugins(plugins)
		return
	}

	if w != wd {
		os.Chdir(w)
	}

	common.Log("wd=%v", wd)
	common.Log("w=%v", w)
	common.SetOldWd(wd)

	wishes := parseWishes(args)
	pluginsHandle(common, plugins, wishes)

}

func parseWishes(args []string) *wish.Wishes {
	wishes := wish.NewWishes()
	for _, arg := range args {
		w, err := wish.Parse(arg)
		if err != nil {
			panic(err)
		}
		wishes.Push(w)
	}
	return wishes
}

func pluginsHandle(common *plugin.Plugin, plugins map[string]pluginHandler, wishes *wish.Wishes) {
	for _, w := range wishes.Get() {
		if p, ok := plugins[w.Plugin]; ok {
			common.Log("handle %v...", w.Plugin)
			err := p.Handle(wishes, w)
			if err != nil {
				panic(err)
			}
		}
	}
}

func showVer() {
	fmt.Printf("%v %v\n", name, version)
}

func showPluginHelp(plugins map[string]pluginHandler, p []string) {
	showVer()
	fmt.Println()
	for _, name := range p {
		if x, ok := plugins[name]; ok {
			x.Help()
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

func showPlugins(plugins map[string]pluginHandler) {
	x := []string{}
	for name := range plugins {
		x = append(x, name)
	}
	sort.Strings(x)
	for _, name := range x {
		fmt.Printf("- %v: %v\n", name, plugins[name].Description())
	}
}

type pluginHandler interface {
	Name() string
	Help()
	Description() string
	Handle(*wish.Wishes, *wish.Wish) error
}

func getPlugins(common *plugin.Plugin) map[string]pluginHandler {
	ret := map[string]pluginHandler{}
	ret["changelog"] = &changelog.Plugin{Plugin: common}
	ret["emd"] = &emd.Plugin{Plugin: common}
	ret["git"] = &git.Plugin{Plugin: common}
	ret["golang"] = &golang.Plugin{Plugin: common}
	ret["go"] = &golang.Plugin{Plugin: common}
	ret["glide"] = &glide.Plugin{Plugin: common}
	ret["gump"] = &gump.Plugin{Plugin: common}
	ret["license"] = &license.Plugin{Plugin: common}
	return ret
}

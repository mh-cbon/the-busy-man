package golang

import (
	"fmt"

	"github.com/mh-cbon/the-busy-man/plugin"
	"github.com/mh-cbon/the-busy-man/wish"
)

// Plugin golang for the busy man.
type Plugin struct {
	*plugin.Plugin
}

// Name of the plugin
func (p *Plugin) Name() string {
	return "golang"
}

// Description of the plugin
func (p *Plugin) Description() string {
	return "Initialize a golang project"
}

// Help of the plugin
func (p *Plugin) Help() {
	fmt.Println("	golang: Initialize a default main.go.")
}

// Handle wishes of the busy man.
func (p *Plugin) Handle(w *wish.Wishes, plugin *wish.Wish) error {
	data := `//Package...
package xx

func main(){

}`
	return p.Write("main.go", data)
}

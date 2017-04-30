package golang

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mh-cbon/the-busy-man/wish"
)

// Plugin golang for the busy man.
type Plugin struct {
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
func (p *Plugin) Handle(w wish.Wishes) error {
	x := w.Filter(wish.FilterByPlugin("go"))
	if x.Len() > 0 {
		data := `//Package...
package xx

func main(){

}`
		return ioutil.WriteFile("main.go", []byte(data), os.ModePerm)
	}
	return nil
}

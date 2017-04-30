package emd

import (
	"fmt"
	"strings"

	"github.com/mh-cbon/the-busy-man/plugin"
	"github.com/mh-cbon/the-busy-man/wish"
)

// Plugin emd for the busy man.
type Plugin struct {
	*plugin.Plugin
}

// Name of the plugin
func (p *Plugin) Name() string {
	return "emd"
}

// Description of the plugin
func (p *Plugin) Description() string {
	return "Initialize a README emd file"
}

// Help of the plugin
func (p *Plugin) Help() {
	fmt.Println("	emd:init: Intialize the README with the default template.")
	fmt.Println("	emd:user/repo: Download a README.e.md file from the repo github.com/user/repo/README.e.md.")
}

// Handle wishes of the busy man.
func (p *Plugin) Handle(w *wish.Wishes) error {
	x := w.Filter(wish.FilterByPlugin("emd"))
	if x.Len() > 0 {
		plugin := x.At(0)
		err := p.Exec("emd", "-version")
		if err != nil {
			err = p.GoGet("github.com/mh-cbon/emd")
			if err != nil {
				return err
			}
		}
		if plugin.Shades.Len() > 0 {
			x := plugin.Shades.At(0)
			if strings.Index(x, "/") > -1 {
				return p.DlGhRawFile("README.e.md", x)
			}
			return p.Exec("emd", "init")
		}
	}
	return nil
}

package changelog

import (
	"fmt"

	"github.com/mh-cbon/the-busy-man/plugin"
	"github.com/mh-cbon/the-busy-man/wish"
)

// Plugin changelog for the busy man.
type Plugin struct {
	*plugin.Plugin
}

// Name of the plugin
func (p *Plugin) Name() string {
	return "changelog"
}

// Description of the plugin
func (p *Plugin) Description() string {
	return "Initialize a changelog file"
}

// Help of the plugin
func (p *Plugin) Help() {
	fmt.Println("	changelog: Initialize an changes.log file.")
}

// Handle wishes of the busy man.
func (p *Plugin) Handle(w *wish.Wishes) error {
	x := w.Filter(wish.FilterByPlugin("changelog"))
	if x.Len() > 0 {
		// plugin := x.At(0)
		err := p.Exec("changelog", "-version")
		if err != nil {
			err = p.GoGet("github.com/mh-cbon/changelog")
			if err != nil {
				return err
			}
		}
		return p.Exec("changelog", "init")
	}
	return nil
}

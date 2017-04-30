package git

import (
	"fmt"

	"github.com/mh-cbon/the-busy-man/plugin"
	"github.com/mh-cbon/the-busy-man/wish"
)

// Plugin git for the busy man.
type Plugin struct {
	*plugin.Plugin
}

// Name of the plugin
func (p *Plugin) Name() string {
	return "glide"
}

// Description of the plugin
func (p *Plugin) Description() string {
	return "Initialize a glide package"
}

// Help of the plugin
func (p *Plugin) Help() {
	fmt.Println("	glide: Run glide init.")
}

// Handle wishes of the busy man.
func (p *Plugin) Handle(w *wish.Wishes) error {
	x := w.Filter(wish.FilterByPlugin("glide"))
	if x.Len() > 0 {
		err := p.Exec("glide", "-version")
		if err != nil {
			err = p.GoGet("github.com/mh-cbon/emd")
			if err != nil {
				return err
			}
		}
		err = p.Exec("glide", "init")
		if err != nil {
			return err
		}
	}
	return nil
}

package license

import (
	"fmt"

	"github.com/mh-cbon/the-busy-man/plugin"
	"github.com/mh-cbon/the-busy-man/wish"
)

// Plugin license for the busy man.
type Plugin struct {
	*plugin.Plugin
}

// Name of the plugin
func (p *Plugin) Name() string {
	return "license"
}

// Description of the plugin
func (p *Plugin) Description() string {
	return "Initialize a license file"
}

// Help of the plugin
func (p *Plugin) Help() {
	fmt.Println("	license:name: Initialize a LICENSE file matching name.")
}

// Handle wishes of the busy man.
func (p *Plugin) Handle(w *wish.Wishes, plugin *wish.Wish) error {
	err := p.Exec("license", "-version")
	if err != nil {
		err = p.GoGet("github.com/nishanths/license")
		if err != nil {
			return err
		}
	}
	if plugin.Shades.Len() > 0 {
		return p.Exec("license", "-o", "LICENSE", plugin.Shades.At(0))
	}
	p.Warn("missing license name in 'license:?' intent")
	return nil
}

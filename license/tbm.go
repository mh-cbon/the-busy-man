package license

import (
	"fmt"
	"os/exec"

	"github.com/mh-cbon/the-busy-man/wish"
)

// Plugin license for the busy man.
type Plugin struct {
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
func (p *Plugin) Handle(w wish.Wishes) error {
	x := w.Filter(wish.FilterByPlugin("license"))
	if x.Len() > 0 {
		plugin := x.At(0)
		err := exec.Command("license", "-version").Run()
		if err != nil {
			err = exec.Command("go", "get", "-u", "github.com/nishanths/license").Run()
			if err != nil {
				return err
			}
		}
		if plugin.Shades.Len() > 0 {
			return exec.Command("license", "-o", "LICENSE", plugin.Shades.At(0)).Run()
		}
	}
	return nil
}

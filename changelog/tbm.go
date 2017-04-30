package changelog

import (
	"fmt"
	"os/exec"

	"github.com/mh-cbon/the-busy-man/wish"
)

// Plugin changelog for the busy man.
type Plugin struct {
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
		err := exec.Command("changelog", "-version").Run()
		if err != nil {
			w.Log("changelog not found, installing...")
			err = exec.Command("go", "get", "-u", "github.com/mh-cbon/changelog").Run()
			if err != nil {
				return err
			}
		}
		w.Log("changelog init")
		return exec.Command("changelog", "init").Run()
	}
	return nil
}

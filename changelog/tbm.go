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
func (p *Plugin) Handle(w *wish.Wishes, plugin *wish.Wish) error {
	if err := p.goGet(w, plugin); err != nil {
		return err
	}
	if err := p.init(w, plugin); err != nil {
		return err
	}
	return nil
}

func (p *Plugin) init(w *wish.Wishes, plugin *wish.Wish) error {
	p.Print("> init changelog...")
	return p.Exec("changelog", "init")
}

func (p *Plugin) goGet(w *wish.Wishes, plugin *wish.Wish) error {
	p.Print("? checking changelog...")
	err := p.Exec("changelog", "-version")
	if err != nil {
		p.Print("? installing changelog...")
		p.GoGet("github.com/Masterminds/glide")
		err = p.GoGet("github.com/mh-cbon/changelog")
		if err != nil {
			return err
		}
		err = p.GlideInstall("github.com/mh-cbon/changelog")
		if err != nil {
			return err
		}
	}
	p.Print("✓ changelog is up!")
	return nil
}

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
	p.Print("> init emd...")
	if plugin.Shades.Len() > 0 {
		x := plugin.Shades.First()
		p.Log("x=%v", x)
		if strings.Index(x, "/") > -1 {
			return p.DlGhRawFile("README.e.md", x)
		}
	}
	return p.Exec("emd", "init")
}

func (p *Plugin) goGet(w *wish.Wishes, plugin *wish.Wish) error {
	p.Print("? checking emd...")
	err := p.Exec("emd", "-version")
	if err != nil {
		p.Print("? installing emd...")
		p.GoGet("github.com/Masterminds/glide")
		err = p.GoGet("github.com/mh-cbon/emd")
		if err != nil {
			return err
		}
		err = p.GlideInstall("github.com/mh-cbon/emd")
		if err != nil {
			return err
		}
	}
	p.Print("✓ emd is up!")
	return nil
}

package gump

import (
	"fmt"
	"strings"

	"github.com/mh-cbon/the-busy-man/plugin"
	"github.com/mh-cbon/the-busy-man/wish"
)

// Plugin gump for the busy man.
type Plugin struct {
	*plugin.Plugin
}

// Name of the plugin
func (p *Plugin) Name() string {
	return "gump"
}

// Description of the plugin
func (p *Plugin) Description() string {
	return "Initialize a release script"
}

// Help of the plugin
func (p *Plugin) Help() {
	fmt.Println("	gump: Initialize an empty .version.sh file.")
	fmt.Println("	gump:user/repo: Initialize a .version.sh file downloaded from github.com/user/repo/.version.sh.")
}

// Handle wishes of the busy man.
func (p *Plugin) Handle(w *wish.Wishes, plugin *wish.Wish) error {
	err := p.Exec("gump", "-version")
	if err != nil {
		p.GoGet("github.com/Masterminds/glide")
		err = p.GoGet("github.com/mh-cbon/gump")
		if err != nil {
			return err
		}
		err = p.GlideInstall("github.com/mh-cbon/gump")
		if err != nil {
			return err
		}
	}
	if plugin.Shades.Len() > 0 {
		x := plugin.Shades.At(0)
		if strings.Index(x, "/") > -1 {
			return p.DlGhRawFile(".version.sh", x)
		}
		data := `PREBUMP=

PREVERSION=

POSTVERSION=
`
		return p.Write(".version.sh", data)
	}
	return nil
}

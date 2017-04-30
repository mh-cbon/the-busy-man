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
	return "git"
}

// Description of the plugin
func (p *Plugin) Description() string {
	return "Initialize a git repository"
}

// Help of the plugin
func (p *Plugin) Help() {
	fmt.Println("	git: Run git init.")
	fmt.Println("	git:commit: Run git add -A && git commit -am with a default message.")
}

// Handle wishes of the busy man.
func (p *Plugin) Handle(w *wish.Wishes) error {
	x := w.Filter(wish.FilterByPlugin("git"))
	if x.Len() > 0 {
		plugin := x.At(0)
		data := ``
		if w.Filter(wish.FilterByPlugin("glide", "dep")).Len() > 0 {
			p.Log("adding vendor to gitignore...")
			data += "vendor/\n"
		}
		err := p.Write(".gitignore", data)
		if err != nil {
			return err
		}
		if plugin.Shades.Len() == 0 || plugin.HasShade("init") {
			err = p.Exec("git", "init")
			if err != nil {
				return err
			}
		}
		if plugin.HasShade("commit") {
			err = p.Exec("git", "add", "-A")
			if err != nil {
				return err
			}
			err = p.Exec("git", "commit", "-am", "Project initialization")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

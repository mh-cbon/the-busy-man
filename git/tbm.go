package git

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/mh-cbon/the-busy-man/wish"
)

// Plugin git for the busy man.
type Plugin struct {
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
func (p *Plugin) Handle(w wish.Wishes) error {
	x := w.Filter(wish.FilterByPlugin("git"))
	if x.Len() > 0 {
		plugin := x.At(0)
		data := ``
		if w.Filter(wish.FilterByPlugin("glide", "dep")).Len() > 0 {
			data += "vendor/\n"
		}
		err := ioutil.WriteFile(".gitignore", []byte(data), os.ModePerm)
		if err != nil {
			return err
		}
		if plugin.Shades.Len() == 0 || plugin.Shades.Index("init") > -1 {
			err = exec.Command("git", "init").Run()
			if err != nil {
				return err
			}
		}
		if plugin.Shades.Index("commit") > -1 {
			err = exec.Command("git", "add", "-A").Run()
			if err != nil {
				return err
			}
			err = exec.Command("git", "commit", "-am", "Project initialization").Run()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

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
	fmt.Println("	git:amend: Add --amend to the commit.")
}

// Handle wishes of the busy man.
func (p *Plugin) Handle(w *wish.Wishes, plugin *wish.Wish) error {
	if err := p.gitIgnore(w, plugin); err != nil {
		return err
	}
	if err := p.gitInit(w, plugin); err != nil {
		return err
	}
	if err := p.gitCommit(w, plugin); err != nil {
		return err
	}
	return nil
}

func (p *Plugin) gitIgnore(w *wish.Wishes, plugin *wish.Wish) error {
	if p.FileExists(".gitignore") && !plugin.HasShade("force") {
		fmt.Println("skipping .gitignore, it already exists")
		fmt.Println("use git:init+force to proceed")
	}
	return p.FileAppendOnce(".gitignore", "vendor/")
}

func (p *Plugin) gitInit(w *wish.Wishes, plugin *wish.Wish) error {
	if plugin.Shades.Len() == 0 || plugin.HasShade("init") {
		if p.DirExists(".git") && !plugin.HasShade("force") {
			fmt.Println("skipping git init, it already exists")
			fmt.Println("use git:init+force to proceed")
		}
		return p.Exec("git", "init")
	}
	return nil
}

func (p *Plugin) gitCommit(w *wish.Wishes, plugin *wish.Wish) error {
	if plugin.HasShade("commit") {
		err := p.Exec("git", "add", "-A")
		if err != nil && !plugin.HasShade("force") {
			p.Print("  git add failed")
			p.Print("  use git:add+force to proceed")
			return err
		}
		p.FYI(err, "git add failed")
		args := []string{"commit", "-am", "Project initialization"}
		if plugin.HasShade("amend") {
			args = append(args, "--amend")
		}
		err = p.Exec("git", args...)
		if err != nil && !plugin.HasShade("force") {
			p.Print("  git commit failed")
			p.Print("  use git:commit+force to proceed")
			return err
		}
		p.FYI(err, "git commit failed")
	}
	return nil
}

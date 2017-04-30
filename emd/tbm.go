package emd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/mh-cbon/the-busy-man/wish"
)

// Plugin emd for the busy man.
type Plugin struct {
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
func (p *Plugin) Handle(w *wish.Wishes) error {
	x := w.Filter(wish.FilterByPlugin("emd"))
	if x.Len() > 0 {
		plugin := x.At(0)
		err := exec.Command("emd", "-version").Run()
		if err != nil {
			w.Log("emd not found, installing...")
			err = exec.Command("go", "get", "-u", "github.com/mh-cbon/emd").Run()
			if err != nil {
				return err
			}
		}
		if plugin.Shades.Len() > 0 {
			x := plugin.Shades.At(0)
			if strings.Index(x, "/") > -1 {
				w.Log("emd downloading from %v", x)
				response, err := http.Get("http://github.com/" + x + "/master/README.e.md")
				if err != nil {
					return err
				}
				defer response.Body.Close()
				f, err := os.Create("README.e.md")
				if err != nil {
					return err
				}
				_, err = io.Copy(f, response.Body)
				if err != nil {
					return err
				}
			} else {
				w.Log("emd init...")
				return exec.Command("emd", "init").Run()

			}
		}
	}
	return nil
}

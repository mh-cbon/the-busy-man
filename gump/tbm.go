package gump

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/mh-cbon/the-busy-man/wish"
)

// Plugin gump for the busy man.
type Plugin struct {
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
func (p *Plugin) Handle(w *wish.Wishes) error {
	x := w.Filter(wish.FilterByPlugin("gump"))
	if x.Len() > 0 {
		plugin := x.At(0)
		err := exec.Command("gump", "-version").Run()
		if err != nil {
			w.Log("gump not found, installing...")
			err = exec.Command("go", "get", "-u", "github.com/mh-cbon/gump").Run()
			if err != nil {
				return err
			}
		}
		if plugin.Shades.Len() > 0 {
			x := plugin.Shades.At(0)
			if strings.Index(x, "/") > -1 {
				w.Log("gump downloading from %v", x)
				response, err := http.Get("http://github.com/" + x + "/master/.version.sh")
				if err != nil {
					return err
				}
				defer response.Body.Close()
				f, err := os.Create(".version.sh")
				if err != nil {
					return err
				}
				_, err = io.Copy(f, response.Body)
				if err != nil {
					return err
				}
			} else {
				w.Log("gump init...")
				data := `PREBUMP=

PREVERSION=

POSTVERSION=
`
				return ioutil.WriteFile(".version.sh", []byte(data), os.ModePerm)

			}
		}
	}
	return nil
}

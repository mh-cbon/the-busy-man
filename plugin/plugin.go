package plugin

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

// Plugin provides common func to plugins.
type Plugin struct {
	oldpwd  string
	verbose bool
}

// SetVerbose to enable logging.
func (p *Plugin) SetVerbose(s bool) {
	p.verbose = s
}

// SetOldWd saves oldpwd.
func (p *Plugin) SetOldWd(oldpwd string) {
	p.oldpwd = oldpwd
}

// Log message if verbose = true
func (p *Plugin) Log(format string, c ...interface{}) {
	if p.verbose {
		fmt.Printf(format+"\n", c...)
	}
}

// Warn message
func (p *Plugin) Warn(format string, c ...interface{}) {
	fmt.Printf("!! "+format+"\n", c...)
}

// Write a file
func (p *Plugin) Write(file string, data string) error {
	p.Log("writes %v...", file)
	return ioutil.WriteFile(file, []byte(data), os.ModePerm)
}

// Exec a command.
func (p *Plugin) Exec(c string, x ...string) error {
	p.Log("exec %v %v...", c, x)
	cmd := exec.Command(c, x...)
	if p.verbose {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// GoGet a package.
func (p *Plugin) GoGet(c string) error {
	return p.Exec("go", "get", "-u", c)
}

// GlideInstall a package.
func (p *Plugin) GlideInstall(c string) error {
	p.Log("exec glide install %v", c)
	cmd := exec.Command("glide", "install")
	cmd.Dir = filepath.Join(os.Getenv("GOPATH"), "src", c)
	if p.verbose {
		cmd.Stdout = os.Stdout
	}
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Dl download a file.
func (p *Plugin) Dl(dest string, u string) error {
	p.Log("emd downloading from %v to %v", u, dest)
	response, err := http.Get(u)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, response.Body)

	return err
}

// DlGhRawFile download a file.
func (p *Plugin) DlGhRawFile(dest string, r string) error {
	u := "https://raw.githubusercontent.com/" + r + "/master/" + dest
	return p.Dl(dest, u)
}

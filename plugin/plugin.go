package plugin

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

// Print message
func (p *Plugin) Print(format string, c ...interface{}) {
	fmt.Printf(format+"\n", c...)
}

// Warn message
func (p *Plugin) Warn(format string, c ...interface{}) {
	fmt.Printf("!! "+format+"\n", c...)
}

// FYI prints a message for your information is err is not nil
func (p *Plugin) FYI(err error, format string, c ...interface{}) {
	if err != nil {
		fmt.Printf("FYI: "+format+"\n", c...)
	}
}

// Write a file
func (p *Plugin) Write(file string, data string) error {
	p.Log("> writes %v...", file)
	return ioutil.WriteFile(file, []byte(data), os.ModePerm)
}

// Read a file
func (p *Plugin) Read(file string) (string, error) {
	p.Log("> reads %v...", file)
	if p.FileExists(file) == false {
		return "", nil
	}
	b, err := ioutil.ReadFile(file)
	if err == nil {
		return string(b), nil
	}
	return "", err
}

// FileExists yes/no
func (p *Plugin) FileExists(file string) bool {
	p.Log("> file exists %v...", file)
	stat, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return stat.IsDir() == false
}

// DirExists yes/no
func (p *Plugin) DirExists(file string) bool {
	p.Log("> dir exists %v...", file)
	stat, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return stat.IsDir() == true
}

// Exec a command.
func (p *Plugin) Exec(c string, x ...string) error {
	p.Print("$ %v %v...", c, strings.Join(x, " "))
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
	p.Print("$ glide install %v", c)
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
	p.Print("> download %v to %v", u, dest)
	response, err := http.Get(u)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	n, err := io.Copy(f, response.Body)
	p.Print("  downloaded %v bytes", n)

	return err
}

// DlGhRawFile download a file.
func (p *Plugin) DlGhRawFile(dest string, r string) error {
	u := "https://raw.githubusercontent.com/" + r + "/master/" + dest
	return p.Dl(dest, u)
}

// FileAppendOnce ensure given line exists in the file at least once or appends it.
func (p *Plugin) FileAppendOnce(file, line string) error {
	data, err := p.Read(file)
	if err != nil {
		return err
	}
	if strings.Contains(data, line) == false {
		return p.Write(file, data+"\n"+line)
	}
	return nil
}

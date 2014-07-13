package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

var homeDir = os.Getenv("HOME")

func main() {
	app := cli.NewApp()
	app.Name = "goshim"
	app.Usage = "Installs or uninstalls go shims. "
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:   "install",
			Usage:  "installs go-related shims into $HOME/.goshim/bin",
			Action: install,
		},
		{
			Name:   "uninstall",
			Usage:  "uninstalls all go-related shims from $HOME/.goshim/bin",
			Action: uninstall,
		},
	}
	app.Run(os.Args)
}

func execute(verbose bool, cmd string, args ...string) string {
	if verbose {
		outs := make([]interface{}, len(args)+1)
		outs[0] = "[goshim]$ " + cmd
		for i, arg := range args {
			outs[i+1] = arg
		}
		fmt.Println(outs...)
	}

	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	out, err := c.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return string(out)
}

// Shortcut for fmt.Sprintf
func s(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func install(c *cli.Context) {
	if _, err := os.Stat(s("%s/.goshim/bin/go", homeDir)); err == nil {
		fmt.Println("[goshim] Shims already installed.")
		os.Exit(0)
	}

	fmt.Println("[goshim] Fetching shims:")
	execute(true, "go", "get", "-d", "github.com/kevin-cantwell/goshim/go")
	fmt.Println("[goshim] Installing shims:")
	execute(true, "go", "build", "-o", s("%s/.goshim/bin/go", homeDir), "github.com/kevin-cantwell/goshim/go")

	fmt.Println("[goshim] Install complete.")
	fmt.Println("[goshim]   ** NOTE **")
	fmt.Println("[goshim]   To use goshim, you must insert the shims dir into your PATH:")
	fmt.Println("[goshim]   Eg: export PATH=$HOME/.goshim/bin:$PATH")
}

func uninstall(c *cli.Context) {
	if _, err := os.Stat(s("%s/.goshim", homeDir)); os.IsNotExist(err) {
		fmt.Println("[goshim] No shims detected.")
		os.Exit(0)
	}
	fmt.Print("[goshim] Unstalling shims...")
	if err := os.RemoveAll(s("%s/.goshim", homeDir)); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println("done")
	fmt.Println("[goshim]   ** NOTE **")
	fmt.Println("[goshim]   To completely uninstall goshim, you must remove $HOME/.goshim/bin from your PATH.")
}

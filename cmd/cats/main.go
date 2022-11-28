package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kong"
)

var CLI struct {
	Path  string `arg:"" required:""`
	Srefl string `name="srefl" enum:"redirect,extension,none" default:"none"`
}

func main() {
	kong.Parse(&CLI)

	if CLI.Srefl == "extension" {
		if strings.HasSuffix(CLI.Path, ":") {
			fmt.Printf(".s.txt")
			return
		}
		ext := filepath.Ext(CLI.Path)
		if ext != "" {
			ext2 := filepath.Ext(CLI.Path[:len(CLI.Path)-len(ext)])
			if ext2 != "" && len(ext2) < 5 {
				ext = ext2 + ext
			}
		}
		if ext == "." {
			return
		}
		fmt.Printf("%s", ext)
		return
	} else if CLI.Srefl == "redirect" {
		tp, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		if len(tp) == 0 {
			fmt.Printf("cats %s", CLI.Path)
			return
		}
		hostsep := strings.Index(CLI.Path, ":")
		host := ""

		if hostsep != -1 {
			host = CLI.Path[:hostsep]
			CLI.Path = CLI.Path[hostsep+1:]
		}

		tgtd := filepath.Dir(CLI.Path)
		tgt := filepath.Join(tgtd, string(tp))
		if host != "" {
			tgt = fmt.Sprintf("%s:%s", host, tgt)
		}
		c := "cats"
		if filepath.Ext(tgt) == "" {
			c = "sls"
		}
		fmt.Printf("%s %s", c, tgt)
		return
	}

	args := []string{"cat"}
	host := ""
	port := "22"

	if CLI.Path != "" {
		hostsep := strings.Index(CLI.Path, ":")
		if hostsep != -1 {
			host = CLI.Path[:hostsep]
			CLI.Path = CLI.Path[hostsep+1:]
			portsep := strings.Index(host, "^")
			if portsep != -1 {
				port = host[portsep+1:]
				host = host[:portsep]
			}
		}
		args = append(args, CLI.Path)
	}

	if host != "" {
		sshcmd := ""
		for _, c := range args {
			if c == "" {
				continue
			}
			if sshcmd != "" {
				sshcmd += " "
			}

			c = fmt.Sprintf(`"%s"`, c)
			c = strings.Replace(c, `\`, `\\`, -1)

			sshcmd += c
		}

		if sshcmd == `"cat"` {
			args = []string{"ssh", "-p", port, host}
		} else {
			args = append([]string{"ssh", "-p", port, host}, sshcmd)
		}
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			os.Exit(ee.ExitCode())
		}
		panic(err)
	}
}

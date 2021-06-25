package gosh_wget

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"git.mrcyjanek.net/mrcyjanek/gosh/_core/h"
)

var flags map[string]string

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	flags, cmd = h.ParseFlags(cmd)
	var url []string
	for i := range cmd {
		if cmd[i] == "wget" {
			continue
		}
		unused := true
		for j := range flags {
			if cmd[i] == flags[j] || cmd[i] == "-"+j || cmd[i] == "--"+j {
				unused = false
			}
		}
		if unused {
			url = append(url, cmd[i])
		}
	}
	if len(url) == 0 {
		STDOUT.WriteString(`wget: missing URL
Usage: wget [OPTION]... [URL]...

Try 'wget --help' for more options.
`)
		return 1
	}
	var output *bufio.Writer
	outputset := false
	if h.GetFlagBool("O", flags) || h.GetFlagBool("output", flags) {
		if h.GetFlagString("O", flags) == "true" || h.GetFlagString("output", flags) == "true" {
			output = STDOUT
			outputset = true
		} else if h.GetFlagBool("O", flags) {
			f, err := os.Open(h.GetFlagString("O", flags))
			if err != nil {
				STDERR.WriteString(err.Error() + "\n")
				return 1
			}
			output = bufio.NewWriter(f)
			outputset = true
		} else if h.GetFlagBool("output", flags) {
			f, err := os.Create(h.GetFlagString("output", flags))
			if err != nil {
				STDERR.WriteString(err.Error() + "\n")
				return 1
			}
			output = bufio.NewWriter(f)
			outputset = true
		}
	}
	for i := range url {
		resp, err := http.Get(url[i])
		if err != nil {
			STDERR.WriteString(err.Error() + "\n")
			return 1
		}
		if !outputset {
			f, err := os.Create(filepath.Base(resp.Request.URL.Path))
			if err != nil {
				STDERR.WriteString(err.Error() + "\n")
				return 1
			}
			output = bufio.NewWriter(f)
		}
		STDERR.WriteString("Downloading '" + url[i] + "'\n")
		io.Copy(output, resp.Body)
	}
	return 0
}

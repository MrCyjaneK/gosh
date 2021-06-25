package gosh_man

import (
	"bufio"
	"embed"

	"git.mrcyjanek.net/mrcyjanek/gosh/_core/h"
)

//go:embed manpages
var mans embed.FS

//go:embed default_fallback_info
var default_fallback_info string

var flags map[string]string

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	flags = h.ParseFlags(cmd)
	manpage := "man"
	if h.GetFlagBool("list", flags) {
		files, err := mans.ReadDir("manpages")
		if err != nil {
			STDERR.WriteString(err.Error() + "\n")
			return 2
		}
		STDOUT.WriteString("List of all available manual pages:\n")
		for i := range files {
			STDOUT.WriteString("\t- " + files[i].Name() + "\n")
		}
		return 0
	} else if len(cmd) > 1 {
		manpage = cmd[1]
	}
	b, err := getManPage(manpage, STDERR)
	if err != nil {
		STDERR.WriteString(err.Error() + "\n")
		STDERR.WriteString(default_fallback_info + "\n")
		return 1
	}
	STDOUT.Write(b)
	return 0
}

func getManPage(page string, STDERR *bufio.Writer) ([]byte, error) {
	b, err := mans.ReadFile("manpages/" + page)
	if err != nil {
		return []byte{}, err
	}
	return b, nil
}

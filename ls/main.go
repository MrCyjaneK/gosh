package gosh_ls

import (
	"bufio"
	"io/ioutil"
	"strings"

	"git.mrcyjanek.net/mrcyjanek/gosh/_core/h"
)

var flags map[string]string

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	files, err := ioutil.ReadDir(CWD)
	if err != nil {
		STDERR.WriteString(err.Error() + "\n")
		return 1
	}
	flags = h.ParseFlags(cmd)
	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, ".") && !(h.GetFlagBool("a", flags) || h.GetFlagBool("all", flags)) {
			continue
		}
		STDOUT.WriteString(name + "\n")
	}
	return 0
}

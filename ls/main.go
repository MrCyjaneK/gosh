package gosh_ls

import (
	"bufio"
	"io/ioutil"
	"strings"

	"git.mrcyjanek.net/mrcyjanek/gosh/_core/h"
)

var flags map[string]string

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	flags, cmd = h.ParseFlags(cmd)
	toret := uint8(0)

	if len(cmd) == 0 {
		files, err := ioutil.ReadDir(CWD)
		if err != nil {
			STDERR.WriteString(err.Error() + "\n")
			return 1
		}
		for _, file := range files {
			name := file.Name()
			if strings.HasPrefix(name, ".") && !(h.GetFlagBool("a", flags) || h.GetFlagBool("all", flags)) {
				continue
			}
			STDOUT.WriteString(name + "\n")
		}
	} else {
		for i := 1; i < len(cmd); i++ {
			files, err := ioutil.ReadDir(CWD + "/" + cmd[i])
			if err != nil {
				STDERR.WriteString(err.Error() + "\n")
				toret = 2
				continue
			}
			if i != 1 {
				STDOUT.WriteString("\n")
			}
			STDOUT.WriteString(cmd[i] + ":\n")
			for _, file := range files {
				name := file.Name()
				if strings.HasPrefix(name, ".") && !(h.GetFlagBool("a", flags) || h.GetFlagBool("all", flags)) {
					continue
				}
				STDOUT.WriteString(name + "\n")
			}
		}
	}
	return toret
}

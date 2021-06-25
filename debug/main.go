package gosh_echo

import (
	"bufio"
	"strconv"

	"git.mrcyjanek.net/mrcyjanek/gosh/_core/h"
)

var flags map[string]string

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	flags, cmd = h.ParseFlags(cmd)

	STDOUT.WriteString("CMD:\n")
	for i := range cmd {
		STDOUT.WriteString(" " + strconv.Itoa(i) + ". '" + cmd[i] + "'\n")
	}
	STDOUT.WriteString("FLAGS:\n")
	for flag := range flags {
		STDOUT.WriteString(" " + flag + ". '" + flags[flag] + "'\n")
	}
	return 0
}

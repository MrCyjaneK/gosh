package gosh_echo

import (
	"bufio"
	"strconv"
)

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	for i := range cmd {
		STDOUT.WriteString(" " + strconv.Itoa(i) + ". '" + cmd[i] + "'\n")
	}
	return 0
}

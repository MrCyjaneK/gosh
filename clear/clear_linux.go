// +build linux

package gosh_clear

import (
	"bufio"
)

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	STDOUT.Write([]byte("\033[2J\033[H"))
	return 0
}

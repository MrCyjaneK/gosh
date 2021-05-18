package gosh_printenv

import (
	"bufio"
)

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	for key := range ENV {
		STDOUT.WriteString(key + "=" + ENV[key] + "\n")
	}
	return 0
}

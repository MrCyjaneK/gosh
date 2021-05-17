package gosh_echo

import (
	"bufio"
	"fmt"
	"strings"
)

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	tolog := strings.Join(cmd[1:], " ")
	STDOUT.Write([]byte(fmt.Sprintf("%v", tolog) + "\n"))
	return 0
}

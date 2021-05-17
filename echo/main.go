package gosh_echo

import (
	"bufio"
	"fmt"
	"strings"
)

func Handle(cmd []string, STDIN *bufio.ReadWriter, STDOUT *bufio.ReadWriter, STDERR *bufio.ReadWriter, CWD string, ENV map[string]string) uint8 {
	tolog := strings.Join(cmd[1:], " ")
	STDOUT.Write([]byte(fmt.Sprintf("%v", tolog) + "\n"))
	return 0
}

package gosh_echo

import (
	"fmt"
	"os"
	"strings"
)

func Handle(cmd []string, STDIN *os.File, STDOUT *os.File, STDERR *os.File, CWD string, ENV map[string]string) uint8 {
	tolog := strings.Join(cmd[1:], " ")
	STDOUT.Write([]byte(fmt.Sprintf("%v", tolog) + "\n"))
	return 0
}

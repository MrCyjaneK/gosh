package gosh_exit

import (
	"bufio"
	"os"
	"strconv"
)

func Handle(cmd []string, STDIN *bufio.ReadWriter, STDOUT *bufio.ReadWriter, STDERR *bufio.ReadWriter, CWD string, ENV map[string]string) uint8 {
	if len(cmd) >= 2 {
		ec, err := strconv.ParseInt(cmd[1], 10, 8)
		if err != nil {
			STDERR.WriteString(err.Error())
			os.Exit(255)
		}
		os.Exit(int(ec))
	}
	os.Exit(0)
	return 255
}

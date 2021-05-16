package gosh_exit

import (
	"os"
	"strconv"
)

func Handle(cmd []string, STDIN *os.File, STDOUT *os.File, STDERR *os.File, CWD string, ENV map[string]string) uint8 {
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

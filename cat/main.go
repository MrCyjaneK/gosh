package gosh_cat

import (
	"bufio"
	"io"
	"os"
)

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	if len(cmd) == 1 {
		io.Copy(STDOUT, STDIN)
		return 0
	}
	var errcode uint8
	for i := range cmd {
		if i == 0 {
			continue
		}
		file, err := os.Open(cmd[i])
		if err != nil {
			errcode = 1
			STDERR.WriteString(err.Error() + "\n")
			continue
		}
		io.Copy(STDOUT, file)
	}
	return errcode
}

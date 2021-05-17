package gosh_ls

import (
	"bufio"
	"io/ioutil"
)

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	files, err := ioutil.ReadDir(CWD)
	if err != nil {
		STDERR.WriteString(err.Error() + "\n")
		return 1
	}

	for _, file := range files {
		STDOUT.WriteString(file.Name() + "\n")
	}
	return 0
}

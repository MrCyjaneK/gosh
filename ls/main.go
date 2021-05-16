package gosh_ls

import (
	"io/ioutil"
	"os"
)

func Handle(cmd []string, STDIN *os.File, STDOUT *os.File, STDERR *os.File, CWD string, ENV map[string]string) uint8 {
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

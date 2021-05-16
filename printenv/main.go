package gosh_printenv

import (
	"os"
)

func Handle(cmd []string, STDIN *os.File, STDOUT *os.File, STDERR *os.File, CWD string, ENV map[string]string) uint8 {
	for key := range ENV {
		STDOUT.WriteString(key + "=" + ENV[key] + "\n")
	}
	return 0
}

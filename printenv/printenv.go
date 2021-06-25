package gosh_printenv

import (
	"bufio"
)

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	if len(cmd) == 1 {
		for key := range ENV {
			STDOUT.WriteString(key + "=" + ENV[key] + "\n")
		}
	} else {
		ret := uint8(0)
		cmd = append(cmd[1:])
		for i := range cmd {
			key := cmd[i]
			val, ok := ENV[key]
			if !ok {
				ret = 1
				continue
			}
			STDOUT.WriteString(val + "\n")
		}
		return ret
	}
	return 0
}

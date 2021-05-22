package gosh_soexec

import (
	"bufio"
	"plugin"
)

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	if len(cmd) <= 1 {
		STDERR.WriteString("Usage: soexec path/to/plugin.so\n")
		return 1
	}
	p, err := plugin.Open(cmd[1])
	if err != nil {
		STDERR.WriteString(err.Error() + " (1)\n")
		return 1
	}
	pHandle, err := p.Lookup("Handle")
	if err != nil {
		STDERR.WriteString(err.Error() + " (2)\n")
		return 1
	}

	return pHandle.(func(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8)(cmd[2:], STDIN, STDOUT, STDERR, CWD, ENV)
}

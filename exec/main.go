package gosh_exec

import (
	"bufio"
	"os/exec"
)

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	if len(cmd) == 1 {
		STDERR.WriteString("No command provided!")
		return 127
	}
	var args []string
	if len(cmd) > 1 {
		args = cmd[2:]
	}
	execcmd := exec.Command(cmd[1], args...)
	execcmd.Stderr = STDERR
	execcmd.Stdout = STDOUT
	execcmd.Stdin = STDIN
	execcmd.Run()
	return 0
}

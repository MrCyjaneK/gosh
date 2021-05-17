package gosh

import (
	"bufio"
	"io/fs"
	"os"
	"path"
	"strings"
)

func cdHandle(cmd []string, STDIN *bufio.ReadWriter, STDOUT *bufio.ReadWriter, STDERR *bufio.ReadWriter, XCWD string, ENV map[string]string) uint8 {
	if len(cmd) != 2 {
		STDOUT.WriteString("Incorrect arguments provided!\n")
		return 1
	}
	var dir fs.FileInfo
	var err error
	if strings.HasPrefix(cmd[1], "/") {
		dir, err = os.Stat(cmd[1])
	} else {
		dir, err = os.Stat(CWD + "/" + cmd[1])
	}
	if os.IsNotExist(err) {
		STDERR.WriteString("Directory doesn't exist\n")
		return 1
	}
	if !dir.IsDir() {
		STDERR.WriteString("'" + cmd[1] + "' is not a directory\n")
		return 1
	}
	if strings.HasPrefix(cmd[1], "/") {
		CWD = path.Join(cmd[1])
	} else {
		CWD = path.Join(CWD, cmd[1])
	}
	return 0
}

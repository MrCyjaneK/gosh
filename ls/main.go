package gosh_ls

import (
	"bufio"
	"io/ioutil"
	"strings"
)

var flags map[string]string

func Handle(cmd []string, STDIN *bufio.ReadWriter, STDOUT *bufio.ReadWriter, STDERR *bufio.ReadWriter, CWD string, ENV map[string]string) uint8 {
	files, err := ioutil.ReadDir(CWD)
	if err != nil {
		STDERR.WriteString(err.Error() + "\n")
		return 1
	}
	flags = parseFlags(cmd)
	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, ".") && !(getFlagBool("a") || getFlagBool("all")) {
			continue
		}
		STDOUT.WriteString(name + "\n")
	}
	return 0
}

func getFlagBool(flag string) bool {
	_, ok := flags[flag]
	return ok
}

func getFlagString(flag string) string {
	a, ok := flags[flag]
	if !ok {
		return ""
	}
	return a
}

func parseFlags(cmd []string) map[string]string {
	var flags = make(map[string]string)
	for i := range cmd {
		if strings.HasPrefix(cmd[i], "--") && len(cmd[1]) != 2 {
			if len(strings.SplitN(cmd[i], "=", 2)) == 2 {
				flags[strings.SplitN(cmd[i][2:], "=", 2)[0]] = strings.SplitN(cmd[i], "=", 2)[1]
			} else if len(cmd) > i+1 {
				flags[cmd[i][2:]] = cmd[i+1]
			} else {
				flags[cmd[i][2:]] = "true"
			}
		} else if strings.HasPrefix(cmd[i], "-") && len(cmd[1]) == 2 {
			if len(strings.SplitN(cmd[i], "=", 2)) == 2 {
				flags[strings.SplitN(cmd[i][1:], "=", 2)[0]] = strings.SplitN(cmd[i], "=", 2)[1]
			} else if len(cmd) > i+1 {
				flags[cmd[i][1:]] = cmd[i+1]
			} else {
				flags[cmd[i][1:]] = "true"
			}
		} else if strings.HasPrefix(cmd[i], "-") && len(cmd[1]) > 2 {
			if len(cmd) > i+1 {
				flags[cmd[i][1:]] = cmd[i+1]
			} else {
				flags[cmd[i][1:]] = "true"
			}
			for j := range cmd[i][1:] {
				flags[string(cmd[i][j+1])] = "true"
			}
		}
	}
	return flags
}

package h

import "strings"

func GetFlagBool(flag string, flags map[string]string) bool {
	_, ok := flags[flag]
	return ok
}

func GetFlagString(flag string, flags map[string]string) string {
	a, ok := flags[flag]
	if !ok {
		return ""
	}
	return a
}

func ParseFlags(cmd []string) map[string]string {
	flags := make(map[string]string)
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

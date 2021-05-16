package gosh

import "strconv"

func getPrompt() string {
	if ERRCODE != 0 {
		return "user@gosh " + CWD + " [" + strconv.Itoa(int(ERRCODE)) + "]> "
	}
	return "user@gosh " + CWD + " > "
}

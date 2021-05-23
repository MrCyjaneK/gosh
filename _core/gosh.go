package gosh

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	STDIN  *bufio.Reader
	STDOUT *bufio.Writer
	STDERR *bufio.Writer

	CWD = "/"

	ERRCODE uint8

	ENV = make(map[string]string)
)

func Start(stdin *os.File, stdout *os.File, stderr *os.File) {
	var err error
	CWD, err = os.Getwd()
	if err != nil {
		CWD = "/"
	}
	STDIN = bufio.NewReader(stdin)
	STDOUT = bufio.NewWriter(stdout)
	STDERR = bufio.NewWriter(stderr)

	loadenv()
	input := bufio.NewScanner(STDIN)
	STDOUT.Write([]byte(getPrompt()))
	STDOUT.Flush()
	for input.Scan() {
		os.Chdir(CWD)
		STDOUT.Flush()
		ERRCODE = 0
		text := input.Text()
		var re = regexp.MustCompile(`\$[a-zA-Z0-9]+`)
		vars := re.FindAllString(text, -1)
		for i := range vars {
			en, ok := ENV[vars[i][1:]]
			if !ok {
				continue
			}
			text = strings.ReplaceAll(text, vars[i], en)
		}
		cmd, err := Split(text)
		if err != nil {
			Err(err)
			STDERR.Flush()
			ERRCODE = 126
			stdout.Write([]byte(getPrompt()))
			STDOUT.Flush()
			continue
		}
		if len(cmd) == 0 {
			Err("No input command given!")
			STDERR.Flush()
			ERRCODE = 126
			stdout.Write([]byte(getPrompt()))
			STDOUT.Flush()
			continue
		}
		if cmd[len(cmd)-1] == "&" {
			cmd = cmd[:len(cmd)-1]
			go handlecmd(cmd, STDIN, STDOUT, STDERR)
			STDOUT.Flush()
			STDERR.Flush()
		} else {
			handlecmd(cmd, STDIN, STDOUT, STDERR)
			STDOUT.Flush()
			STDERR.Flush()
		}
		stdout.Write([]byte(getPrompt()))
		STDOUT.Flush()
		STDERR.Flush()
	}
}

func Log(tolog interface{}) {
	STDOUT.Write([]byte(fmt.Sprintf("%v", tolog) + "\n"))
}

func Err(tolog interface{}) {
	STDERR.Write([]byte(fmt.Sprintf("%v", tolog) + "\n"))
}

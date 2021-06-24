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
		var re = regexp.MustCompile(`\$[a-zA-Z0-9_]+`)
		vars := re.FindAllString(text, -1)
		for i := range vars {
			en, ok := ENV[vars[i][1:]]
			if !ok {
				continue
			}
			text = strings.ReplaceAll(text, vars[i], en)
		}
		cmds, err := Split(text)
		var cmd []string
		for c := range cmds {
			cmd = append(cmd, cmds[c])
			if cmds[c] == ";" || cmds[c][len(cmds[c])-1] == ';' || len(cmds)-1 == c {
				if err != nil {
					Err(err)
					STDERR.Flush()
					ERRCODE = 126
					STDOUT.Flush()
					cmd = []string{}
					continue
				}
				if len(cmd) == 0 {
					Err("No input command given!")
					STDERR.Flush()
					ERRCODE = 126
					STDOUT.Flush()
					cmd = []string{}
					continue
				}
				if cmd[len(cmd)-1] == "&" {
					cmd = cmd[:len(cmd)-1]
					go func() {
						handlecmd(cmd, STDIN, STDOUT, STDERR)
						STDOUT.Flush()
						STDERR.Flush()
					}()
				} else {
					handlecmd(cmd, STDIN, STDOUT, STDERR)
					STDOUT.Flush()
					STDERR.Flush()
				}
				cmd = []string{}
			}
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

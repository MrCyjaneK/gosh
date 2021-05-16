package gosh

import (
	"bufio"
	"fmt"
	"os"

	gosh_cat "git.mrcyjanek.net/mrcyjanek/gosh/cat"
	gosh_echo "git.mrcyjanek.net/mrcyjanek/gosh/echo"
	gosh_ls "git.mrcyjanek.net/mrcyjanek/gosh/ls"
	gosh_printenv "git.mrcyjanek.net/mrcyjanek/gosh/printenv"
)

var (
	STDIN  *os.File
	STDOUT *os.File
	STDERR *os.File

	CWD = "/"

	ERRCODE uint8

	ENV = make(map[string]string)
)

func Start(stdin *os.File, stdout *os.File, stderr *os.File) {
	STDIN = stdin
	STDOUT = stdout
	STDERR = stderr
	ERRCODE = 0
	loadenv()
	input := bufio.NewScanner(STDIN)
	STDOUT.Write([]byte(getPrompt()))
	for input.Scan() {
		os.Chdir(CWD)

		ERRCODE = 0
		cmd, err := Split(input.Text())
		if err != nil {
			Err(err)
			ERRCODE = 126
			STDOUT.Write([]byte(getPrompt()))
			continue
		}
		if len(cmd) == 0 {
			Err("No input command given!")
			ERRCODE = 126
			STDOUT.Write([]byte(getPrompt()))
			continue
		}
		switch cmd[0] {
		case "cat":
			ERRCODE = gosh_cat.Handle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
		case "cd":
			ERRCODE = cdHandle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
		case "echo":
			ERRCODE = gosh_echo.Handle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
		case "ls":
			ERRCODE = gosh_ls.Handle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
		case "printenv":
			ERRCODE = gosh_printenv.Handle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
		default:
			ERRCODE = 127
			Err("Command '" + cmd[0] + "' not found.")
		}
		STDOUT.Write([]byte(getPrompt()))
	}
}

func Log(tolog interface{}) {
	STDOUT.Write([]byte(fmt.Sprintf("%v", tolog) + "\n"))
}

func Err(tolog interface{}) {
	STDERR.Write([]byte(fmt.Sprintf("%v", tolog) + "\n"))
}

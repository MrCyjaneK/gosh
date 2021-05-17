package gosh

import (
	"bufio"
	"log"
	"os"

	gosh_cat "git.mrcyjanek.net/mrcyjanek/gosh/cat"
	gosh_echo "git.mrcyjanek.net/mrcyjanek/gosh/echo"
	gosh_exec "git.mrcyjanek.net/mrcyjanek/gosh/exec"
	gosh_exit "git.mrcyjanek.net/mrcyjanek/gosh/exit"
	gosh_ls "git.mrcyjanek.net/mrcyjanek/gosh/ls"
	gosh_printenv "git.mrcyjanek.net/mrcyjanek/gosh/printenv"
)

func Handlecmd(cmd []string, tostdin []byte) (stdout *bufio.ReadWriter, stderr *bufio.ReadWriter) {
	//var b bytes.Buffer
	//b.Write(tostdin)
	//stdin := bufio.NewReadWriter(bufio.NewReader(&b), bufio.NewWriter(&b))
	//var out bytes.Buffer
	//stdout = bufio.NewReadWriter(bufio.NewReader(*out), bufio.NewWriter(&out))
	//var err bytes.Buffer
	//stderr = bufio.NewReadWriter(bufio.NewReader(&err), bufio.NewWriter(&err))4
	stdinr, stdinw, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	stdoutr, stdoutw, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	stderrr, stderrw, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	STDIN = bufio.NewReadWriter(bufio.NewReader(stdinr), bufio.NewWriter(stdinw))
	STDOUT = bufio.NewReadWriter(bufio.NewReader(stdoutr), bufio.NewWriter(stdoutw))
	STDERR = bufio.NewReadWriter(bufio.NewReader(stderrr), bufio.NewWriter(stderrw))
	handlecmd(cmd, STDIN, STDOUT, STDERR)
	return STDOUT, STDERR
}

func handlecmd(cmd []string, stdin *bufio.ReadWriter, stdout *bufio.ReadWriter, stderr *bufio.ReadWriter) {
	switch cmd[0] {
	case "cat":
		ERRCODE = gosh_cat.Handle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
	case "cd":
		ERRCODE = cdHandle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
	case "echo":
		ERRCODE = gosh_echo.Handle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
	case "exec":
		ERRCODE = gosh_exec.Handle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
	case "exit":
		ERRCODE = gosh_exit.Handle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
	case "ls":
		ERRCODE = gosh_ls.Handle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
	case "printenv":
		ERRCODE = gosh_printenv.Handle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
	default:
		ERRCODE = 127
		Err("Command '" + cmd[0] + "' not found.")
	}
}

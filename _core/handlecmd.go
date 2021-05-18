package gosh

import (
	"bufio"
	"log"
	"os"

	gosh_cat "git.mrcyjanek.net/mrcyjanek/gosh/cat"
	gosh_debug "git.mrcyjanek.net/mrcyjanek/gosh/debug"
	gosh_echo "git.mrcyjanek.net/mrcyjanek/gosh/echo"
	gosh_exec "git.mrcyjanek.net/mrcyjanek/gosh/exec"
	gosh_exit "git.mrcyjanek.net/mrcyjanek/gosh/exit"
	gosh_ls "git.mrcyjanek.net/mrcyjanek/gosh/ls"
	gosh_printenv "git.mrcyjanek.net/mrcyjanek/gosh/printenv"
)

func Handlecmd(cmd []string, tostdin []byte) (stdout *bufio.Reader, stderr *bufio.Reader) {
	//var b bytes.Buffer
	//b.Write(tostdin)
	//stdin := bufio.NewReadWriter(bufio.NewReader(&b), bufio.NewWriter(&b))
	//var out bytes.Buffer
	//stdout = bufio.NewReadWriter(bufio.NewReader(*out), bufio.NewWriter(&out))
	//var err bytes.Buffer
	//stderr = bufio.NewReadWriter(bufio.NewReader(&err), bufio.NewWriter(&err))4
	_, stdinr, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	stdoutw, stdoutr, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	stderrw, stderrr, err := os.Pipe()
	if err != nil {
		log.Fatal(err)
	}
	STDIN = bufio.NewReaderSize(bufio.NewReader(stdinr), 1)
	STDOUT = bufio.NewWriterSize(bufio.NewWriter(stdoutw), 1)
	STDERR = bufio.NewWriterSize(bufio.NewWriter(stderrw), 1)
	handlecmd(cmd, STDIN, STDOUT, STDERR)
	return bufio.NewReader(stdoutr), bufio.NewReader(stderrr)
}

func handlecmd(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer) {
	switch cmd[0] {
	case "cat":
		ERRCODE = gosh_cat.Handle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
	case "cd":
		ERRCODE = cdHandle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
	case "debug":
		ERRCODE = gosh_debug.Handle(cmd, STDIN, STDOUT, STDERR, CWD, ENV)
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

package gosh

import (
	gosh_cat "git.mrcyjanek.net/mrcyjanek/gosh/cat"
	gosh_echo "git.mrcyjanek.net/mrcyjanek/gosh/echo"
	gosh_exec "git.mrcyjanek.net/mrcyjanek/gosh/exec"
	gosh_exit "git.mrcyjanek.net/mrcyjanek/gosh/exit"
	gosh_ls "git.mrcyjanek.net/mrcyjanek/gosh/ls"
	gosh_printenv "git.mrcyjanek.net/mrcyjanek/gosh/printenv"
)

func Handlecmd(cmd []string) {
	handlecmd(cmd)
}

func handlecmd(cmd []string) {
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

package gosh

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var (
	STDIN  *bufio.ReadWriter
	STDOUT *bufio.ReadWriter
	STDERR *bufio.ReadWriter

	CWD = "/"

	ERRCODE uint8

	ENV = make(map[string]string)
)

func Start(stdin *os.File, stdout *os.File, stderr *os.File) {
	STDIN = bufio.NewReadWriter(bufio.NewReader(stdin), bufio.NewWriter(stdin))
	STDOUT = bufio.NewReadWriter(bufio.NewReader(stdout), bufio.NewWriter(stdout))
	STDERR = bufio.NewReadWriter(bufio.NewReader(stderr), bufio.NewWriter(stderr))
	go func() {
		for {
			STDOUT.Flush()
			STDERR.Flush()
			time.Sleep(time.Second / 10)
		}
	}()
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
		if cmd[len(cmd)-1] == "&" {
			cmd = cmd[:len(cmd)-1]
			go handlecmd(cmd, STDIN, STDOUT, STDERR)
		} else {
			handlecmd(cmd, STDIN, STDOUT, STDERR)
		}
		STDOUT.Write([]byte(getPrompt()))
	}
}

func Log(tolog interface{}) {
	STDOUT.Flush()
	STDERR.Flush()
	STDOUT.Write([]byte(fmt.Sprintf("%v", tolog) + "\n"))
}

func Err(tolog interface{}) {
	STDOUT.Flush()
	STDERR.Flush()
	STDERR.Write([]byte(fmt.Sprintf("%v", tolog) + "\n"))
}

package main

// Build with
// go build -buildmode=plugin
import "bufio"

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	STDOUT.WriteString("Hello, World!\n")
	return 69
}

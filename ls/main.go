package gosh_ls

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"strings"

	"git.mrcyjanek.net/mrcyjanek/gosh/_core/h"
)

var flags map[string]string

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	flags, cmd = h.ParseFlags(cmd)
	toret := uint8(0)
	if len(cmd) == 1 {
		files, err := ioutil.ReadDir(CWD)
		if err != nil {
			STDERR.WriteString(err.Error() + "\n")
			return 1
		}
		printFiles(files, STDOUT)
	} else {
		for i := 1; i < len(cmd); i++ {
			var files []fs.FileInfo
			var err error
			if cmd[i][:1] == "/" {
				files, err = ioutil.ReadDir(cmd[i])
			} else {
				files, err = ioutil.ReadDir(CWD + "/" + cmd[i])
			}
			if err != nil {
				STDERR.WriteString(err.Error() + "\n")
				toret = 2
				continue
			}
			if i != 1 {
				STDOUT.WriteString("\n")
			}
			STDOUT.WriteString(cmd[i] + ":\n")
			printFiles(files, STDOUT)
		}
	}
	return toret
}

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

func getSize(size int64) string {
	if h.GetFlagBool("h", flags) {
		return ByteCountIEC(size)
	} else {
		return fmt.Sprintf("%d", size)
	}
}

func printFiles(files []fs.FileInfo, STDOUT *bufio.Writer) {
	for _, file := range files {
		name := file.Name()
		if strings.HasPrefix(name, ".") && !(h.GetFlagBool("a", flags) || h.GetFlagBool("all", flags)) {
			continue
		}
		if !h.GetFlagBool("l", flags) {
			STDOUT.WriteString(name + "\n")
		} else {
			STDOUT.WriteString(file.Mode().String() + " " + getSize(file.Size()) + "\t" + file.ModTime().String() + "\t" + file.Name() + "\n")
		}
	}
}

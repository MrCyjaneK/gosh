package gosh_sha512sum

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"os"
	"strings"

	"git.mrcyjanek.net/mrcyjanek/gosh/_core/h"
)

var flags map[string]string

func Handle(cmd []string, STDIN *bufio.Reader, STDOUT *bufio.Writer, STDERR *bufio.Writer, CWD string, ENV map[string]string) uint8 {
	flags = h.ParseFlags(cmd)

	tolog := strings.Join(cmd[1:], " ")
	if tolog == "" {
		// No input file provided, let's read stdin.
	}
	errcode := uint8(0)
	var files []string
	for i := range cmd {
		if cmd[i] == "wget" {
			continue
		}
		unused := true
		for j := range flags {
			if cmd[i] == flags[j] || cmd[i] == "-"+j || cmd[i] == "--"+j {
				unused = false
			}
		}
		if unused {
			files = append(files, cmd[i])
		}
	}
	for i := range files {
		if i == 0 {
			continue
		}
		file, err := os.Open(files[i])
		if err != nil {
			errcode = 1
			STDERR.WriteString(err.Error() + "\n")
			continue
		}
		hasher := sha512.New()
		io.Copy(hasher, file)
		sum := hex.EncodeToString(hasher.Sum(nil))
		if !h.GetFlagBool("quiet", flags) {
			STDOUT.WriteString(sum + "\t" + files[i] + "\n")
		}
	}

	//STDOUT.Write([]byte(fmt.Sprintf("%v", tolog) + "\n"))
	return errcode
}

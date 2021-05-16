package main

import (
	"os"

	gosh "git.mrcyjanek.net/mrcyjanek/gosh/_core"
)

func main() {
	gosh.Start(os.Stdin, os.Stdout, os.Stderr)
}

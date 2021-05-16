package gosh

import (
	"os"
	"strings"
)

func loadenv() {
	envs := os.Environ()
	for i := range envs {
		env := strings.SplitN(envs[i], "=", 2)
		ENV[env[0]] = env[1]
	}
}

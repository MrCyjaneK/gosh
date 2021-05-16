# Gosh

Simple shell written in golang.

To start an interactive shell you can do:

```go
package main

import (
	"os"

	gosh "git.mrcyjanek.net/mrcyjanek/gosh/_core"
)

func main() {
	gosh.Start(os.Stdin, os.Stdout, os.Stderr)
}
```

Every command runs in separated environment (changing path in one command doesn't change it for rest). That's why `cd` is part of `_core/`, and `echo` is not. I think that this is a good idea... But I'm not sure (yet), about that.

Currently only basics work `cd`, `ls`, `cat`, `printenv`, nothing that could be actually used, no `;`, `|`, `&&`, `&` and even no external command execution. Nothing.

If you need only one command, for example `ls`, you can include only it in your project [check source](https://git.mrcyjanek.net/mrcyjanek/gosh/src/branch/master/ls/main.go).

Exit codes support is provided.

I'll try to stay as close as possible to `sh`.
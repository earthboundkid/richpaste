package main

import (
	"os"

	"github.com/carlmjohnson/exitcode"
	"github.com/earthboundkid/richpaste/pasteboard"
)

func main() {
	exitcode.Exit(pasteboard.CLI(os.Args[1:]))
}

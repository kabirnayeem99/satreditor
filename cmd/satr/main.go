package main

import (
	"fmt"
	"os"

	"github.com/kabirnayeem99/satreditor/internal/editor"
)

func main() {
	if err := editor.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

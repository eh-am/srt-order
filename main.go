package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/eh-am/srt-order/internal"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprint(os.Stderr, "usage ./strorder {files}")
		os.Exit(1)
	}
	wordPtr := flag.String("word", "foo", "a string")

	files := os.Args[1:]
	for _, f := range files {
		err := internal.Process(f)
		if err != nil {
			fmt.Fprint(os.Stderr, "err: %w", err)
			os.Exit(1)
		}
	}
}

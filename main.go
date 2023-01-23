package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/eh-am/srt-order/internal"
)

func main() {
	inPlace := flag.Bool("in-place", false, "substitute in place")
	flag.Parse()

	values := flag.Args()
	if len(values) <= 0 {
		fmt.Fprint(os.Stderr, "usage ./strorder {flags} {files}\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	for _, f := range values {
		err := internal.Process(f, *inPlace)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
			os.Exit(1)
		}
	}
}

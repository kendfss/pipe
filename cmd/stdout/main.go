// automatically genrated
// date: 2025-02-13 12:13:28.171104619 &#43;0100 WAT m=&#43;0.000123328
package main

import (
	"flag"
	"os"
	"strings"

	"github.com/kendfss/pipe"
)

var mirrorFlag bool

func init() {
	flag.BoolVar(&mirrorFlag, "m", false, "also write output to stderr")
}

func main() {
	flag.Parse()
	input := pipe.Get()
	if len(input) == 0 {
		input = []byte(strings.Join(flag.Args(), " "))
		if len(input) == 0 {
			flag.Usage()
		}
		input = append(input, '\n')
	}
	os.Stdout.Write(input)
	if mirrorFlag {
		os.Stderr.Write(input)
	}
}

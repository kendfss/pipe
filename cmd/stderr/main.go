// automatically genrated
// date: 2025-02-12 16:50:20.079232747 &#43;0100 WAT m=&#43;0.000151510
package main

import (
	"flag"
	"os"
	"strings"

	"github.com/kendfss/pipe"
)

func main() {
	flag.Parse()
	input := pipe.Get()
	if len(input) == 0 {
		input = []byte(strings.Join(flag.Args(), " "))
		if len(input) == 0 {
			flag.Usage()
		}
	}
	os.Stderr.Write(append(input, '\n'))
}


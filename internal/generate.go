//go:generate go run .
package main

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

func main() {
	date := time.Now()
	todos := []string{
		"out",
		"err",
	}
	for _, todo := range todos {
		f := mustv(os.Create(fname(todo)))
		defer f.Close()
		must(template.Must(template.New("").Parse(templateString)).Execute(f, struct {
			Date time.Time
			Todo string
		}{
			Date: date,
			Todo: todo,
		}))
	}
}

const templateString = `// automatically genrated
// date: {{ printf "%s" .Date }}
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
	os.Std{{ .Todo }}.Write(append(input, '\n'))
}

`

func fname(s string) string {
	return fmt.Sprintf("../cmd/std%s/main.go", s)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func mustv[T any](t T, err error) T {
	must(err)
	return t
}

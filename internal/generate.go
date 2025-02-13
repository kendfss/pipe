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
		must(template.Must(template.New("").Funcs(template.FuncMap{
			"mirror": func(s string) string {
				switch s {
				case "err":
					return "out"
				case "out":
					return "err"
				default:
					panic(fmt.Errorf("%q is not a supported output", s))
				}
			},
		}).Parse(templateString)).Execute(f, struct {
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

var mirrorFlag bool

func init() {
	flag.BoolVar(&mirrorFlag, "m", false, "also write output to std{{ mirror .Todo }}")
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
	os.Std{{ .Todo }}.Write(input)
	if mirrorFlag {
		os.Std{{ mirror .Todo }}.Write(input)
	}
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

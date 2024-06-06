pipe
---

Small utility for getting pipes from stdin and other files

# exports

```go
func Empty() bool
    check if Stdin is empty

func Emptyf(f fs.File) bool
    check if a file is empty

func From(f fs.File) ([]byte, error)
    Extract the bytes from a file

func FromErr() ([]byte, error)
    Extract the bytes from /dev/stderr

func FromIn() ([]byte, error)
    Extract the bytes from /dev/stdin

func FromOut() ([]byte, error)
    Extract the bytes from /dev/stdout

func Get() []byte
    get data from Stdin

func Getf(f fs.File) []byte
    get data from a named pipe

func Loaded(stat fs.FileInfo) bool
    Loaded checks if a file's status implies that it can be read from

func Receiving() bool
    check for incoming data

func Size() int64
    check size (in bytes) of Stdin

func Sizef(f fs.File) int64
    check size (in bytes) of a file.
```

# usage

```go
package main

import "github.com/kendfss/pipe"


func main() {
    if data := pipe.Get(); data != nil {
        processInput(data) // this exercise is left to the reader
    } 
}
```

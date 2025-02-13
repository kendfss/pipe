package pipe

import (
	"io"
	"io/fs"
	"os"
)

// Extract the bytes from a file
func From(f fs.File) ([]byte, error) {
	var data []byte
	info, err := f.Stat()
	if err == nil {
		if Loaded(info) {
			data, err = io.ReadAll(f)
		}
	}
	return data, err
}

// Extract the bytes from /dev/stdin
func FromIn() ([]byte, error) {
	return From(os.Stdin)
}

// Extract the bytes from /dev/stdout
func FromOut() ([]byte, error) {
	return From(os.Stdout)
}

// Extract the bytes from /dev/stderr
func FromErr() ([]byte, error) {
	return From(os.Stderr)
}

// Loaded checks if a file's status implies that it can be read from
func Loaded(stat fs.FileInfo) bool {
	return (stat.Mode() & os.ModeCharDevice) == 0
}

// Receiving reports whether or not we're receiving on stdin
func Receiving() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	return (stat.Mode() & os.ModeCharDevice) == 0
}

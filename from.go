package pipe

import (
	"io/fs"
	"os"
)

// Extract the bytes from a file
func From(f *os.File) ([]byte, error) {
	var data []byte
	info, err := f.Stat()
	if err == nil {
		if Loaded(info) {
			data = make([]byte, info.Size())
			_, err = f.Read(data)
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

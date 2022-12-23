package pipe

import (
	"os"

	"github.com/kendfss/but"
)

func Stdin() *os.File {
	f, err := os.Open("/dev/stdin")
	but.Exiff(err, "pipe: %s", err)
	return f
}

func Stdout() *os.File {
	f, err := os.Open("/dev/stdout")
	but.Exiff(err, "pipe: %s", err)
	return f
}

func Stderr() *os.File {
	f, err := os.Open("/dev/stderr")
	but.Exiff(err, "pipe: %s", err)
	return f
}

// get data from Stdin
func Get() []byte {
	// return Getf(os.Stdout)
	return Getf(os.Stdin)
}

// get data from a file
// panic on error
func Getf(f *os.File) []byte {
	data, err := From(f)
	but.Exiff(err, "pipe: %s", err)
	return data
}

// check if Stdin is empty
func Empty() bool {
	return Emptyf(os.Stdin)
}

// check if a file is empty
func Emptyf(f *os.File) bool {
	stat, err := f.Stat()
	but.Exif(err)
	return Loaded(stat)
}

// check size (in bytes) of Stdin
func Size() int64 {
	return Sizef(os.Stdin)
}

// check size (in bytes) of a file.
// Exits if any Stat error is incurred error
func Sizef(f *os.File) int64 {
	info, err := f.Stat()
	but.Exiff(err, "pipe: %s", err)
	return info.Size()
}

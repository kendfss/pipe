package pipe

import (
	"io/fs"
	"os"
)

// get data from Stdin
func Get() []byte {
	data, _ := From(os.Stdin)
	return data
}

// get data from a named pipe
func Getf(f fs.File) []byte {
	data, _ := From(f)
	return data
}

// check if Stdin is empty
func Empty() bool {
	return Emptyf(os.Stdin)
}

// check if a file is empty
func Emptyf(f fs.File) bool {
	return Sizef(f) == 0
}

// check size (in bytes) of Stdin
func Size() int64 {
	return Sizef(os.Stdin)
}

// check size (in bytes) of a file.
func Sizef(f fs.File) int64 {
	info, _ := f.Stat()
	return info.Size()
}

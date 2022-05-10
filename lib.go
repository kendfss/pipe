package pipe

import (
	"io/ioutil"
	"os"

	"github.com/kendfss/but"
)

// get data from stdin
// panic on error
func Get() []byte {
	var (
		data []byte
		err  error
	)

	if !Empty() {
		data, err = ioutil.ReadAll(os.Stdin)
		but.Must(err)
	}

	return data
}

// get data, with an error from stdin
func GetE() ([]byte, error) {
	var data []byte
	info, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if info.Size() > 0 {
		data, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

// check if stdin is empty
func Empty() bool {
	return Size() == 0
}

// check size (in bytes) of stdin
func Size() int64 {
	info, err := os.Stdin.Stat()
	but.Must(err)

	return info.Size()
}

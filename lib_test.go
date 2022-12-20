package pipe

import (
	"os"
	"testing"
)

const (
	msg  = "hello dev"
	mLen = len(msg)
)

func TestFrom(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	n, err := w.Write([]byte(msg))
	if err != nil {
		t.Fatal(err)
	}
	if n != mLen {
		t.Fatalf("should've written %d bytes, wrote %d", n, mLen)
	}

	data, err := From(r)
	if err != nil {
		t.Fatal(err)
	}

	if out := string(data); out == msg {
		t.Fatalf("got %q, wanted %q", data, msg)
	}
}

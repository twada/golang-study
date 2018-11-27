package mustcopy

import (
	"bytes"
	"fmt"
	"testing"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func TestMustCopy(t *testing.T) {
	var buf bytes.Buffer
	mustCopy(os.Stdout, conn)

	s := &IntSet{}
	actual := s.Has(3)
	expected := false
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}

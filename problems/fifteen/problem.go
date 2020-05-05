package fifteen

import (
	"io"
	"math/rand"
)

// Fifteen returns a random element from a large stream.
func Fifteen(e io.Reader) byte {

	var r byte
	b := make([]byte, 1)
	for i := 0; ; i++ {
		_, err := e.Read(b)

		if err == io.EOF {
			break
		}

		if i == 0 {
			r = b[0]
		} else if rand.Intn(i+1) == i {
			r = b[0]
		}
	}

	return r

}

package showstart

// see https://codeberg.org/ac/base62/src/branch/main/base62.go
import (
	"errors"
	"math"
	"strings"
)

const Base = 62

// CharacterSet consists of 62 characters [0-9][A-Z][a-z].
const CharacterSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Base62 struct {
}

// Encode returns a base62 representation as
// string of the given integer number.
func (*Base62) Encode(num uint32) string {
	b := make([]byte, 0)

	// loop as long the num is bigger than zero
	for num > 0 {
		// receive the rest
		r := math.Mod(float64(num), float64(Base))

		// devide by Base
		num /= Base

		// append chars
		b = append([]byte{CharacterSet[int(r)]}, b...)
	}

	return string(b)
}

// Decode returns a integer number of a base62 encoded string.
func (*Base62) Decode(s string) (uint32, error) {
	var r, pow int

	// loop through the input
	for i, v := range s {
		// convert position to power
		pow = len(s) - (i + 1)

		// IndexRune returns -1 if v is not part of CharacterSet.
		pos := strings.IndexRune(CharacterSet, v)

		if pos == -1 {
			return 0, errors.New("invalid character: " + string(v))
		}

		// calculate
		r += pos * int(math.Pow(float64(Base), float64(pow)))
	}

	return uint32(r), nil
}

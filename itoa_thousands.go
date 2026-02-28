package diagnosticism

import (
	"strconv"
)

func itoaThousands_int64(v int64) string {
	neg := v < 0

	s := strconv.FormatInt(v, 10)

	r := make([]byte, 0, len(s) * 4 / 3)

	if neg {
		s = s[1:]
		r = append(r, '-')
	}

	var n int
	switch len(s) % 3 {
	case 0:
		n = 0
	case 1:
		n = 2
	case 2:
		n = 1
	}

	for _, b := range []byte(s) {

		if n == 3 {
			r = append(r, ',')
			n = 0
		}
		n++

		r = append(r, b)
	}

	return string(r)
}

func itoaThousands_uint64(v uint64) string {
	s := strconv.FormatUint(v, 10)

	r := make([]byte, 0, len(s) * 4 / 3)

	var n int
	switch len(s) % 3 {
	case 0:
		n = 0
	case 1:
		n = 2
	case 2:
		n = 1
	}

	for _, b := range []byte(s) {

		if n == 3 {
			r = append(r, ',')
			n = 0
		}
		n++

		r = append(r, b)
	}

	return string(r)
}

// Returns the string form of v in base 10, with thousands separators.
func ItoaThousands[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr](v T) string {

	switch any(v).(type) {
	case int, int8, int16, int32, int64:

		return itoaThousands_int64(int64(v))
	case uint, uint8, uint16, uint32, uint64, uintptr:

		return itoaThousands_uint64(uint64(v))
	default:
		panic("VIOLATION: unexpected type")
	}
}

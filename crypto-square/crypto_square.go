package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func filterString(st string) string {
	// I later found that there is a strings.Map function
	// but that was slower. Benchmark results with this method
	// gave ~7000 ns/op whereas for that, it was ~8800 ns/op.
	var sb strings.Builder
	for _, c := range strings.ToLower(st) {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func determineRectSize(len int) (c, r int) {
	r = int(math.Sqrt(float64(len)))
	// Depending on the length, there are three possible sizes
	// 1. r x r
	// 2. r x (r+1)
	// 3. (r+1) x (r+1)
	c = r
	if len == r*c {
		return
	}
	c++
	if len <= r*c {
		return
	}
	r++
	return
}

// Encode does a cryptosquare encoding
func Encode(st string) string {
	st = filterString(st)
	c, r := determineRectSize(len(st))

	var sb strings.Builder
	for j := 0; j < c; j++ {
		for i := 0; i < r; i++ {
			x := i*c + j
			if x < len(st) {
				sb.WriteByte(st[x])
			} else {
				sb.WriteRune(' ')
			}
		}

		if j < c-1 {
			// Don't write a space for the last iteration
			sb.WriteRune(' ')
		}
	}

	return sb.String()
}

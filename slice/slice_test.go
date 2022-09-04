package slice

import (
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestSliceCut(t *testing.T) {
	x1 := make([]int, 2, 10)

	t.Log(x1[2:])

	x2 := appendSlice(x1)
	t.Log(x2[2:])

	t.Log(len(x1))
	t.Log(x1[:])

	// slice_test.go:12: []
	// slice_test.go:15: [2]
	// slice_test.go:17: 2
	// slice_test.go:18: [99 0]

	/*type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}*/
}

func appendSlice(slice []int) []int {
	slice[0] = 99
	slice = append(slice, 2)
	return slice
}

package _interface

import "testing"

func TestMain(m *testing.M) {
	m.Run()
}

func TestCompare(t *testing.T) {
	var a interface{}
	var b interface{}

	var c interface{} = []int{1, 2, 3}
	var d interface{} = []int{1, 2, 3}

	t.Log(a == b)
	t.Log(c == d)

	// output:
	//  interface_test.go:16: true
	//  interface_test.go:17: panic: runtime error: comparing uncomparable type []int [recovered]
}

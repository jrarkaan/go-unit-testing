package helper

import "testing"

// go test
// go test -v
// go test -v -run [name_func]

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// unit test failed
		panic("Result isnt hello Eko")
	}
}

func TestHelloWorldRaka(t *testing.T) {
	result := HelloWorld("Raka")
	if result != "Hello Raka" {
		// unit test failed
		panic("Result isnt hello Raka")
	}
}

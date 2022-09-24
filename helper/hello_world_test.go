package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// go test
// go test -v
// go test -v -run [name_func]

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// unit test failed
		//panic("Result isnt hello Eko")
		//t.Fail()
		t.Error("Harusnya Hello Eko")
	}
	fmt.Println("Done 1")
}

func TestHelloWorldRaka(t *testing.T) {
	result := HelloWorld("Raka")
	if result != "Hello Raka" {
		// unit test failed
		//panic("Result isnt hello Raka")
		//t.FailNow()
		t.Fatal("Harusnya Hello Raka")
	}
	fmt.Println("Done 2")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("Done 1")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("Done 1")
}

func TestSkip(t *testing.T) {
	result := HelloWorld("Eko")
	if runtime.GOOS == "darwin" {
		t.Skip("g bsa jalan di mac")
	}
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

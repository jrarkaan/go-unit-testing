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

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")
}

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

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result)
	})
	t.Run("Raka", func(t *testing.T) {
		result := HelloWorld("Raka")
		require.Equal(t, "Hello Raka", result)
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Raka",
			request:  "Raka",
			expected: "Hello Raka",
		},
		{
			name:     "Eko",
			request:  "Eko",
			expected: "Hello Eko",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Raka")
	}
}

func BenchmarkSubHelloWorld(b *testing.B) {
	b.Run("Raka", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Raka")
		}
	})
}

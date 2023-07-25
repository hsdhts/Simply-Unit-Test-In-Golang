package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

//Benchmark

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Husada")
	}
}
func BenchmarkHelloHutasoit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hutasoit")
	}
}

// Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Husada",
			request:  "Husada",
			expected: "Hello Husada",
		},
		{
			name:     "Hutasoit",
			request:  "Hutasoit",
			expected: "Hello Hutasoit",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}

}

// Sub Test
func TestSubTest(t *testing.T) {
	t.Run("Husada", func(t *testing.T) {
		result := HelloWorld("Husada")
		assert.Equal(t, "Hello Husada", result, "Result must be 'Hello Husada'")
	})
	t.Run("Hutasoit", func(t *testing.T) {
		result := HelloWorld("Hutasoit")
		assert.Equal(t, "Hello Hutasoit", result, "Result must be 'Hello Hutasoit'")
	})
}

// Main Test
func TestMain(m *testing.M) {
	//before
	fmt.Println("BEFORE RUN UNIT TEST")

	m.Run()

	//after
	fmt.Println("AFTER RUN UNIT TEST")
}

// Skip test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "Darwin" {
		t.Skip("Cannot run on Mac OS")
	}
	result := HelloWorld("Husada")
	assert.Equal(t, "Hello Husada", result, "Result must be 'Hello Husada'")

}

// assertion
func TestHelloWorldHusada(t *testing.T) {
	result := HelloWorld("Husada")
	assert.Equal(t, "Hello Husada", result, "Result must be 'Hello Husada'")
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldBudi(t *testing.T) {
	result := HelloWorld("Budi")
	assert.Equal(t, "Hello Budi", result, "Result must be 'Hello Budi'")
	fmt.Println("TestHelloWorldBudi Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Require")
	require.Equal(t, "Hello Require", result, "Result must be 'Hello Budi'")
	fmt.Println("TestHelloWorldBudi Done")
}

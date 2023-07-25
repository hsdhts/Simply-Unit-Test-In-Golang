package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

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

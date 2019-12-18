package phony

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	a := Get("name")
	b := Get("name")
	assert.NotEqual(t, a, "")
	assert.NotEqual(t, b, "")
}

func TestEmpty(t *testing.T) {
	assert.Equal(t, Get("foo"), "")
}

func TestAll(t *testing.T) {
	for _, p := range List() {
		assert.NotEqual(t, Get(p), "")
	}
}

package test

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestFoo(t *testing.T) {
	expectd := 1
	actual := 1

	assert.Equal(t, expectd, actual)
}

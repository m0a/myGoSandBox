package mymath_test

import (
	"testing"

	aaaa "github.com/m0a/myGoSandBox/test"

	// mymath "github.com/m0a/myGoSandBox/test"
)

///Users/m0a/go/src/github.com/m0a/myGoSandBox/test

func TestSum(t *testing.T) {
	actual := aaaa.Sum(10, 20)
	expected := 30
	if actual != expected {
		t.Errorf("got %v\n want %v ", actual, expected)
	}
}

package graphql

import (
	"testing"
	s "github.com/goplayground/schema"
)

func TestPlay(t *testing.T) {
	if Play() == 2 {
		t.Error("expected not to be 2")
	}
}

func TestPlay2(t *testing.T) {
	if s.Play() != 2 {
		t.Error("expected 2")
	}
}

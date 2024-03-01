package errorbase

import (
	"errors"
	"io"
	"testing"
)

func TestErrorChain(t *testing.T) {
	base := NewError("base error")
	if m := base.Error(); m != "base error" {
		t.Errorf("unexpected error message: %s", m)
	}

	derived := base.Derive("derived error")
	if m := derived.Error(); m != "derived error" {
		t.Errorf("unexpected error message: %s", m)
	}

	if !errors.Is(derived, base) {
		t.Errorf("expected derived error to wrap base error")
	}
}

func TestErrorDerive(t *testing.T) {
	err := Derive(io.EOF, "my EOF")
	if m := err.Error(); m != "my EOF" {
		t.Errorf("unexpected error message: %s", m)
	}

	if !errors.Is(err, io.EOF) {
		t.Errorf("expected derived error to wrap io.EOF")
	}
}

func TestDeriveNilError(t *testing.T) {
	err := Derive(nil, "my error")
	if err != nil {
		t.Errorf("expected nil error, got: %v", err)
	}
}

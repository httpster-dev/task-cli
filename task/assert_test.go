package task_test

import (
	"errors"
	"testing"
)

// assertEqual checks that got == want for any comparable type.
// The [T comparable] syntax is a Go generic — like a Ruby method that accepts any type
// as long as it supports ==. This replaces writing separate helpers for int, string, etc.
func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

// assertNoError fails the test if err is non-nil.
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// assertError fails the test if err is nil.
func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error but got nil")
	}
}

// assertErrorIs fails the test if err does not match target via errors.Is.
// Use this instead of assertError when you care which error was returned.
func assertErrorIs(t *testing.T, err, target error) {
	t.Helper()
	if !errors.Is(err, target) {
		t.Errorf("got error %v, want %v", err, target)
	}
}

// assertNotNil fails the test if got is nil.
// The [T any] constraint accepts any pointer or interface type.
func assertNotNil[T any](t *testing.T, got *T) {
	t.Helper()
	if got == nil {
		t.Error("expected a non-nil value but got nil")
	}
}

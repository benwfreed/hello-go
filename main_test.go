package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	expectedSalutation := "Hello"
	actualSalutation := getSalutation()
	if expectedSalutation != actualSalutation {
		t.Errorf("getSalutation() = %s; want %s", actualSalutation, expectedSalutation)
	}
}

package main

import (
	"bytes"
	"testing"
)

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("abc")

	expected := 3
	actual := count(false, true, b)

	if actual != expected {
		t.Errorf("expected %d, got %d instead.\n", expected, actual)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

	expected := 3
	actual := count(true, false, b)

	if actual != expected {
		t.Errorf("expected %d, got %d instead.\n", expected, actual)
	}
}

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	expected := 4
	actual := count(false, false, b)

	if actual != expected {
		t.Errorf("expected %d, got %d instead.\n", expected, actual)
	}
}

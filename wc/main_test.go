package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	expected := 4
	actual := count(b)

	if actual != expected {
		t.Errorf("expected %d, got %d instead.\n", expected, actual)
	}
}

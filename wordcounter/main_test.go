package main

import (
	"bytes"
	"testing"
)

func TestCountWord(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	res := count(b, false, false)
	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}

func TestCountLines(t *testing.T){
	b :=bytes.NewBufferString("word1 word2 word3 \n line2 \n line3 word1")
	exp := 3
	res := count(b, true, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}

func TestCountBytes(t *testing.T){
	b := bytes.NewBufferString("word1 word2 word3")
	exp := 17
	res := count(b, false, true)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
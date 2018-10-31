package main

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	orig := "marina de los angeles"
	tokens := Tokenize(orig)
	expected := []string{"marina", "de", "los", "angeles"}
	if !reflect.DeepEqual(expected, tokens) {
		t.Errorf("expected %v but got %v", expected, tokens)
	}
}

func TestTokenizeCaseInsensitive(t *testing.T) {
	orig := "marina DE lOs ANGELES"
	tokens := TokenizeLowerCase(orig)
	expected := []string{"marina", "de", "los", "angeles"}
	if !reflect.DeepEqual(expected, tokens) {
		t.Errorf("expected %v but got %v", expected, tokens)
	}
}

func TestTokenizeMultipleSpaces(t *testing.T) {
	orig := "marina       de    los   angeles"
	tokens := Tokenize(orig)
	expected := []string{"marina", "de", "los", "angeles"}
	if !reflect.DeepEqual(expected, tokens) {
		t.Errorf("expected %v but got %v", expected, tokens)
	}
}

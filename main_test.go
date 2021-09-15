package main

import "testing"

func TestOla(t *testing.T) {
	res := HelloWorld()
	expected := "hello world"

	if res != expected {
		t.Errorf("resultado '%s', esperado '%s'", res, expected)
	}
}

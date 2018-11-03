package main

import (
	"testing"
)

func TestGetUniversalMessage(t *testing.T) {
	if GetUniversalMessage() != "Hello world" {
		t.Error("did not return the expected message")
	} else {
		t.Log("Hello")
	}
}

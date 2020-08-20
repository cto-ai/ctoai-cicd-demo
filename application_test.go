package main

import (
	"fmt"
	"testing"
)

func TestTrueIsTrue(t *testing.T) {
	if true != true {
		t.Error()
	}

	fmt.Println("success")
}

func TestFalseIsFalse(t *testing.T) {
	if false != false {
		t.Error()
	}

	fmt.Println("success")
}

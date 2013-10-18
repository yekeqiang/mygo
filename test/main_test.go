package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	if sum(1, 2, 3) != 16 {
		t.Log("我")
		t.FailNow()
	}
}

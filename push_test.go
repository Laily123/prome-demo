package main

import "testing"

func TestPush(t *testing.T) {
	value := 10
	err := push(value)
	if err != nil {
		t.Error(err)
	}
}

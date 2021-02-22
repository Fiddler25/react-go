package main

import "testing"

func TestHello(t *testing.T) {
	result := GetHello("Kn")
	expect := "Hello Kin!"
	if result != expect {
		t.Error("\n正: ", expect, "\n誤: ", result)
	}

	t.Log("end")
}

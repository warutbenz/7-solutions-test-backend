package main

import "testing"

func TestDecodeLeftRightEquals(t *testing.T) {
	input := "LLRR="
	decode := decodeLeftRightEquals(input)
	if decode != "210122" {
		t.Error("TestDecodeLeftRightEquals should be 'Fail' because decode is ", decode)
	}
}

func TestEncodeLeftRightEquals(t *testing.T) {
	input := "410233"
	encode := encodeLeftRightEquals(input)
	if encode != "LLRR=" {
		t.Error("TestEncodeLeftRightEquals should be 'Fail' because encode is ", encode)
	}
}

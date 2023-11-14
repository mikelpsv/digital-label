package model

import "testing"

func TestEnc62_Encode(t *testing.T) {
	s := NewEnc62("").Encode(105456554)
	if s != "hiEek" {
		t.Error("base62 encode, invalid result")
	}
}

func TestEnc62_Decode(t *testing.T) {
	d := NewEnc62("").Decode("hiEek")
	if d != 105456554 {
		t.Error("base62 decode, invalid result")
	}
}

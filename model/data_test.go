package model_test

import (
	"testing"
	"wib-project/model"
)

func TestEncode(t *testing.T) {
	s := model.NewEnc62("").Encode(105456554)
	println(s)
}

func TestDecode(t *testing.T) {
	d := model.NewEnc62("").Decode("hiEek")
	println(d)
}

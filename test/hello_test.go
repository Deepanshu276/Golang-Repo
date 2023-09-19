package test

import (
	"testing"
	"github.com/Deepanshu276/Golang-Repo/Lect1"
)

func TestSayHello(t *testing.T){
	want:= "Hello , test!"
	got:=hello.Say("test")
	if want!=got{
		t.Errorf("wanted %s, got %s", want, got)
	}


}
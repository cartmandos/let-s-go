package hello

import "testing"

func TestSayHello(t *testing.T) {
	want:= "Hello, test!"
	got := Say("test")

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}

func TestSayV2Hello(t *testing.T) {
	want:= "Hello, test!"
	got := SayV2([]string{"test"})

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}
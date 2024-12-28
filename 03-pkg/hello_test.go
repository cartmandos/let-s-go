package hello

import "testing"

func TestSayV1Hello(t *testing.T) {
	want:= "Hello, test!"
	got := Say("test")

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}

func TestSayV2Hello(t *testing.T) {
	want := "Hello, test!"
	got := SayV2([]string{"test"})

	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}

func TestSayV2HelloCases(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items: []string{"Damien"},
			result: "Hello, Damien!",
		},
	}

	for _, st := range subtests {
		if s := SayV2(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.items,
				st.result, s)
		}
	}
}

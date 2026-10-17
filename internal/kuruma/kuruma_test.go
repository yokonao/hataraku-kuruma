package kuruma

import (
	"strings"
	"testing"
)

func TestArt(t *testing.T) {
	seen := map[string]bool{}
	for _, v := range All {
		if seen[v.Name] {
			t.Errorf("duplicate name %q", v.Name)
		}
		seen[v.Name] = true
		if len(v.Lights) == 0 {
			t.Errorf("%s: no lights", v.Name)
		}
		for _, l := range v.Art {
			for _, r := range l {
				if r > '~' || r < ' ' {
					t.Errorf("%s: non-ASCII %q in %q", v.Name, r, l)
				}
			}
		}
	}
}

func TestFrameClips(t *testing.T) {
	v := Vehicle{Color: red, Lights: []string{blue}, Art: []string{"ab(@)!", "%%%%"}}
	tests := []struct {
		n, x, width int
		want        []string
	}{
		{0, 0, 10, []string{"ab(-)*", "=-=-"}},
		{1, -2, 10, []string{"(\\)*", "-="}},
		{4, 7, 10, []string{"       ab(", "       =-="}},
		{0, 10, 10, []string{"          ", "          "}},
		{0, -6, 10, []string{"", ""}},
	}
	for _, tt := range tests {
		got := v.Frame(tt.n, tt.x, tt.width, false)
		if strings.Join(got, "\n") != strings.Join(tt.want, "\n") {
			t.Errorf("Frame(%d, %d, %d) = %q, want %q", tt.n, tt.x, tt.width, got, tt.want)
		}
	}
}

func TestFrameColor(t *testing.T) {
	v := Vehicle{Color: red, Lights: []string{blue}, Art: []string{"a !"}}
	got := v.Frame(0, 0, 10, true)[0]
	want := red + "a " + blue + "*" + reset
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestLookup(t *testing.T) {
	if _, ok := Lookup("firetruck"); !ok {
		t.Error("firetruck not found")
	}
	if _, ok := Lookup("tank"); ok {
		t.Error("tank found")
	}
}

func TestRandomNoRepeat(t *testing.T) {
	next := Random()
	prev := next()
	for range 1000 {
		v := next()
		if v.Name == prev.Name {
			t.Fatalf("%s twice in a row", v.Name)
		}
		prev = v
	}
}

func TestSmoke(t *testing.T) {
	v := Vehicle{Lights: []string{""}, Art: []string{"ab"}, Smoke: true, Exhaust: 1}
	got := v.Frame(0, 0, 100, false)
	if v.Height() != len(smoke[0])+1 || len(got) != v.Height() {
		t.Fatalf("Height() = %d, %d lines, want %d", v.Height(), len(got), len(smoke[0])+1)
	}
	if want := " " + smoke[0][0]; got[0] != want {
		t.Errorf("top row = %q, want %q", got[0], want)
	}
	if got[len(got)-1] != "ab" {
		t.Errorf("bottom row = %q, want %q", got[len(got)-1], "ab")
	}
}

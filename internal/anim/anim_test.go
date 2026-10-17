package anim

import (
	"context"
	"strings"
	"testing"

	"github.com/yokonao/hataraku-kuruma/internal/kuruma"
)

func TestRun(t *testing.T) {
	v := kuruma.Vehicle{Lights: []string{""}, Art: []string{"ab"}}
	var b strings.Builder
	if err := Run(context.Background(), &b, v, 3, 1, 1, false); err != nil {
		t.Fatal(err)
	}
	got := strings.Split(b.String(), "\x1b[1;1H\x1b[2K")
	want := []string{"\x1b[?1049h\x1b[?25l", "  a", " ab", "ab", "b\x1b[?25h\x1b[?1049l"}
	if strings.Join(got, "|") != strings.Join(want, "|") {
		t.Errorf("got %q, want %q", got, want)
	}
}

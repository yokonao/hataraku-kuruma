package anim

import (
	"context"
	"slices"
	"strings"
	"testing"

	"github.com/yokonao/hataraku-kuruma/internal/kuruma"
)

func TestRun(t *testing.T) {
	a := kuruma.Vehicle{Lights: []string{""}, Art: []string{"ab"}}
	b := kuruma.Vehicle{Lights: []string{""}, Art: []string{"c"}}
	var out strings.Builder
	if err := Run(context.Background(), &out, slices.Values([]kuruma.Vehicle{a, b}), 3, 1, 1, false); err != nil {
		t.Fatal(err)
	}
	got := strings.Split(out.String(), "\x1b[1;1H\x1b[2K")
	want := []string{"\x1b[?1049h\x1b[?25l\x1b[2J", "  a", " ab", "ab", "b\x1b[2J", "  c", " c", "c\x1b[?25h\x1b[?1049l"}
	if strings.Join(got, "|") != strings.Join(want, "|") {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRunStops(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	forever := func(yield func(kuruma.Vehicle) bool) {
		for {
			if !yield(kuruma.Vehicle{Lights: []string{""}, Art: []string{"ab"}}) {
				return
			}
		}
	}
	var out strings.Builder
	if err := Run(ctx, &out, forever, 3, 1, 1, false); err != nil {
		t.Fatal(err)
	}
}

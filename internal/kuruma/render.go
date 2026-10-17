package kuruma

import "strings"

const (
	reset  = "\x1b[0m"
	red    = "\x1b[31m"
	green  = "\x1b[32m"
	yellow = "\x1b[33m"
	blue   = "\x1b[34m"
	cyan   = "\x1b[36m"
	white  = "\x1b[37m"
)

// Vehicle is an ASCII art working vehicle. Art must be ASCII so that one
// byte is one terminal column.
type Vehicle struct {
	Name   string
	Label  string
	Color  string
	Lights []string
	Art    []string
}

func (v Vehicle) Width() int {
	w := 0
	for _, l := range v.Art {
		w = max(w, len(l))
	}
	return w
}

func (v Vehicle) Height() int { return len(v.Art) }

// Frame renders animation frame n with the vehicle's left edge at column x,
// clipped to columns [0, width).
func (v Vehicle) Frame(n, x, width int, color bool) []string {
	lights := n / 4 % 2
	lines := make([]string, len(v.Art))
	for row, art := range v.Art {
		var b strings.Builder
		cur := ""
		for i := range len(art) {
			col := x + i
			if col < 0 || col >= width {
				continue
			}
			ch, c := art[i], v.Color
			switch ch {
			case '@':
				ch = `-\|/`[n%4]
			case '%':
				ch = "=-"[(i+n)%2]
			case '!':
				ch = "*o"[lights]
				c = v.Lights[lights%len(v.Lights)]
			case ' ':
				c = cur
			}
			if color && c != cur {
				b.WriteString(c)
				cur = c
			}
			b.WriteByte(ch)
		}
		if color && cur != "" {
			b.WriteString(reset)
		}
		lines[row] = b.String()
	}
	return lines
}

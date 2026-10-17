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
	gray   = "\x1b[90m"
)

// Vehicle is an ASCII art working vehicle. Art must be ASCII so that one
// byte is one terminal column.
type Vehicle struct {
	Name   string
	Label  string
	Color  string
	Lights []string
	Art    []string
	// Smoke rises above the top row of Art from column Exhaust.
	Smoke   bool
	Exhaust int
}

func (v Vehicle) Width() int {
	w := 0
	for _, l := range v.sprite(0) {
		w = max(w, len(l))
	}
	return w
}

func (v Vehicle) Height() int { return len(v.sprite(0)) }

// sprite returns the rows of frame n: smoke, if any, then Art.
func (v Vehicle) sprite(n int) []string {
	if !v.Smoke {
		return v.Art
	}
	pad := strings.Repeat(" ", v.Exhaust)
	var rows []string
	for _, l := range smoke[n/4%len(smoke)] {
		rows = append(rows, pad+l)
	}
	return append(rows, v.Art...)
}

// Frame renders animation frame n with the vehicle's left edge at column x,
// clipped to columns [0, width). Lines start at column 0.
func (v Vehicle) Frame(n, x, width int, color bool) []string {
	lights := n / 4 % 2
	rows := v.sprite(n)
	smokeRows := len(rows) - len(v.Art)
	lines := make([]string, len(rows))
	for row, art := range rows {
		base := v.Color
		if row < smokeRows {
			base = gray
		}
		var b strings.Builder
		b.WriteString(strings.Repeat(" ", min(max(x, 0), width)))
		cur := ""
		for i := range len(art) {
			col := x + i
			if col < 0 || col >= width {
				continue
			}
			ch, c := art[i], base
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

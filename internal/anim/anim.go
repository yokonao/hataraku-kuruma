package anim

import (
	"context"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/yokonao/hataraku-kuruma/internal/kuruma"
)

// Run drives v from the right edge of a width x height screen until it has
// left the screen or ctx is done.
func Run(ctx context.Context, w io.Writer, v kuruma.Vehicle, width, height int, delay time.Duration, color bool) (err error) {
	// Alternate screen, hidden cursor; restored on return.
	if _, err := io.WriteString(w, "\x1b[?1049h\x1b[?25l"); err != nil {
		return err
	}
	defer func() {
		if _, rerr := io.WriteString(w, "\x1b[?25h\x1b[?1049l"); err == nil {
			err = rerr
		}
	}()

	top := max(0, (height-v.Height())/2)
	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for n, x := 0, width-1; x > -v.Width(); n, x = n+1, x-1 {
		var b strings.Builder
		for i, line := range v.Frame(n, x, width, color) {
			if top+i >= height {
				break
			}
			b.WriteString("\x1b[" + strconv.Itoa(top+i+1) + ";1H\x1b[2K" + line)
		}
		if _, err := io.WriteString(w, b.String()); err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
		}
	}
	return nil
}

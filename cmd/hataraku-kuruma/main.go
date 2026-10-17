package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/term"

	"github.com/yokonao/hataraku-kuruma/internal/anim"
	"github.com/yokonao/hataraku-kuruma/internal/kuruma"
)

// version is set at build time via -ldflags "-X main.version=...".
var version = "dev"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := newRootCmd().ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	var (
		kind    string
		list    bool
		loop    bool
		delay   time.Duration
		noColor bool
	)
	root := &cobra.Command{
		Use:     "hataraku-kuruma",
		Short:   "A random working vehicle drives across your terminal",
		Version: version,
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if list {
				for _, v := range kuruma.All {
					cmd.Printf("%-10s %s\n", v.Name, v.Label)
				}
				return nil
			}
			if delay <= 0 {
				return fmt.Errorf("--delay must be positive")
			}
			next := kuruma.Random()
			if kind != "" {
				v, ok := kuruma.Lookup(kind)
				if !ok {
					return fmt.Errorf("unknown vehicle %q (see --list)", kind)
				}
				next = func() kuruma.Vehicle { return v }
			}
			vehicles := func(yield func(kuruma.Vehicle) bool) {
				for {
					if !yield(next()) || !loop {
						return
					}
				}
			}
			fd := int(os.Stdout.Fd())
			width, height, err := term.GetSize(fd)
			if err != nil {
				width, height = 80, 24
			}
			color := !noColor && os.Getenv("NO_COLOR") == "" && term.IsTerminal(fd)
			return anim.Run(cmd.Context(), os.Stdout, vehicles, width, height, delay, color)
		},
	}
	root.Flags().StringVarP(&kind, "kind", "k", "", "vehicle to drive instead of a random one")
	root.Flags().BoolVarP(&list, "list", "l", false, "list the vehicles")
	root.Flags().BoolVar(&loop, "loop", false, "keep driving vehicles until interrupted")
	root.Flags().DurationVar(&delay, "delay", 40*time.Millisecond, "time per frame")
	root.Flags().BoolVar(&noColor, "no-color", false, "disable colors (also NO_COLOR)")
	return root
}

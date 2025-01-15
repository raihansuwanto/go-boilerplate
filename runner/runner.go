package runner

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Runner interface {
	Run(ctx context.Context) error
}

type RunnerFunc func(ctx context.Context) error

func (f RunnerFunc) Run(ctx context.Context) error {
	return f(ctx)
}

func Run(ctx context.Context, runners ...Runner) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	for _, runner := range runners {
		r := runner
		g.Go(func() error {
			return r.Run(ctx)
		})
	}

	return g.Wait()
}

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker(ctx context.Context, service string, sleep time.Duration) error {
	for {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				return fmt.Errorf("service %q was not available before the timeout", service)
			}

			return fmt.Errorf("context was cancelled: %w", ctx.Err())
		default:
			_, err := net.Dial("tcp", service)

			if err == nil {
				slog.Info(fmt.Sprintf("service %q is available", service))
				return nil
			}

			time.Sleep(sleep * time.Second)
		}
	}
}

func main() {
	timeoutOpt := flag.Int("timeout", 60, "How long, in seconds, should the program wait for each service to be available before timing out")
	sleepOpt := flag.Int("sleep", 1, "How long, in seconds, each worker waits between two consecutive polls of a given service")
	flag.Parse()

	args := flag.Args()
	services := make([]string, 0, len(args))
	services = append(services, args...)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeoutOpt)*time.Second)
	defer cancel()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	g, _ := errgroup.WithContext(ctx)

	for _, service := range services {
		g.Go(func() error {
			return worker(ctx, service, time.Duration(*sleepOpt))
		})
	}

	if err := g.Wait(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

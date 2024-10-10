package server

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/version", VersionHandler)
	mux.HandleFunc("/decode", DecodeHandler)
	mux.HandleFunc("/hard-op", HardOpHandler)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		fmt.Printf("Listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("err in listen: %s\n", err)
			return fmt.Errorf("failed to serve http server: %w", err)
		}
		fmt.Println("after listen http server")
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("failed to shutdown http server: %w", err)
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Printf("after wait: %s\n", err)
		return
	}
	fmt.Println("server shutdown gracefully")
}

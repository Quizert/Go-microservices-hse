package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand/v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type DecodeRequest struct {
	Input string `json:"input"`
}

type DecodeResponse struct {
	Output string `json:"output"`
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("version: 1.0.0")
		w.Write([]byte("version: 1.0.0"))
	default:
		fmt.Printf("Method not allowed\n")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var req DecodeRequest
		json.NewDecoder(r.Body).Decode(&req)
		fmt.Printf("Input string: %s\n", req.Input)
		decodeString, err := base64.StdEncoding.DecodeString(req.Input)
		if err != nil {
			fmt.Printf("Invalid base64 input\n")
			http.Error(w, "Invalid base64 input", http.StatusBadRequest)
			return
		}
		res := DecodeResponse{string(decodeString)}
		json.NewEncoder(w).Encode(res)
		fmt.Printf("Decoded string: %s\n", res.Output)
	default:
		fmt.Printf("Method not allowed\n")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func HardOpHandler(w http.ResponseWriter, r *http.Request) {
	duration := time.Duration(rand.IntN(10)+11) * time.Second
	time.Sleep(duration)

	if rand.IntN(2) == 0 {
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
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
			fmt.Printf("err in shutdown: %s\n", err)
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

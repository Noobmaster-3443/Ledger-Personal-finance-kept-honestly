package main

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

//go:embed index.html
var appHTML []byte

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		_, _ = io.WriteString(w, string(appHTML))
	})

	// Bind to an OS-assigned free port on loopback instead of a hardcoded
	// port, so a second instance (or anything else already using that port)
	// can't silently fail to start.
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalf("Ledger could not start: %v", err)
	}

	srv := &http.Server{Handler: mux}

	serveErrCh := make(chan error, 1)
	go func() {
		serveErrCh <- srv.Serve(listener)
	}()

	// Give the server a brief moment to fail fast before we trust it
	// enough to open a browser tab.
	select {
	case err := <-serveErrCh:
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ledger server stopped unexpectedly: %v", err)
		}
		return
	case <-time.After(150 * time.Millisecond):
	}

	url := fmt.Sprintf("http://%s/", listener.Addr().String())
	if err := openBrowser(url); err != nil {
		fmt.Println("Couldn't open a browser automatically. Open this URL manually:")
	}
	fmt.Println("Ledger running at", url)
	fmt.Println("Press Ctrl+C to stop.")

	// Wait for Ctrl+C / termination signal, then shut down cleanly instead
	// of blocking forever with select{}.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stop:
	case err := <-serveErrCh:
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ledger server stopped unexpectedly: %v", err)
		}
		return
	}

	fmt.Println("\nShutting down…")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Shutdown error: %v", err)
	}
}

func openBrowser(url string) error {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return exec.Command("xdg-open", url).Start()
	}
}

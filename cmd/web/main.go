package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	// Initialize a new instance of application containing the dependencies
	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Serve static files using the file server
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register routes
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.viewSnippet)
	mux.HandleFunc("GET /snippet/create", app.createSnippet)
	mux.HandleFunc("POST /snippet/create", app.createSnippetPost)

	logger.Info("starting server", slog.String("addr", *addr))

	err := http.ListenAndServe(*addr, mux)

	if err != nil {
		logger.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}

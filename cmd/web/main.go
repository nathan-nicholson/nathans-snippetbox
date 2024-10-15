package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"snippetbox.nathan-r-nicholson.com/internal/models"
)

type application struct {
	logger        *slog.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	dbUser := flag.String("dbuser", "web", "MySQL database user")
	dbPass := flag.String("dbpass", "Snippets4Days", "MySQL database password")
	dbHost := flag.String("dbhost", "my-snippetbox-mysql.my-snippetbox.svc.cluster.local", "MySQL database host")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	db, err := openDB(fmt.Sprintf("%s:%s@tcp(%s:3306)/snippetbox?parseTime=true", *dbUser, *dbPass, *dbHost))

	if err != nil {
		logger.Error("cannot connect to database", "error", err)
		os.Exit(1)
	}

	defer db.Close()

	// Initialize a new template cache
	templateCache, err := newTemplateCache()

	if err != nil {
		logger.Error("cannot create template cache", "error", err)
		os.Exit(1)
	}

	// Initialize a new instance of application containing the dependencies
	app := &application{
		logger:        logger,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	mux := app.routes()

	logger.Info("starting server", slog.String("addr", *addr))

	err = http.ListenAndServe(*addr, mux)

	if err != nil {
		logger.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

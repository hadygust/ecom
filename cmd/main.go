package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/hadygust/ecom/internal/env"
	"github.com/jackc/pgx/v5"
)

// func main() {
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)
// 	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("welcome"))
// 	})
// 	http.ListenAndServe(":3000", r)
// }

func main() {
	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres dbname=ecommerce  password=janganlupa123 sslmode=disable"),
		},
	}

	// Logger

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Database

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	logger.Info("connected to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     *conn,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("server has failed to start", "error", err)
		os.Exit(1)
	}

}

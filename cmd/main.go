package main

import (
	"log/slog"
	"os"
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

	cfg := config{
		addr: ":8080",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	// Logger

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		slog.Error("server has failed to start", "error", err)
		os.Exit(1)
	}

}

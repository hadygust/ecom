package main

import (
	"log"
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

	if err := api.run(api.mount()); err != nil {
		log.Printf("server has failed to start, err: %s", err)
		os.Exit(1)
	}

}

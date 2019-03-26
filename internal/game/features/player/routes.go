package player

import "github.com/go-chi/chi"

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{playerID}", GetPlayer)
	return router
}

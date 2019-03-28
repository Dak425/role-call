package player

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{playerID}", getPlayer)
	r.Delete("/{playerID}", deletePlayer)
	r.Post("/", createPlayer)
	r.Get("/", getAllPlayers)
	return r
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{
		"feature": "getPlayer",
		"status":  "not yet implemented",
	})
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{
		"feature": "deletePlayer",
		"status":  "not yet implemented",
	})
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{
		"feature": "createPlayer",
		"status":  "not yet implemented",
	})
}

func getAllPlayers(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{
		"feature": "getAllPlayers",
		"status":  "not yet implemented",
	})
}

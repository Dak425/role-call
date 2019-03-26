package player

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

var players = map[string]Player{
	"1": {ID: "1", RobloxID: "555555"},
	"2": {ID: "2", RobloxID: "666666"},
	"3": {ID: "3", RobloxID: "777777"},
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	playerID := chi.URLParam(r, "playerID")
	player := players[playerID]
	render.JSON(w, r, player)
}

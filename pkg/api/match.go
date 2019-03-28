package api

import "time"

type Match struct {
	ID       int           `json:"id"`
	Duration time.Duration `json:"duration"`
	Players  []Player      `json:"players"`
}

type MatchRepository interface {
	Match(id int) (match *Match, err error)
	Matches() (matches []*Match, err error)
	UpsertMatch(m *Match) (err error)
	DeleteMatch(id int) (err error)
}

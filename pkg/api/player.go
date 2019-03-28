package api

type Player struct {
	ID       int `json:"id"`
	RobloxID int `json:"roblox_id"`
}

type PlayerRepository interface {
	Player(id int) (*Player, error)
	Players() ([]*Player, error)
	UpsertPlayer(*Player) error
	DeletePlayer(id int) error
}

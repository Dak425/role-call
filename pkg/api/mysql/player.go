package mysql

import (
	"database/sql"
	"github.com/Dak425/role-call/pkg/api"
	_ "github.com/go-sql-driver/mysql"
)

type PlayerRepository struct {
	DB *sql.DB
}

func (pr *PlayerRepository) Player(id int) (player *api.Player, err error) {
	query, err := pr.DB.Prepare(`SELECT * FROM Player WHERE id = ?`)

	if err != nil {
		return
	}

	row := query.QueryRow(id)

	defer query.Close()

	err = row.Scan(&player.ID, &player.RobloxID)

	if err != nil {
		player = nil
	}

	return
}

func (pr *PlayerRepository) Players() (players []*api.Player, err error) {
	result, err := pr.DB.Query(`SELECT * FROM Players`)

	if err != nil {
		return
	}

	defer result.Close()

	for result.Next() {
		var player *api.Player
		err = result.Scan(&player.ID, &player.RobloxID)

		if err != nil {
			players = nil
			return
		}

		players = append(players, player)
	}
	return
}

func (pr *PlayerRepository) UpsertPlayer(player *api.Player) error {
	// If no id, player is new, insert into table
	// If id is present, update player's data
	if player.ID == 0 {
		query, err := pr.DB.Prepare("INSERT INTO Players (RobloxID) VALUES (?)")

		if err != nil {
			return err
		}

		defer query.Close()

		_, err = query.Exec(player.RobloxID)

		return err
	} else {
		query, err := pr.DB.Prepare("UPDATE Players SET RobloxID = ? WHERE ID = ?")

		if err != nil {
			return err
		}

		defer query.Close()

		_, err = query.Exec(player.RobloxID, player.ID)

		return err
	}
}

func (pr *PlayerRepository) DeletePlayer(id int) error {
	query, err := pr.DB.Prepare("DELETE FROM Players WHERE ID = ?")

	if err != nil {
		return err
	}

	_ , err = query.Exec(id)

	return err
}
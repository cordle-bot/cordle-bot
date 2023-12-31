package users

import "fmt"

type User struct {
	Id     string
	Wins   int
	Losses int
	Draws  int
	Elo    int
}

func (u User) ToStr() string {
	return fmt.Sprintf(
		"Id: %s\nLosses: %d\nLosses: %d\nDraws: %d\nGames Played: %d\nElo: %d",
		u.Id,
		u.Wins,
		u.Losses,
		u.Draws,
		u.Wins+u.Losses+u.Draws,
		u.Elo,
	)
}

func (u User) ToSqlAdd() string {
	return fmt.Sprintf(
		"%s, %d, %d, %d, %d",
		u.Id,
		u.Wins,
		u.Losses,
		u.Draws,
		u.Elo,
	)
}

func (u User) ToSqlUpdate() string {
	return fmt.Sprintf(
		"wins=%d, losses=%d, draws=%d, elo=%d",
		u.Wins,
		u.Losses,
		u.Draws,
		u.Elo,
	)
}

func (u User) ToStat() string {
	return fmt.Sprintf(
		"Wins: %d\nLosses: %d\nDraws: %d\nGames: %d\nElo: %d",
		u.Wins,
		u.Losses,
		u.Draws,
		u.Wins+u.Losses+u.Draws,
		u.Elo,
	)
}

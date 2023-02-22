package game

import (
	"cordle/wordle"

	"github.com/bwmarrin/discordgo"
)

// DuelGame holds the information about a DuelGame
type DuelGame struct {
	// games stores a map of user IDs to their game
	games map[string]*wordle.WordleGame
	// menus stores the interaction to edit to display games to each user
	menus map[string]*discordgo.Interaction
}

// NewDuelGame creates a specialized Game struct representing a Cordle Duel Game
func NewDuelGame(th string, p []*discordgo.User) {
	// Create the shared game
	g0 := wordle.NewRandomGame()
	// Manually create a second game with the same goal word
	// This is more efficient than doing a deep copy
	g1 := &wordle.WordleGame{
		Guesses:  []*wordle.Guess{},
		GoalWord: g0.GoalWord,
	}

	// Create the game struct and store it
	games.mu.Lock()
	games.g[th] = &DuelGame{
		games: map[string]*wordle.WordleGame{
			p[0].ID: g0,
			p[1].ID: g1,
		},
		menus: make(map[string]*discordgo.Interaction),
	}
	games.mu.Unlock()
}

// PlayerInGame returns true if the given player is a part of the game
func (g *DuelGame) PlayerInGame(p *discordgo.User) bool {
	_, exists := g.games[p.ID]
	return exists
}

// PlayerHasGuesses returns true if the player has guesses remaining in the game
func (g *DuelGame) PlayerHasGuesses(p *discordgo.User) bool {
	return g.games[p.ID].GuessesRemaining() > 0
}

// GetPlayerInteractionMenu searches for and returns the interaction menu for the given player
// Returns a boolean to indicate whether or not the menu was found
func (g *DuelGame) GetPlayerInteractionMenu(p *discordgo.User) (*discordgo.Interaction, bool) {
	r, exists := g.menus[p.ID]
	return r, exists
}

// SetPlayerInteractionMenu stores an interaction to be used as the user's menu
func (g *DuelGame) SetPlayerInteractionMenu(p *discordgo.User, m *discordgo.Interaction) {
	g.menus[p.ID] = m
}

// SubmitGuess allows a guess to be submitted to the game of a given player
// Returns the result as an array of wordle.GuessState
func (g *DuelGame) SubmitGuess(guess string, p *discordgo.User) (*wordle.Guess, error) {
	pg := g.games[p.ID]
	return pg.Guess(guess)
}

// PlayerGuessHistory returns the guess history of the given player
func (g *DuelGame) PlayerGuessHistory(p *discordgo.User) []*wordle.Guess {
	return g.games[p.ID].Guesses
}

// GoalWord returns the goal word for this game
func (g *DuelGame) GoalWord(p *discordgo.User) string {
	return g.games[p.ID].GoalWord
}

// PlayerSurrender allows a player to quit an ongoing game
// In a duel game, this should immediately end the game
func (g *DuelGame) PlayerSurrender(p *discordgo.User) {
	for id, g := range g.games {
		if p.ID != id {
			g.Won = true
			return
		}
	}
}

// GameWon returns true if the game has been won, as well as the ID of the winner
func (g *DuelGame) GameWon() (bool, string) {
	for id, g := range g.games {
		if g.Won {
			return true, id
		}
	}
	return false, ""
}

// ShouldEndInDraw returns true if the current game has reached a stalemate and should end in a draw
func (g *DuelGame) ShouldEndInDraw() bool {
	for _, g := range g.games {
		if g.GuessesRemaining() > 0 {
			return false
		}
	}
	return true
}

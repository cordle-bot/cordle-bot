//go:build unit

package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElo(t *testing.T) {
	// Create test elo score values (should be arbitrary)
	// Assuming a correct elo formula, the sums of elo should remain constant (zero-sum)
	ra := 1000
	rb := 1050
	// Test a win
	eloa := calculateElo(ra, rb, ScoreWin)
	elob := calculateElo(rb, ra, ScoreLoss)
	assert.Equal(t, eloa+elob, ra+rb)
	// Test a draw
	eloa = calculateElo(ra, rb, ScoreDraw)
	elob = calculateElo(rb, ra, ScoreDraw)
	assert.Equal(t, eloa+elob, ra+rb)
}

package tictactoe

import "boardgames/board"

// PlayerWins erwartet ein Spielfeld und ein Zeichen,
// das einen Spieler repräsentiert (meist "X" oder "O").
// Liefert true, falls dieser Spieler gewonnen hat.
func PlayerWins(b board.Board, player string) bool {
	result := false
	for i := range b {
		if b.RowContainsOnly(i, player) {
			result = true
		} else if b.ColumnContainsOnly(i, player) {
			result = true
		} else if b.DiagDownRightContainsOnly(player) {
			result = true
		} else if b.DiagUpRightContainsOnly(player) {
			result = true
		}

	}
	return result
}

// PlayerAllowed erwartet ein Spielfeld und eine Position
// in Form einer Zeilen- und einer Spaltennummer.
// Liefert true, falls auf dieses Feld ein Zug erlaubt ist.
// Ein Zug ist erlaubt, wenn das Feld leer ist
// und wenn die Position gültig ist.
func PlayerAllowed(b board.Board, row, col int) bool {
	// TODO
	return false
}

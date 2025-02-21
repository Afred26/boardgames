package queens

import "boardgames/board"

// QueenAllowed erwartet ein Spielfeld und eine Position.
// Liefert true, falls es erlaubt ist, dort eine Dame zu setzen.
func QueenAllowed(b board.Board, row, col int) bool {

	return b.RowEmpty(row) && b.ColumnEmpty(col) && b.DiagDownRightEmpty(col-row) && b.DiagUpRightEmpty(col+row-len(b)+1)
}

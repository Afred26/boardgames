package board

// RowContainsOnly erwartet eine Zeilennummer row und einen String s.
// Liefert true, wenn die row-te Zeile nur aus dem gegebenen String s besteht.
func (board Board) RowContainsOnly(row int, s string) bool {

	return ContainsOnly(board.GetRow(row), s)
}

// ColumnContainsOnly erwartet eine Spaltennummer col und einen String s.
// Liefert true, wenn die col-te Spalte nur aus dem gegebenen String s besteht.
func (board Board) ColumnContainsOnly(col int, s string) bool {

	return ContainsOnly(board.GetColumn(col), s)
}

// DiagDownRightContainsOnly erwartet einen String s.
// Liefert true, wenn die Diagonale, die in Spalte 0 und Zeile 0 beginnt,
// nur aus dem gegebenen String s besteht.
func (board Board) DiagDownRightContainsOnly(s string) bool {

	return ContainsOnly(board.GetDiagDownRight(0), s)
}

// DiagUpRightContainsOnly erwartet einen String s.
// Liefert true, wenn die Diagonale, die in Spalte 0 und der letzten Zeile beginnt,
// nur aus dem gegebenen String s besteht.
func (board Board) DiagUpRightContainsOnly(s string) bool {

	return ContainsOnly(board.GetDiagUpRight(0), s)
}

// RowEmpty erwartet ein Spielfeld und eine Zeilennummer row.
// Liefert true, wenn die row-te Zeile leer ist, sonst false.
// Liefert true, wenn die Zeile nicht existiert.
func (board Board) RowEmpty(row int) bool {
	if row < 0 || row >= len(board) {
		return true
	}

	return ContainsOnly(board.GetRow(row), " ")
}

// ColumnEmpty erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, wenn die col-te Spalte leer ist, sonst false.
// Liefert true, wenn die Spalte nicht existiert.
func (board Board) ColumnEmpty(col int) bool {
	if col < 0 || col >= len(board) {
		return true
	}

	return ContainsOnly(board.GetColumn(col), " ")
}

// DiagDownRightEmpty erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die Diagonale, die in Spalte col und Zeile 0 beginnt,
// nur Leerzeichen enthält.
// Bei ungültigen Spaltennummern wird ggf. eine Teil-Diagonale betrachtet.
func (board Board) DiagDownRightEmpty(col int) bool {

	return ContainsOnly(board.GetDiagDownRight(col), " ")
}

// DiagUpRightEmpty erwartet ein Spielfeld und eine Spaltennummer col.
// Liefert true, falls die Diagonale, die in Spalte col und der letzten Zeile beginnt,
// nur Leerzeichen enthält.
// Bei ungültigen Spaltennummern wird ggf. eine Teil-Diagonale betrachtet.
func (board Board) DiagUpRightEmpty(col int) bool {
	return ContainsOnly(board.GetDiagUpRight(col), " ")
}

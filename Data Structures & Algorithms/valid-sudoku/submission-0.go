func isValidSudoku(board [][]byte) bool {
    rows := make([]int, 9)
    cols := make([]int, 9)
    squares := make([]int, 9)
    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {
            if board[row][col] == '.' {
                continue
            }
            // process num into bit
            val := board[row][col] - '1'
            bit := 1 << val
            squareIdx := (row/3) * 3 + (col/3)
            // check if bit taken with & bit
            if cols[col]&bit != 0 || rows[row]&bit != 0 || squares[squareIdx]&bit != 0 {
                return false
            }
            // update
            rows[row] |= bit
            cols[col] |= bit
            squares[squareIdx] |= bit
        }
    }
    return true
}

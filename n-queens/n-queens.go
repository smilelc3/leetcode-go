package n_queens

import "strings"

func solveNQueens(n int) [][]string {
	var ans [][]string
	var chessboard []string
	for i := 0; i < n; i++ {
		builder := strings.Builder{}
		for j := 0; j < n; j++ {
			builder.WriteByte('.')
		}
		chessboard = append(chessboard, builder.String())
	}
	NQueuesBackTrack(0, n, 0, [2]uint32{0, 0}, chessboard, &ans)

	return ans
}

func NQueuesBackTrack(row int, n int, colExistQueue uint16, slashExistQueue [2]uint32, cur []string, pSave *[][]string) {
	// 1 <= n <= 9
	// slashExistQueue[0]  means /
	// slashExistQueue[0]  means \
	for col := 0; col < n; col++ {
		if colExistQueue&(1<<col) != 0 ||
			slashExistQueue[0]&(1<<(row+col)) != 0 ||
			slashExistQueue[1]&(1<<(row-col+n-1)) != 0 {
			continue
		} else {
			if row+1 == n {
				back := make([]string, n)
				copy(back, cur)
				tempBytes := []byte(back[row])
				tempBytes[col] = 'Q'
				back[row] = string(tempBytes)
				*pSave = append(*pSave, back)
				return
			}
			colExistQueue |= 1 << col
			slashExistQueue[0] |= 1 << (row + col)
			slashExistQueue[1] |= 1 << (row - col + n - 1)
			tempBytes := []byte(cur[row])
			tempBytes[col] = 'Q'
			cur[row] = string(tempBytes)

			NQueuesBackTrack(row+1, n, colExistQueue, slashExistQueue, cur, pSave)

			// resume
			tempBytes = []byte(cur[row])
			tempBytes[col] = '.'
			cur[row] = string(tempBytes)
			colExistQueue = colExistQueue ^ (1 << col)
			slashExistQueue[0] ^= 1 << (row + col)
			slashExistQueue[1] ^= 1 << (row - col + n - 1)
		}
	}
}

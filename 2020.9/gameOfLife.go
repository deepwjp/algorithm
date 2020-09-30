package main

//所有细胞同事更新，判断初始细胞周围状态
func gameOfLife(board [][]int) {
	//1 原来活现在活
	//0 原来死现在死
	//2 原来死现在活
	//-1 原来活现在死
	for i, rowArr := range board {
		for j, _ := range rowArr {
			//获取初始状态当前细胞周围存活的细胞个数
			temp := isAliveGameOfLife(i-1, j-1, board) + isAliveGameOfLife(i-1, j, board) + isAliveGameOfLife(i-1, j+1, board) + isAliveGameOfLife(i, j-1, board) + isAliveGameOfLife(i, j+1, board) + isAliveGameOfLife(i+1, j-1, board) + isAliveGameOfLife(i+1, j, board) + isAliveGameOfLife(i+1, j+1, board)
			if temp < 2 && board[i][j] == 1 {
				board[i][j] = -1
			} else if temp > 3 && board[i][j] == 1 {
				board[i][j] = -1
			} else if temp == 3 && board[i][j] == 0 {
				board[i][j] = 2
			}
		}
	}
	//更新切片
	for i, rowArr := range board {
		for j, _ := range rowArr {
			//将现在活着的置为1.否则置为0
			if board[i][j] == 1 || board[i][j] == 2 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}
func isAliveGameOfLife(i int, j int, board [][]int) int {
	//判断数组是否越界
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return 0
	}
	//判断初始状态是否活着
	if board[i][j] == 1 || board[i][j] == -1 {
		return 1
	}
	return 0
}

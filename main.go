package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	boardSize = 16  // 棋盘大小
	emptyCell = "-" // 空棋子
	playerX   = "X" // 玩家X的棋子
	playerO   = "O" // 玩家O的棋子
)

func main() {
	// 创建新的实例
	game := NewGame()
	// 创建扫描器用于读取标准输入
	scanner := bufio.NewScanner(os.Stdin)

	game.saveBoard() // 保存棋盘状态

	for {
		game.printBoard() // 打印当前棋盘状态
		fmt.Printf("玩家%s的回合，请输入行和列 7 7 : ", game.player)

		var row, col int
		if scanner.Scan() {
			_, err := fmt.Sscanf(scanner.Text(), "%d %d", &row, &col) // 从扫描结果中解析出行和列
			if err != nil {
				fmt.Println("无效的输入，请输入行和列数。")
				continue
			}
		}

		if game.makeMove(row, col) { // 进行移动操作
			if game.checkWin(row, col) { // 检查是否获胜
				game.printBoard() // 打印棋盘状态
				fmt.Printf("玩家%s获胜！\n", game.player)
				break // 结束游戏
			}

			game.switchPlayer() // 切换玩家
		} else {
			fmt.Println("无效的移动，请重试。")
		}
	}
}

type Module struct {
	// 棋盘数组
	board  [boardSize][boardSize]string
	player string // 当前玩家
}

func NewGame() *Module {
	// 创建新实例
	game := &Module{player: playerX}
	for i := 1; i < boardSize; i++ {
		for j := 1; j < boardSize; j++ {
			// 初始化棋盘
			game.board[i][j] = emptyCell
		}
	}
	return game
}

func (g *Module) saveBoard() {
	// 连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/module")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 执行插入操作
	insertStmt, err := db.Prepare("INSERT INTO board (board_info) VALUES (?)")
	if err != nil {
		panic(err.Error())
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec("value1") // 执行插入操作
	if err != nil {
		panic(err.Error())
	}
}

func (g *Module) printBoard() {
	fmt.Print("  ")
	for i := 1; i < boardSize; i++ {
		// 打印横行
		fmt.Printf("%2d ", i)
	}
	fmt.Println()

	for i := 1; i < boardSize; i++ {
		// 打印纵列
		fmt.Printf("%2d ", i)
		for j := 1; j < boardSize; j++ {
			// 打印棋盘上的棋子
			fmt.Printf("%2s ", g.board[i][j])
		}
		fmt.Println()
	}
}

func (g *Module) makeMove(row, col int) bool {
	// 判断移动是否有效
	if row < 1 || row >= boardSize || col < 1 || col >= boardSize || g.board[row][col] != emptyCell {
		return false
	}
	// 在棋盘上放置棋子
	g.board[row][col] = g.player
	return true
}

func (g *Module) switchPlayer() {
	if g.player == playerX {
		g.player = playerO
	} else {
		g.player = playerX
	}
}

func (g *Module) checkWin(row, col int) bool {
	// 校验落子位置
	return g.checkDirection(row, col, 0, 1) || // 横向
		g.checkDirection(row, col, 1, 0) || // 纵向
		g.checkDirection(row, col, 1, 1) || // 斜向 \
		g.checkDirection(row, col, 1, -1) // 斜向 /
}

func (g *Module) checkDirection(row, col, dr, dc int) bool {
	count := 1
	for i := 1; i <= 4; i++ {
		r := row + i*dr
		c := col + i*dc
		if r < 1 || r >= boardSize || c < 1 || c >= boardSize || g.board[row][col] != g.board[r][c] {
			break
		}
		count++
	}

	for i := 1; i <= 4; i++ {
		r := row - i*dr
		c := col - i*dc

		if r < 1 || r >= boardSize || c < 1 || c >= boardSize || g.board[row][col] != g.board[r][c] {
			break
		}

		count++
	}

	return count >= 5
}

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	Empty     = 0
	Player1   = 1
	Player2   = 2
	BoardSize = 15
)

type Gobang struct {
	board       [][]int
	currentTurn int
}

func NewGobang() *Gobang {
	board := make([][]int, BoardSize)
	for i := range board {
		board[i] = make([]int, BoardSize)
	}
	return &Gobang{
		board:       board,
		currentTurn: Player1,
	}
}

func (g *Gobang) PrintBoard() {
	for _, row := range g.board {
		for _, cell := range row {
			switch cell {
			case Empty:
				fmt.Print("- ")
			case Player1:
				fmt.Print("X ")
			case Player2:
				fmt.Print("O ")
			}
		}
		fmt.Println()
	}
}

func (g *Gobang) MakeMove(row, col int) bool {
	if row < 0 || row >= BoardSize || col < 0 || col >= BoardSize || g.board[row][col] != Empty {
		return false
	}

	g.board[row][col] = g.currentTurn
	g.currentTurn = 3 - g.currentTurn
	return true
}

func (g *Gobang) IsGameOver() bool {
	return g.checkWin() || g.isBoardFull()
}

func (g *Gobang) checkWin() bool {
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if g.board[row][col] != Empty {
				if g.checkDirection(row, col, 1, 0) || // 水平方向
					g.checkDirection(row, col, 0, 1) || // 垂直方向
					g.checkDirection(row, col, 1, 1) || // 正斜方向
					g.checkDirection(row, col, 1, -1) { // 反斜方向
					return true
				}
			}
		}
	}
	return false
}

func (g *Gobang) checkDirection(row, col, dx, dy int) bool {
	player := g.board[row][col]
	count := 1

	for i := 1; i <= 4; i++ {
		nr := row + dx*i
		nc := col + dy*i

		if nr < 0 || nr >= BoardSize || nc < 0 || nc >= BoardSize || g.board[nr][nc] != player {
			break
		}

		count++
	}

	return count == 5
}

func (g *Gobang) isBoardFull() bool {
	for _, row := range g.board {
		for _, cell := range row {
			if cell == Empty {
				return false
			}
		}
	}
	return true
}

func main() {
	g := NewGobang()
	rand.Seed(time.Now().UnixNano())

	// 隐藏模式选择
	//fmt.Println("欢迎来到Gobang游戏！")
	//fmt.Println("请选择游戏模式：")
	//fmt.Println("1. 玩家对战")
	//fmt.Println("2. 人机对战")
	//var mode int
	//fmt.Scanln(&mode)

	mode := 2
	switch mode {
	case 1:
		fmt.Println("玩家对战模式")
		for !g.IsGameOver() {
			g.PrintBoard()
			fmt.Printf("轮到玩家%d下棋，请输入行和列号（以空格分隔）：", g.currentTurn)
			var row, col int
			fmt.Scanln(&row, &col)
			if !g.MakeMove(row, col) {
				fmt.Println("无效的位置，请重新输入！")
			}
		}
	case 2:
		fmt.Println("人机对战模式")
		for !g.IsGameOver() {
			g.PrintBoard()
			if g.currentTurn == Player1 {
				fmt.Println("轮到AI下棋...")
				time.Sleep(1 * time.Second) // 模拟思考时间

				row, col := g.generateAIMove()
				g.MakeMove(row, col)
				fmt.Printf("AI下棋在行%d，列%d\n", row, col)
			} else {
				fmt.Printf("轮到玩家%d下棋，请输入行和列号（以空格分隔）：", g.currentTurn)
				var row, col int
				fmt.Scanln(&row, &col)
				if !g.MakeMove(row, col) {
					fmt.Println("无效的位置，请重新输入！")
				}
			}
		}
	}

	g.PrintBoard()
	fmt.Println("游戏结束！")
}

func (g *Gobang) generateAIMove() (int, int) {
	advantageMap := make(map[int]int) // 记录每个空位的权重

	// 遍历棋盘上的每个空位
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if g.board[row][col] == Empty {
				advantage := g.calculateAdvantage(row, col) // 计算权重
				advantageMap[row*BoardSize+col] = advantage
			}
		}
	}

	// 根据权重降序排序
	sortedPositions := make([]int, 0, len(advantageMap))
	for position := range advantageMap {
		sortedPositions = append(sortedPositions, position)
	}
	sort.Slice(sortedPositions, func(i, j int) bool {
		return advantageMap[sortedPositions[i]] > advantageMap[sortedPositions[j]]
	})

	// 选择权重最高的位置进行下棋
	for _, position := range sortedPositions {
		row := position / BoardSize
		col := position % BoardSize
		if g.board[row][col] == Empty {
			return row, col
		}
	}

	// 如果没有合适的位置，则随机选择一个空位置下棋
	return g.getRandomEmptyPosition()
}

func (g *Gobang) calculateAdvantage(row, col int) int {
	advantage := 0

	// 遍历附近三格
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			nr := row + dr
			nc := col + dc

			// 判断是否在边界内
			if nr >= 0 && nr < BoardSize && nc >= 0 && nc < BoardSize {
				// 判断是否为敌方棋子
				if g.board[nr][nc] == 3-g.currentTurn {
					advantage++
				}
			}
		}
	}

	return advantage
}

func (g *Gobang) getRandomEmptyPosition() (int, int) {
	for {
		row := rand.Intn(BoardSize)
		col := rand.Intn(BoardSize)
		if g.board[row][col] == Empty {
			return row, col
		}
	}
}

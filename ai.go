package main

import (
	"log"
	"math/rand"
	"sort"
)


func (g *Gobang) generateAIMove() (int, int) {
	advantageMap := make(map[int]int) // 记录每个空位的权重

	// 遍历棋盘上的每个空位
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			// 计算空位权重
			if g.board[row][col] == Empty {
				advantage := g.calculateAdvantage(row, col)
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
		}else {
			log.Println("position:{} is not empty,{},{}",position,row,col)
		}
	}

	// 如果没有合适的位置，则随机选择一个空位置下棋
	return g.getRandomEmptyPosition()
}

func (g *Gobang) calculateAdvantage(row, col int) int {
	advantage := 0

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
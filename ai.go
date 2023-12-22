package main


//func (g *Gobang) GenerateAIMove() (int, int) {
//	advantageMap := make(map[int]int) // 记录每个空位的权重
//
//	// 遍历棋盘上的每个空位
//	for row := 0; row < BoardSize; row++ {
//		for col := 0; col < BoardSize; col++ {
//			// 计算空位权重
//			if g.board[row][col] == Empty {
//				advantage := g.CalculateAdvantage(row, col)
//				advantageMap[row*BoardSize+col] = advantage
//			}
//		}
//	}
//
//	// 根据权重降序排序
//	sortedPositions := make([]int, 0, len(advantageMap))
//	for position := range advantageMap {
//		sortedPositions = append(sortedPositions, position)
//	}
//	sort.Slice(sortedPositions, func(i, j int) bool {
//		return advantageMap[sortedPositions[i]] > advantageMap[sortedPositions[j]]
//	})
//
//	// 选择权重最高的位置进行下棋
//	for _, position := range sortedPositions {
//		row := position / BoardSize
//		col := position % BoardSize
//		if g.board[row][col] == Empty {
//			return row, col
//		}else {
//			log.Println("position:{} is not empty,{},{}",position,row,col)
//		}
//	}
//
//	// 如果没有合适的位置，则随机选择一个空位置下棋
//	return g.GetRandomEmptyPosition()
//}
//
//func (g *Gobang) GetRandomEmptyPosition() (int, int) {
//	for {
//		row := rand.Intn(BoardSize)
//		col := rand.Intn(BoardSize)
//		if g.board[row][col] == Empty {
//			return row, col
//		}
//	}
//}
//
//func (g *Gobang) CalculateAdvantage(row, col int) int {
//	score := 0
//
//	// 遍历八个方向
//	for dr := -1; dr <= 1; dr++ {
//		for dc := -1; dc <= 1; dc++ {
//			if dr == 0 && dc == 0 {
//				continue
//			}
//
//			nr := row + dr
//			nc := col + dc
//
//			// 判断是否在边界内
//			if nr >= 0 && nr < BoardSize && nc >= 0 && nc < BoardSize {
//				// 判断当前方向上的局势评分
//				directionScore := g.CalculateDirectionScore(row, col, dr, dc)
//				score += directionScore
//			}
//		}
//	}
//
//	return score
//}
//
//func (g *Gobang) CalculateDirectionScore(row, col, dr, dc int) int {
//	enemy := 3 - g.currentTurn
//	score := 0
//
//	// 判断当前位置是否为敌方棋子
//	if g.board[row][col] == enemy {
//		return score
//	}
//
//	// 计算当前方向上的连续棋子数量
//	count := 0
//	emptyCount := 0
//	for i := 1; i <= 4; i++ {
//		nr := row + i*dr
//		nc := col + i*dc
//
//		// 判断是否在边界内
//		if nr < 0 || nr >= BoardSize || nc < 0 || nc >= BoardSize {
//			break
//		}
//
//		if g.board[nr][nc] == g.currentTurn {
//			count++
//		} else if g.board[nr][nc] == enemy {
//			break
//		} else {
//			emptyCount++
//		}
//	}
//
//	// 根据连续棋子数量和空位数量计算局势评分
//	switch count {
//	case 0:
//		switch emptyCount {
//		case 0:
//			score += 10 // 一个子
//		case 1:
//			score += 20 // 死二
//		case 2:
//			score += 50 // 冲二
//		case 3:
//			score += 100 // 活二
//		}
//	case 1:
//		switch emptyCount {
//		case 0:
//			score += 30 // 死三
//		case 1:
//			score += 100 // 冲三
//		case 2:
//			score += 1000 // 活三
//		}
//	}
//
//	return score
//}
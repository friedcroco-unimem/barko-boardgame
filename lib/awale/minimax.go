package awale

const minimaxDepth = 5

func (b *Board) GetBestMove() *Move {
	move, _ := minimax(minimaxDepth, b.Clone(), true, b.curTurn)
	return move
}

func minimax(depth int, state *Board, isMax bool, player int) (*Move, int) {
	state.BeginOfTurn()
	allMoves := state.GetAllMoves()
	if len(allMoves) == 0 {
		// panic(fmt.Errorf("no move available to choose"))
		return nil, state.GetScoreOfPlayer(player)
	}

	curMove := allMoves[0]
	newState := state.Clone()
	newState.ApplyMove(curMove)
	curScore := newState.GetScoreOfPlayer(player)

	if !newState.IsEndGame() && depth > 1 {
		newState.EndOfTurn()
		_, curScore = minimax(depth-1, newState.Clone(), !isMax, player)
	}

	for _, move := range allMoves[1:] {
		newState := state.Clone()
		newState.ApplyMove(move)
		score := newState.GetScoreOfPlayer(player)

		if !newState.IsEndGame() && depth > 1 {
			newState.EndOfTurn()
			_, score = minimax(depth-1, newState.Clone(), !isMax, player)
		}

		if isMax && score > curScore {
			curScore = score
			curMove = move
		} else if !isMax && score < curScore {
			curScore = score
			curMove = move
		}
	}

	return curMove, curScore
}

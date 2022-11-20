package awale

import (
	"fmt"
	"strconv"
)

type Move struct {
	index     int
	direction int
}

func NewMove(index int, direction int) *Move {
	return &Move{index, direction}
}

func (m *Move) ToString() string {
	return fmt.Sprintf("[%v,%v]", m.index, m.direction)
}

type Board struct {
	squares []Square
	curTurn int
	scores  []int
}

func NewBoard() *Board {
	res := make([]Square, 0)

	for j := 0; j < 2; j++ {
		res = append(res, NewBossSquareDefault())

		for i := 0; i < 5; i++ {
			res = append(res, NewNormalSquareDefault())
		}
	}

	return &Board{res, 0, []int{0, 0}}
}

func NewCustomBoard(units []int) *Board {
	res := make([]Square, 0)

	for i := 0; i < 12; i++ {
		if i%6 == 0 {
			res = append(res, NewBossSquareWithCount(units[i]))
		} else {
			res = append(res, NewNormalSquareWithCount(units[i]))
		}
	}

	return &Board{res, 0, []int{0, 0}}
}

func (b *Board) nextTurn() {
	b.curTurn = 1 - b.curTurn
}

func (b *Board) EndOfTurn() {
	if b.IsEndGame() {
		b.ClearBoard()
	}

	b.nextTurn()
}

func (b *Board) GetScore() int {
	return (b.scores[0] - b.scores[1]) * (1 - 2*b.curTurn)
}

func (b *Board) GetScoreOfPlayer(player int) int {
	return (b.scores[0] - b.scores[1]) * (1 - 2*player)
}

func (b *Board) BeginOfTurn() {
	for i := 1 + 6*b.curTurn; i < 6+6*b.curTurn; i++ {
		if !b.squares[i].IsEmpty() {
			return
		}
	}

	for i := 1 + 6*b.curTurn; i < 6+6*b.curTurn && b.scores[b.curTurn] > 0; i++ {
		b.squares[i].AddUnit(normalUnitSample)
		b.scores[b.curTurn]--
	}
}

func (b *Board) GetAllMoves() []*Move {
	res := make([]*Move, 0)
	for i := 1 + 6*b.curTurn; i < 6+6*b.curTurn; i++ {
		if !b.squares[i].IsEmpty() {
			res = append(res, &Move{i, 1})
			res = append(res, &Move{i, -1})
		}
	}

	return res
}

func (b *Board) ExecuteTurnWithMove(move *Move) {
	b.BeginOfTurn()
	b.ApplyMove(move)
	b.EndOfTurn()
}

func (b *Board) ApplyMove(move *Move) {
	i := (move.index + move.direction + len(b.squares)) % len(b.squares)
	j := move.index

	for !b.squares[j].IsEmpty() && b.squares[j].IsMovable() {
		for ; !b.squares[j].IsEmpty(); i = (i + move.direction + len(b.squares)) % len(b.squares) {
			b.squares[i].AddUnit(b.squares[j].PopUnit())
		}

		j = i
		i = (i + move.direction + len(b.squares)) % len(b.squares)
	}

	for b.squares[j].IsEmpty() && b.squares[j].IsPassable() && !b.squares[i].IsEmpty() {
		for !b.squares[i].IsEmpty() {
			b.scores[b.curTurn] += b.squares[i].PopUnit().GetValue()
		}

		i = (i + move.direction + len(b.squares)) % len(b.squares)
		j = i
		i = (i + move.direction + len(b.squares)) % len(b.squares)
	}
}

func (b *Board) Clone() *Board {
	res := make([]Square, 0)

	for _, square := range b.squares {
		res = append(res, square.Clone())
	}

	return &Board{res, b.curTurn, []int{b.scores[0], b.scores[1]}}
}

func (b *Board) IsEndGame() bool {
	if b.squares[0].IsEmpty() && b.squares[6].IsEmpty() {
		return true
	}

	for player := 0; player < 2; player++ {
		isEmpty := true
		for i := 1 + player*6; i < 6+player*6; i++ {
			if !b.squares[i].IsEmpty() {
				isEmpty = false
			}
		}

		if !isEmpty && b.scores[player] == 0 {
			return true
		}
	}

	return false
}

func (b *Board) ClearBoard() {
	for i := 1; i < 6; i++ {
		for !b.squares[i].IsEmpty() {
			b.scores[0] += b.squares[i].PopUnit().GetValue()
		}
	}

	for i := 7; i < 12; i++ {
		for !b.squares[i].IsEmpty() {
			b.scores[1] += b.squares[i].PopUnit().GetValue()
		}
	}
}

func (b *Board) ToString() string {
	res := "_______________\n| |"
	for i := 1; i < 6; i++ {
		res += strconv.Itoa(b.squares[i].GetSize()) + "|"
	}

	res += " |  Score: " + strconv.Itoa(b.scores[0]) + "\n|" + strconv.Itoa(b.squares[0].GetSize()) +
		"|IIIIIIIII|" + strconv.Itoa(b.squares[6].GetSize()) + "|     Player:" + strconv.Itoa(b.curTurn) + "\n| |"
	for i := 11; i >= 7; i-- {
		res += strconv.Itoa(b.squares[i].GetSize()) + "|"
	}

	res += " |  Score: " + strconv.Itoa(b.scores[1]) + "\nTTTTTTTTTTTTTTT\n\n"
	return res
}

// _______________
// | | | | | | | |
// | |IIIIIIIII| |
// | | | | | | | |
// TTTTTTTTTTTTTTT

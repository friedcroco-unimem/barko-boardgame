package awale

import (
	"barkdo/lib/barko"
)

const interval = 0.7

const (
	easy = 0
	hard = 1
)

type gameScene struct {
	*barko.BaseScene
	board        *Board
	units        [][]barko.Sprite
	background   barko.Sprite
	playerTurn   int
	autoOpponent bool
	hardLevel    int
	arrows       []barko.Sprite
	scoreGroups  []*digitGroup
}

func NewGameScene(playerTurn int, autoOpponent bool, hardLevel int) *gameScene {
	units := make([][]barko.Sprite, 0)
	for j := 0; j < 2; j++ {
		units = append(units, NewBossSquareSprites())

		for i := 0; i < 5; i++ {
			units = append(units, NewNormalSquareSprites())
		}
	}

	res := &gameScene{barko.NewBaseScene(), NewBoard(), units, barko.NewSprite("background"),
		playerTurn, autoOpponent, hardLevel, make([]barko.Sprite, 0), []*digitGroup{}}

	return res
}

func NewGameSceneWithSaveData() *gameScene {
	data, _ := isSaveDataExist()
	save := stringToFileData(data)

	board := NewBoard()
	for _, move := range save.Moves {
		board.BeginOfTurn()
		board.ApplyMove(NewMove(move.SquareIndex, move.Direction))
		board.EndOfTurn()
	}

	units := make([][]barko.Sprite, 12)
	for i := 0; i < 12; i++ {
		units[i] = make([]barko.Sprite, 0)
		for _, u := range board.squares[i].GetAllUnits() {
			if u.GetValue() == 1 {
				units[i] = append(units[i], NewNormalUnitSprite())
			} else {
				units[i] = append(units[i], NewBossUnitSprite())
			}
		}
	}

	res := &gameScene{barko.NewBaseScene(), board, units, barko.NewSprite("background"),
		0, true, save.HardLevel, make([]barko.Sprite, 0), []*digitGroup{}}

	return res
}

func NewBossSquareSprites() []barko.Sprite {
	return []barko.Sprite{NewBossUnitSprite()}
}

func NewNormalSquareSprites() []barko.Sprite {
	res := make([]barko.Sprite, 0)
	for i := 0; i < 5; i++ {
		res = append(res, NewNormalUnitSprite())
	}

	return res
}

func (s *gameScene) updateShowScore() {
	if s.board.scores[0] != s.scoreGroups[0].GetValue() {
		s.scoreGroups[0].ChangeValue(s.board.scores[0])
	}

	if s.board.scores[1] != s.scoreGroups[1].GetValue() {
		s.scoreGroups[1].ChangeValue(s.board.scores[1])
	}
}

func (s *gameScene) Init() {
	manager := barko.GetGameManager()
	screenWidth, screenHeight := manager.GetGameLayout()

	// draw background
	s.AddChild(s.background)

	// draw foreground
	foreground := barko.NewSprite("foreground")
	foreground.SetAnchor(0.5, 0.5)
	foreground.SetPosition(float32(screenWidth)/2, float32(screenHeight)/2)
	s.AddChild(foreground)

	// score
	s.scoreGroups = append(s.scoreGroups, NewDigitGroup(s, float32(screenWidth/2-240), float32(screenHeight/2-150)))
	s.scoreGroups = append(s.scoreGroups, NewDigitGroup(s, float32(screenWidth/2+240), float32(screenHeight/2+150)))

	// add child
	squares := len(s.units)
	for i := 0; i < squares; i++ {
		units := len(s.units[i])
		for j := 0; j < units; j++ {
			posX, posY := getUnitAbsolutePosition(i, j)
			s.units[i][j].SetAnchor(0.5, 0.5)
			s.units[i][j].SetPosition(posX, posY)
			s.AddChild(s.units[i][j])
		}
	}

	s.setUpStartTurn()
}

func (s *gameScene) setUpStartTurn() {
	s.updateShowScore()
	steps := s.board.BeginOfTurnToSteps()
	turn := yourTurn
	if s.playerTurn != s.board.curTurn {
		turn = oppoturn
	}

	if len(steps) == 0 {
		s.AddChild(NewTextSprite(turn, s, barko.NewFuncCallWithFunc(func() {
			s.waitForMove()
		})))
	} else {
		s.AddChild(NewTextSprite(turn, s, barko.NewFuncCallWithFunc(func() {
			s.background.AddAction(s.ApplyMoveStepSequence(steps, barko.NewFuncCallWithFunc(func() {
				s.updateShowScore()
				s.waitForMove()
			})))
		})))
	}
}

func (s *gameScene) waitForMove() {
	if s.board.curTurn == s.playerTurn {
		s.waitForPlayerInput()
	} else {
		if s.autoOpponent {
			if s.hardLevel == hard {
				s.setUpApplyMove(s.board.GetBestMove())
			} else {
				s.setUpApplyMove(s.board.GetRandomMove())
			}
		} else {
			s.waitForOpponent()
		}
	}
}

func (s *gameScene) waitForPlayerInput() {
	moves := s.board.GetAllMoves()
	for _, move := range moves {
		arrow := NewArrowSprite(move.Index, move.Direction, s)
		s.AddChild(arrow)
		s.arrows = append(s.arrows, arrow)
	}
}

func (s *gameScene) waitForOpponent() {
	// Over network
	go func() {
		for {
			msg := getNetworkMsg()
			if msg != nil {
				setNetworkMsg(nil)
				s.setUpApplyMove(NewMove(msg.SquareIndex, msg.Direction))
				break
			}
		}
	}()
}

func (s *gameScene) setUpApplyMove(move *Move) {
	steps := s.board.ApplyMoveToSteps(move)
	s.background.AddAction(s.ApplyMoveStepSequence(steps, barko.NewFuncCallWithFunc(func() {
		s.updateShowScore()
		s.setUpEndTurn()
	})))

	if s.autoOpponent {
		addMoveToSaveData(move.Index, move.Direction, s.hardLevel)
	}
}

func (s *gameScene) setUpEndTurn() {
	steps := s.board.EndOfTurnToSteps()
	if len(steps) == 0 {
		if s.board.IsEndGame() {
			s.raiseEndGameSignal()
		} else {
			s.setUpStartTurn()
		}
	} else {
		s.background.AddAction(s.ApplyMoveStepSequence(steps, barko.NewFuncCallWithFunc(func() {
			s.updateShowScore()
			if s.board.IsEndGame() {
				s.raiseEndGameSignal()
			} else {
				s.setUpStartTurn()
			}
		})))
	}
}

func (s *gameScene) raiseEndGameSignal() {
	text := win
	if s.board.scores[s.playerTurn] < s.board.scores[1-s.playerTurn] {
		text = lose
	} else if s.board.scores[s.playerTurn] == s.board.scores[1-s.playerTurn] {
		text = draw
	}

	s.AddChild(NewTextSprite(text, s, barko.NewFuncCall()))
	deleteSaveData()
}

func (s *gameScene) ApplyMoveStepSequence(steps []*MoveStep, finalAction barko.Action) barko.Action {
	res := finalAction
	for i := len(steps) - 1; i >= 0; i-- {
		tempI := i
		tempRes := res
		cur := barko.NewFuncCallWithFunc(func() {
			seq := s.applyMoveStep(steps[tempI])
			seq.AddAction(tempRes)
			s.background.AddAction(seq)
		})
		res = cur
	}

	return res
}

func (s *gameScene) applyMoveStep(step *MoveStep) barko.Sequence {
	manager := barko.GetGameManager()
	width, height := manager.GetGameLayout()
	fromX, fromY := float32(0), float32(0)
	toX, toY := float32(0), float32(0)

	if step.FromSquareIndex == -1 {
		fromX, fromY = float32(width/2), -40
	} else if step.FromSquareIndex == -2 {
		fromX, fromY = float32(width/2), float32(height+40)
	} else {
		fromX, fromY = getUnitAbsolutePosition(step.FromSquareIndex, step.UnitIndex)
	}

	if step.ToSquareIndex == -1 {
		toX, toY = float32(width/2), -40
	} else if step.ToSquareIndex == -2 {
		toX, toY = float32(width/2), float32(height+40)
	} else {
		toX, toY = getUnitAbsolutePosition(step.ToSquareIndex, len(s.units[step.ToSquareIndex]))
	}

	res := barko.NewSequence()
	if step.FromSquareIndex < 0 {
		newUnit := NewNormalUnitSprite()
		newUnit.SetAnchor(0.5, 0.5)
		newUnit.SetPosition(fromX, fromY)
		res.AddAction(barko.NewFuncCallWithFunc(func() {
			s.AddChild(newUnit)
			newUnit.AddAction(barko.NewMoveBy(toX-fromX, toY-fromY, interval))
			s.units[step.ToSquareIndex] = append(s.units[step.ToSquareIndex], newUnit)
		}))

		res.AddAction(barko.NewWait(interval * 0.6))
		return res
	}

	if step.ToSquareIndex < 0 {
		res.AddAction(barko.NewFuncCallWithFunc(func() {
			unit := s.units[step.FromSquareIndex][step.UnitIndex]
			s.units[step.FromSquareIndex] = s.units[step.FromSquareIndex][:len(s.units[step.FromSquareIndex])-1]
			unit.AddAction(barko.NewSequenceWithActions([]barko.Action{
				barko.NewMoveBy(toX-fromX, toY-fromY, interval),
				barko.NewDisappear(),
			}))
		}))

		res.AddAction(barko.NewWait(interval * 0.6))
		return res
	}

	res.AddAction(barko.NewFuncCallWithFunc(func() {
		unit := s.units[step.FromSquareIndex][step.UnitIndex]
		s.units[step.FromSquareIndex] = s.units[step.FromSquareIndex][:len(s.units[step.FromSquareIndex])-1]
		s.units[step.ToSquareIndex] = append(s.units[step.ToSquareIndex], unit)
		unit.AddAction(barko.NewJumpBy(toX-fromX, toY-fromY, 100, interval))
	}))

	res.AddAction(barko.NewWait(interval * 0.6))
	return res
}

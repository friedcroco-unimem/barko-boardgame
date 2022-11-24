package awale

import (
	"encoding/json"
	"io/ioutil"
)

type writtenMove struct {
	SquareIndex int `json:"square_index"`
	Direction   int `json:"direction"`
}

type fileData struct {
	Moves     []writtenMove `json:"moves"`
	HardLevel int           `json:"hard_level"`
}

const savefile = "awale.data"

func stringToFileData(str []byte) fileData {
	var res fileData
	json.Unmarshal(str, &res)
	return res
}

func isSaveDataExist() ([]byte, bool) {
	data, err := ioutil.ReadFile(savefile)
	if err != nil {
		return []byte{}, false
	}

	if string(data) == "" {
		return []byte{}, false
	}

	return data, true
}

func deleteSaveData() {
	ioutil.WriteFile(savefile, []byte(""), 0777)
}

func addMoveToSaveData(squareIndex int, direction int, hardLevel int) {
	data, exists := isSaveDataExist()
	var save fileData = fileData{make([]writtenMove, 0), hardLevel}
	if exists {
		save = stringToFileData(data)
	}

	save.Moves = append(save.Moves, writtenMove{squareIndex, direction})
	res, _ := json.Marshal(save)
	ioutil.WriteFile(savefile, res, 0777)
}

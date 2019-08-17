package habit

import (
	"encoding/json"
	"os"
)

type List struct {
	HabitLists []struct {
		Time  string   `json:"time"`
		Lists []string `json:"lists"`
	} `json:"habitLists"`
}

func (l *List) LoadListFile(file string) error {
	listFile, err := os.Open(file)
	defer listFile.Close()
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(listFile)
	jsonParser.Decode(l)
	return nil
}

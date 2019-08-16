package habit

import (
	"encoding/json"
	"os"
)

type List struct {
	HabitList []struct {
		Time string   `json:"time"`
		List []string `json:"list"`
	} `json:"habitList"`
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

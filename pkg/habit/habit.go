package habit

import (
	"fmt"
	"os"
)

func Do(filename string) {
	list := new(List)
	err := list.LoadListFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(list)
}

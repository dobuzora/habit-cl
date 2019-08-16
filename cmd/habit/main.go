package main

import (
	"fmt"

	"github.com/dobuzora/habit-cl/pkg/habit"
)

func main() {
	habit.Do("list.json")
	fmt.Println("habit")
}

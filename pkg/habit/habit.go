package habit

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Do(filename string) {
	list := new(List)
	err := list.LoadListFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:%q", err)
		os.Exit(1)
	}

	ch := input(os.Stdin)
	for _, l := range list.HabitLists {
		fmt.Fprintf(os.Stdout, "[%s]", l.Time)
		for _, ls := range l.Lists {
			<-ch
			fmt.Fprintf(os.Stdout, "\t%s", ls)
		}
		<-ch
	}
}

func input(r io.Reader) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		sc := bufio.NewScanner(r)
		for sc.Scan() {
			ch <- struct{}{}
		}
		close(ch)
	}()
	return ch
}

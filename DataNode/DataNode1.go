package main

import (
	"bufio"
	"fmt"
	"os"
)

type dataNodeData struct {
	playerNumber int
	playerStage  int
	actions      []int
}

func (p *dataNodeData) saveAction() int {
	filename := fmt.Sprintf("jugador_%d__ronda_%d.txt", p.playerNumber, p.playerStage)

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0600)

	check(err)

	for i := 0; i < len(p.actions); i++ {
		fmt.Fprintf(f, "%d\n", p.actions[i])
	}
	f.Close()
	return 0
}

func (p *dataNodeData) getStageData() {
	filename := fmt.Sprintf("jugador_%d__ronda_%d.txt", p.playerNumber, p.playerStage)

	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)
	var num int
	var str string
	var pl []int
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {

		str = scanner.Text()
		fmt.Sscanf(str, "%d", &num)
		pl = append(pl, num)

	}
	p.actions = pl
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

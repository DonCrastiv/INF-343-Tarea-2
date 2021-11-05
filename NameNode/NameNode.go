package main

import (
	"bufio"
	"fmt"
	"os"
)

type nameNodeData struct {
	playerNumber int
	playerStage  int
	ip           string
}

func (p *nameNodeData) savePlayerData() {
	filename := "nameNodeDataLocation.txt"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0600)

	check(err)
	fmt.Fprintf(f, "Jugador_%d Ronda_%d 10.0.1.10\n", p.playerNumber, p.playerStage)
	f.Close()
}

func (p *nameNodeData) getStoredIP() {
	filename := "nameNodeDataLocation.txt"

	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)
	var str, ip string
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {

		str = scanner.Text()
		var pl, st int
		fmt.Sscanf(str, "Jugador_%d Ronda_%d %s", &pl, &st, &ip)

		if pl == p.playerNumber && st == p.playerStage {
			p.ip = ip
			break
		}
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

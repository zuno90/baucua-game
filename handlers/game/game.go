package game

import (
	"log"
	"time"
)

type Stages string

const (
	READYTIME  Stages = "READY TIME"
	BETTIME    Stages = "BET TIME"
	RESULTTIME Stages = "RESULT TIME"
	readyTime         = 5
	betTime           = 45
	resultTime        = 10
)

type Game struct {
	Stage Stages `json:"stage"`
	No    int    `json:"no"`
	Sec   int    `json:"sec"`
}

func (Server) NewGame(stage Stages, no, sec int) *Game {
	return &Game{
		Stage: stage,
		No:    no,
		Sec:   sec,
	}
}

func HandleGameInfo() (Stages, int, int) {
	h := time.Now().Hour()
	m := time.Now().Minute()
	s := time.Now().Second()
	log.Printf("%d giờ %d phút %d giây", h, m, s)
	currentGameNo := ((h*3600 + m*60 + s) / 60) + 1
	var stage Stages
	var c int
	if s >= 0 && s <= readyTime {
		stage = Stages(READYTIME)
		c = readyTime - s
	} else if s > readyTime && s <= betTime+readyTime {
		stage = Stages(BETTIME)
		c = betTime - (s - readyTime)
	} else {
		stage = Stages(RESULTTIME)
		c = resultTime - (s - betTime - readyTime)
	}
	return stage, currentGameNo, c
}

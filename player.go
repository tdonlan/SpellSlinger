package main

import (
	"fmt"
	"strings"
)

type Player struct{
	name string
	hp int
	totalHp int
	mp int
	totalMp int
	alive bool
	deathClock int
}

func createPlayer(name string)(player *Player){
	return &Player{strings.TrimSpace(name), 1000, 1000, 1000, 1000, true, 0}

}

func getStats(p *Player)(stats string){
	return fmt.Sprintf("%s -  HP: %d/%d MP: %d/%d Alive: %b", p.name,p.hp,p.totalHp,p.mp,p.totalMp,p.alive)
}
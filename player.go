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

func createDummy(name string)(player *Player){
	return &Player{name, 1000, 1000, 1000, 1000, true, 0}
}

func getStats(p *Player)(stats string){
	return fmt.Sprintf("%s -  HP: %d/%d MP: %d/%d Alive: %b", p.name,p.hp,p.totalHp,p.mp,p.totalMp,p.alive)
}

//Convert these to methods?

func spendMana(p *Player, amt int)(canCast bool){
	if p.alive{
		if p.mp >= amt{
			p.mp -= amt
			return true
		}
	}
	return false
}

func getMana(p *Player, amt int){
	if p.alive{
		p.mp += amt
		if p.mp > p.totalMp{
			p.mp = p.totalMp
		}
	}
}

func damage(p *Player, dmg int)(isAlive bool){
	if p.alive {
		p.hp -= dmg
		if p.hp <= 0 {
			p.hp = 0
			p.alive = false
		}
	}
	return p.alive
}

func heal(p *Player, amt int){
	if p.alive{
		p.hp += amt
		if p.hp > p.totalHp{
			p.hp = p.totalHp
		}
	}
}
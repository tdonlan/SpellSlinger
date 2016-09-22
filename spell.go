package main

import (
	"strings"
	"fmt"
)

const Help  = `Welcome to Spell!

Commands:
help: this message
list: list of spells
cast: cast spell
stats: show stats
leaders: leader board

`

const List = `Spell List:

Fire (cast fire <target>)
Heal (cast heal <target>

`

func parseMsg(p *Player, msg string)(out string) {

	strSplit := strings.Split(string(msg), " ")

	cmd := strings.TrimSpace(strSplit[0])

	fmt.Print("parsing: " + cmd)

	switch(cmd){
	case "help":
		out += Help
		break
	case "list":
		out += List
		break
	case "stats":
		out += getStats(p)
	case "cast":
		out += "casting..."

	default:
		out = msg
		break
	}
	return out
}

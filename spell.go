package main

import (
	"strings"
	"fmt"
)


const Help  = `Welcome to Spell!

Commands:
help: this message
list: list of spells
who: display who is in the room
cast: cast spell
stats: show stats
leaders: leader board

`

const List = `Spell List:

Fire (cast fire <target>)
Heal (cast heal <target>)
Dummy (cast dummy <name>)

`


func parseMsg(p *Player, msg string)(out string) {

	strSplit := strings.Split(string(msg), " ")

	cmd := strings.TrimSpace(strSplit[0])

	fmt.Printf("%s: %s", p.name, msg)

	switch(cmd){
	case "help":
		out += Help
		break
	case "list":
		out += List
		break
	case "who":
		out += getWho()
	case "stats":
		out += getStats(p)
	case "cast":
		if len(strSplit) >= 3{
			spell := strings.TrimSpace(strSplit[1])
			target := strings.TrimSpace(strSplit[2])
			switch(spell){
			case "fire":
				out = castFire(p, target)
			case "heal":
				out = castHeal(p, target)

			case "dummy":
				out = castDummy(p,target)
			}
		} else{
			out = fmt.Sprintf("Spell fizzles! %s")
		}
	default:
		out = msg
		break
	}
	return out
}

func getWho()(out string){
	out = "Wizard list: \n"
	for _,v := range playerList {
		out += fmt.Sprintf("%s\n",v.name )
	}
	for k := range dummyList {
		out += fmt.Sprintf("%s\n",k )
	}
	out += "\n"
	return out
}

func castFire(player *Player, target string)(out string){
	out = ""
	//put this somewhere?
	dmg := 250
	cost := 100

	//check if target exists
	if _,ok := dummyList[target]; !ok{
		out = "Invalid Target!\n"
		return out
	}

	//spend mana
	if !spendMana(player, cost){
		out = fmt.Sprintf("Not Enough mana! (need %d)", cost)
	}

	//do damage
	targetDead := damage(dummyList[target],dmg)

	//check for target death
	out += fmt.Sprintf("%s casts Fire at %s for %d!\n",player.name,target, dmg)
	if targetDead{
		out += fmt.Sprintf("%s was killed by %s!",target,player.name)
	}
	return out
}

func castHeal(player *Player, target string)(out string){
	return fmt.Sprintf("%s casts Heal at %s!\n",player.name,target)
}

func castDummy(player *Player, target string)(out string){
	dummy := createDummy(target)
	dummyList[dummy.name] = dummy
	out = fmt.Sprintf("%s summons a dummy named %s!\n",player.name,dummy.name)
	return out
}


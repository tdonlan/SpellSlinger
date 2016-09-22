package main

import (
	"net"
	"bufio"
	"strconv"
	"fmt"
	//"log"
)

const PORT = 3540

var playerList map[net.Conn]*Player

func main() {

	playerList = make(map[net.Conn]*Player)

	server, err := net.Listen("tcp", ":" + strconv.Itoa(PORT))
	if server == nil {
		panic("couldn't start listening: " + err.Error())
	}
	conns := clientConns(server)
	for {
		go handleConn(<-conns)
	}
}

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if client == nil {
				fmt.Printf("couldn't accept: " + err.Error())
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())

			client.Write([]byte("Name?"))

			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if _, ok := playerList[client]; !ok {
			playerList[client] = createPlayer(string(line))
		}

		if err != nil { // EOF, or worse
			break
		}
		//log.Print(string(line))

		client.Write([]byte( parseMsg(playerList[client], string(line))))
	}
}
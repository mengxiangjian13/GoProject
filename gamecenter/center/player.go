package center

import (
	"fmt"
)

type Player struct {
	Name  string
	Level int
	Exp   int

	mq chan *Message
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"", 0, 0, m}

	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "receive message", msg.Content)
		}
	}(player)

	return player
}

package mp

import (
	"fmt"
	"time"
)

type Player interface {
	Play(source string)
}

type MP3Player struct {
	stat     int
	progress int
}

func (p *MP3Player) Play(source string) {
	fmt.Println("Playing mp3 music", source)

	p.progress = 0

	for p.progress < 100 {
		// 假装播放。0.1秒
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}

	fmt.Println("\nfinished playing", source)
}

func Play(source, mtype string) {
	var p Player

	switch mtype {
	case "MP3":
		p = &MP3Player{}
	}
	p.Play(source)
}

package player

import "github.com/solidity247/go_game_v1/world"

type Player struct {
	CurrentLocation *world.Location
	BackPack        *[]string
}

func New() *Player {
	return &Player{}
}

func (p *Player) GoTo(l *world.Location) {
	p.CurrentLocation = l
}

package player

import (
	"slices"
	"strings"

	"github.com/solidity247/go_game_v1/world"
)

type Player struct {
	CurrentLocation world.GamingLocation
	BackPack        *BackPack
	Todos           Todos
}

func New() *Player {
	return &Player{
		BackPack: &BackPack{},
		Todos:    Todos{[]string{"собрать рюкзак", "идти в универ"}},
	}
}

func (p *Player) GoTo(l world.GamingLocation) {
	p.CurrentLocation = l
}

func (p *Player) WearBackPack() {
	p.BackPack.Activate()
}

func (p *Player) GrabItem(item string) {
	p.BackPack.Store(item)
}

type Todos struct {
	todos []string
}

func (t *Todos) Task() string {
	if len(t.todos) == 0 {
		return ""
	}
	return t.todos[0]
}

func (t *Todos) Complete(task string) {
	for i, v := range t.todos {
		if v == task {
			t.todos = slices.Delete(t.todos, i, i+1)
		}
	}
}

func (t *Todos) All() string {
	return "надо " + strings.Join(t.todos, " и ")
}

type BackPack struct {
	Items  []string
	active bool
}

func (b *BackPack) Activate() {
	b.active = true
}

func (b *BackPack) IsActive() bool {
	return b.active
}

func (b *BackPack) Store(item string) {
	b.Items = append(b.Items, item)
}

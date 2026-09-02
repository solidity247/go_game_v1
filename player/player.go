package player

import (
	"slices"
	"strings"

	"github.com/solidity247/go_game_v1/world"
)

type Player struct {
	CurrentLocation *world.Location
	BackPack        *[]string
	Todos           Todos
}

func New() *Player {
	return &Player{
		Todos: Todos{[]string{"собрать рюкзак", "идти в универ"}},
	}
}

func (p *Player) GoTo(l *world.Location) {
	p.CurrentLocation = l
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

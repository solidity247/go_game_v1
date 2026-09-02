package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/solidity247/go_game_v1/player"
	"github.com/solidity247/go_game_v1/world"
)

type Game struct {
	scanner *bufio.Scanner
	player  *player.Player
	world   *world.World
}

func New() *Game {
	scanner := bufio.NewScanner(os.Stdin)
	w := world.New()
	p := player.New()
	p.GoTo(w.GetLocation("кухня"))

	game := Game{
		scanner: scanner,
		player:  p,
		world:   w,
	}
	return &game
}

func (g *Game) Play() {
	for {
		command, err := g.getCommand("")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if command[0] == "exit" {
			break
		}
		fmt.Println(g.RunCommand(command...))
	}
}

func (g *Game) getCommand(message string) ([]string, error) {
	fmt.Print(message)
	g.scanner.Scan()
	g.scanner.Text()
	vals := strings.Fields(g.scanner.Text())

	if len(vals) == 1 || len(vals) == 2 || len(vals) == 3 {
		return vals, nil
	}

	return vals, errors.New("WTF are you typing?")
}

func (g *Game) RunCommand(commands ...string) string {
	switch commands[0] {
	case lookAround: // осмотреться ()
		return g.handleLookAround()
	case goTo: // идти <route>
		return g.handleGoToLocation(commands[1])
	case wear: // надеть <item>
		return g.handleWearItem(commands[1])
	case take: // взять <item>
		return commands[0]
	case apply: // применить <item> -> <object>
		return commands[0]
	default:
		return "неизвестная команда"
	}
}

func (g *Game) RunRawCommand(inp string) string {
	var subCommands []string = strings.Fields(inp)
	return g.RunCommand(subCommands...)
}

func (g *Game) handleLookAround() string {
	l := g.player.CurrentLocation
	manifest := l.Manifest
	items := renderRoomObjects(&l.Items)
	var todos string
	if l.ShowTodo {
		todos = g.player.Todos.All()
	}
	avaliableRoutes := l.GetAvaliableRoutes()

	return fmt.Sprintf("%s. %s", joinValues(", ", manifest, items, todos), avaliableRoutes)
}

func (g *Game) handleGoToLocation(destination string) string {
	if !world.IsValidPath(destination) {
		return ""
	}
	canGo, reason := g.player.CurrentLocation.CanGoTo(destination)
	if !canGo {
		return reason
	}

	g.player.GoTo(g.world.GetLocation(destination))

	return renderArrivalMessage(g.player.CurrentLocation)
}

func renderRoomObjects(items *[]world.RoomObj) string {
	if items == nil {
		return ""
	}
	objs := make([]string, len(*items))
	for i, v := range *items {
		if v.IsEmpty() {
			objs[i] = ""
		} else {
			objs[i] = v.Render()
		}
	}
	return joinValues(", ", objs...)
}

func renderArrivalMessage(l *world.Location) string {
	return fmt.Sprintf("%s. %s", l.WelcomeMessage, l.GetAvaliableRoutes())
}

func joinValues(delimiter string, values ...string) string {
	vals := slices.Collect(func(yield func(string) bool) {
		for _, v := range values {
			if v == "" {
				continue
			}
			if !yield(v) {
				return
			}
		}
	})
	return strings.Join(vals, delimiter)
}

func (g *Game) handleWearItem(item string) string {
	if g.player.CurrentLocation.TakeItem(item) {
		g.player.Todos.Complete("собрать рюкзак")
		return "вы надели: рюкзак"
	}
	return "нет такого"
}

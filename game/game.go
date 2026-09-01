package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/solidity247/go_game_v1/player"
	"github.com/solidity247/go_game_v1/world"
)

type Game struct {
	scanner *bufio.Scanner
	player  *player.Player
	world		*world.World
}

func New() *Game {
	scanner := bufio.NewScanner(os.Stdin)
	w :=	world.New()
	p := player.New()
	p.GoTo(w.GetLocation("кухня"))

	game := Game{
		scanner:	scanner,
		player:		p,
		world:		w,
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

func (g *Game) RunCommand(command ...string) string {
	switch command[0] {
	case lookAround: // осмотреться ()
		return takeAlookAround(g.player.CurrentLocation)
	case goTo: // идти <route>
		return command[0]
	case wear: // надеть <item>
		// need second arg Item
		return command[0]
	case take: // взять <item>
		return command[0]
	case apply: // применить <item> -> <object>
		return command[0]
	default:
		return "неизвестная команда"
	}
}

func (g *Game) RunRawCommand(inp string) string {
	var subCommands []string = strings.Fields(inp)
	return g.RunCommand(subCommands...)
}

func takeAlookAround(l *world.Location) string {
	manifest := l.Manifest
	items := renderItems(&l.Items)
	todos := l.Todos
	avaliableRoutes := l.GetAvaliableRoutes()
		
	return fmt.Sprintf("%s. %s", strings.Join([]string{manifest, items, todos}, ", "), avaliableRoutes)
}

func renderItems (items *[]map[string][]string) string {
	res := ""
	if items == nil {
		return  res
	}
	for _, obj := range *items {
		for name, itms := range obj {
			res += fmt.Sprintf("%s: %s", name, strings.Join(itms, ", "))
		}
	}
	return res
}
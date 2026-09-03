package game

import (
	"fmt"

	"github.com/solidity247/go_game_v1/utils"
	"github.com/solidity247/go_game_v1/world"
)

func (g *Game) handleLookAround() string {
	l := g.player.CurrentLocation
	manifest := g.player.CurrentLocation.GetManigest()
	items := g.player.CurrentLocation.RenderObjects()
	var todos string
	if l.ShowTodo() {
		todos = g.player.Todos.All()
	}
	avaliableRoutes := l.GetAvaliableRoutes()

	return utils.JoinValues(". ", utils.JoinValues(", ", manifest, items, todos), avaliableRoutes)
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

	return g.player.CurrentLocation.RenderArrivalMessage()
}

func (g *Game) handleWearItem(item string) string {
	if g.player.CurrentLocation.TakeItem(item) {
		g.player.Todos.Complete("собрать рюкзак")
		g.player.WearBackPack()
		return "вы надели: рюкзак"
	}
	return "нет такого"
}

func (g *Game) handleTakeItem(item string) string {
	if !g.player.CurrentLocation.HasItem(item) {
		return "нет такого"
	} else {
		if !g.player.BackPack.IsActive() {
			return "некуда класть"
		}
		g.player.CurrentLocation.TakeItem(item)
		g.player.GrabItem(item)
		return fmt.Sprintf("предмет добавлен в инвентарь: %s", item)
	}
}

func (g *Game) handleApplyItem(item, subject string) string {
	if !g.player.BackPack.HasItem(item) {
		return fmt.Sprintf("нет предмета в инвентаре - %s", item)
	}
	return g.player.CurrentLocation.UtilizeItem(item, subject)
}

package world

import (
	"fmt"
	"slices"
	"strings"

	"github.com/solidity247/go_game_v1/utils"
)

type GamingLocation interface {
	GetManigest() string
	ShowTodo() bool
	RenderObjects() string
	RenderArrivalMessage() string
	GetAvaliableRoutes() string
	CanGoTo(path string) (bool, string)
	HasItem(item string) bool
	HasNoItems() bool
	TakeItem(item string) bool
}

type Location struct {
	Items           []RoomObj
	showTodo        bool
	AvaliableRoutes []string
	WelcomeMessage  string
	Manifest        string
	Door            Door
}

type MyRoom struct {
	Location
}

func (l *Location) GetManigest() string {
	return l.Manifest
}

func (r *MyRoom) RenderObjects() string {
	if r.HasNoItems() {
		return "пустая комната"
	}

	return r.Location.RenderObjects()
}

func (l *Location) ShowTodo() bool{
	return l.showTodo
}

func (l *Location) RenderObjects() string {
	items := l.Items
	objs := make([]string, len(items))
	for i, v := range items {
		if v.IsEmpty() {
			objs[i] = ""
		} else {
			objs[i] = v.Render()
		}
	}
	return utils.JoinValues(", ", objs...)
}

func (l *Location) RenderArrivalMessage() string {
	return fmt.Sprintf("%s. %s", l.WelcomeMessage, l.GetAvaliableRoutes())
}

func (l *Location) GetAvaliableRoutes() string {
	prefix := "можно пройти - "
	options := strings.Join(l.AvaliableRoutes, ", ")

	return prefix + options
}

func IsValidPath(path string) bool {
	return slices.Contains(allPaths, path)
}

func (l *Location) CanGoTo(path string) (bool, string) {
	if !slices.Contains(l.AvaliableRoutes, path) {
		return false, fmt.Sprintf("нет пути в %s", path)
	}
	if l.Door.toPath == path && l.Door.isLocked {
		return false, "дверь закрыта"
	}
	return true, ""
}

func (l *Location) HasItem(item string) bool {
	for _, obj := range l.Items {
		if obj.Has(item) {
			return true
		}
	}
	return false
}

func (l *Location) HasNoItems() bool {
	for _, obj := range l.Items {
		if !obj.IsEmpty() {
			return false
		}
	}
	return true
}

func (l *Location) TakeItem(item string) bool {
	for _, obj := range l.Items {
		if obj.Has(item) {
			obj.TakeOutItem(item)
			return true
		}
	}
	return false
}

// Just keep this small object //

type Door struct {
	toPath   string
	isLocked bool
}

func (d *Door) Unlock() {
	d.isLocked = false
}

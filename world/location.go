package world

import (
	"fmt"
	"slices"
	"strings"
)

type Location struct {
	Id              string
	Title           string
	Items           []RoomObj
	Todos           string
	AvaliableRoutes []string
	WelcomeMessage  string
	Manifest        string
	Door            Door
}

type Door struct {
	toPath   string
	isLocked bool
}

func (d *Door) Unlock() {
	d.isLocked = false
}

// func (l *Location) WelcomeMessage() string {
// 	ann := l.WelcomeMessage
// 	paths := l.GetAvaliableRoutes()

// }

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

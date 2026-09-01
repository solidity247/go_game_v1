package world

import (
	"fmt"
	"strings"
)

type Location struct {
	Id              string
	Title           string
	Items           []map[string][]string
	Todos           string
	AvaliableRoutes []string
	Interests       string
	Manifest        string
}

func (l *Location) RenderArrival() {
	ann := l.Interests
	paths := l.GetAvaliableRoutes()

	fmt.Printf("%s. %s", ann, paths)
}

func (l *Location) GetAvaliableRoutes() string {
	prefix := "можно пройти - "
	options := strings.Join(l.AvaliableRoutes, ", ")

	return prefix + options
}

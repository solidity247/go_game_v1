package world

import (
	"fmt"
	"slices"
	"strings"
)

func createKitchen() *Location {
	return &Location{
		// Title:           "кухня",
		Items:           []RoomObj{{"на столе", &[]string{"чай"}}},
		AvaliableRoutes: []string{"коридор"},
		showTodo:        true,
		WelcomeMessage:  "кухня, ничего интересного",
		Manifest:        "ты находишься на кухне",
	}
}

func createCorridor() *Location {
	return &Location{
		AvaliableRoutes: []string{"кухня", "комната", "улица"},
		WelcomeMessage:  "ничего интересного",
		Manifest:        "ничего интересного",
		Door:            Door{true, "улица", true},
	}
}

func createRoom() *MyRoom {
	return &MyRoom{
		Location: Location{
			Items:           []RoomObj{{"на столе", &[]string{"ключи", "конспекты"}}, {"на стуле", &[]string{"рюкзак"}}},
			AvaliableRoutes: []string{"коридор"},
			WelcomeMessage:  "ты в своей комнате",
		},
	}
}

func createStreet() *Location {
	return &Location{
		AvaliableRoutes: []string{"домой"},
		WelcomeMessage:  "на улице весна",
	}
}

func createHouse() *Location {
	return &Location{
		AvaliableRoutes: []string{"улица"},
	}
}

type RoomObj struct {
	name    string
	content *[]string
}

func (ro *RoomObj) Render() string {
	return fmt.Sprintf("%s: %s", ro.name, strings.Join((*ro.content), ", "))
}

func (ro *RoomObj) IsEmpty() bool {
	return len(*ro.content) == 0
}

func (ro *RoomObj) Has(item string) bool {
	return slices.Contains(*ro.content, item)
}

func (ro *RoomObj) TakeOutItem(item string) {
	*ro.content = slices.DeleteFunc(*ro.content, func(s string) bool {
		return s == item
	})
}

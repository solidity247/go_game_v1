package world

func createKitchen() *Location {
	return &Location{
		Title:           "кухня",
		Items:           []map[string][]string{{"на столе": []string{"чай"}}},
		AvaliableRoutes: []string{"коридор"},
		Todos:           "надо собрать рюкзак и идти в универ",
		WelcomeMessage:  "кухня, ничего интересного",
		Manifest:        "ты находишься на кухне",
	}
}

func createCorridor() *Location {
	return &Location{
		Title:           "коридор",
		AvaliableRoutes: []string{"кухня", "комната", "улица"},
		WelcomeMessage:  "ничего интересного",
		Manifest:        "ничего интересного",
		Door:            Door{"улица", true},
	}
}

func createRoom() *Location {
	return &Location{
		Title:           "комната",
		Items:           []map[string][]string{{"на столе": []string{"чай"}}},
		AvaliableRoutes: []string{"коридор"},
		// Todos: "надо собрать рюкзак и идти в универ",
		WelcomeMessage: "ты в своей комнате",
	}
}

func createStreet() *Location {
	return &Location{
		Title:           "улица",
		AvaliableRoutes: []string{"домой"},
		// Todos: "надо собрать рюкзак и идти в универ",
		WelcomeMessage: "на улице весна",
	}
}

func createHouse() *Location {
	return &Location{
		// Title: "домой",
		// Items: []map[string][]string{},
		AvaliableRoutes: []string{"улица"},
		// Todos: "надо собрать рюкзак и идти в универ",
		// WelcomeMessage: "кухня, ничего интересного",
	}
}

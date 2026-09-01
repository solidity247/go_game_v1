package world

func createKitchen() *Location {
	return &Location{
		Id:								"kitchen",
		Title:						"кухня",
		Items:						[]map[string][]string{{"на столе": []string{"чай"}}},
		AvaliableRoutes:	[]string{"коридор"},
		Todos:						"надо собрать рюкзак и идти в универ",
		Interests:				"кухня, ничего интересного",
		Manifest: 				"ты находишься на кухне",
	}
}

func createCorridor() *Location {
	return &Location{
		Id:    "corridor",
		Title: "коридор",
		// Items: []map[string][]string{},
		AvaliableRoutes: []string{"кухня", "комната", "улица"},
		Todos:           "надо собрать рюкзак и идти в универ",
		Interests:       "кухня, ничего интересного",
	}
}

func createRoom() *Location {
	return &Location{
		// Id: "room",
		// Title: "коридор",
		// Items: []map[string][]string{},
		// AvaliableRoutes:					[]string{"кухня", "комната", "улица"},
		// Todos: "надо собрать рюкзак и идти в универ",
		// Interests: "кухня, ничего интересного",
	}
}

func createStreet() *Location {
	return &Location{
		// Id: "room",
		// Title: "коридор",
		// Items: []map[string][]string{},
		// AvaliableRoutes:					[]string{"кухня", "комната", "улица"},
		// Todos: "надо собрать рюкзак и идти в универ",
		// Interests: "кухня, ничего интересного",
	}
}

func createHouse() *Location {
	return &Location{
		// Id: "room",
		// Title: "коридор",
		// Items: []map[string][]string{},
		// AvaliableRoutes:					[]string{"кухня", "комната", "улица"},
		// Todos: "надо собрать рюкзак и идти в универ",
		// Interests: "кухня, ничего интересного",
	}
}

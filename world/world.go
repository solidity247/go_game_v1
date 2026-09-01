package world

type World struct {
	maps map[string]*Location
}

func New() *World {
	return &World{
		make(map[string]*Location),
	}
}

func (w *World) GetLocation(locName string) *Location {
	val, ok := w.maps[locName]
	if ok {
		return val
	}
	val = loadLocation(locName)
	w.maps[locName] = val
	return val
}

func loadLocation(l string) *Location {
	switch l {
	case "кухня":
		return createKitchen()
	case "коридор":
		return createCorridor()
	case "комната":
		return createRoom()
	case "улица":
		return createStreet()
	case "домой":
		return createHouse()
	default:
		return &Location{}
	}
}

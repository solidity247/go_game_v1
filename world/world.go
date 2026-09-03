package world

type World struct {
	locations map[string]GamingLocation
}

func New() *World {
	return &World{
		make(map[string]GamingLocation),
	}
}

func (w *World) GetLocation(locName string) GamingLocation {
	val, ok := w.locations[locName]
	if ok {
		return val
	}
	val = loadLocation(locName)
	w.locations[locName] = val
	return val
}

func loadLocation(l string) GamingLocation {
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

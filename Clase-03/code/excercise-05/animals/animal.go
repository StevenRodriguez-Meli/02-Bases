package animals

const (
	dogFood       = 10.0 // kg
	catFood       = 5.0  // kg
	hamsterFood   = 0.25 // kg
	tarantulaFood = 0.15 // kg
)

func Animal(animal string) (func(float64) float64, string) {
	switch animal {
	case "dog":
		return Dog, ""
	case "cat":
		return Cat, ""
	case "hamster":
		return Hamster, ""
	case "tarantula":
		return Tarantula, ""
	default:
		return nil, "Animal not found"
	}
}

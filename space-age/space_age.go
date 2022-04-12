package space

type Planet string

// Age returns your age on a planet.
//
// Benchmark  35714338  ~31 ns/op
func Age(seconds float64, planet Planet) float64 {
	t := float64(31556700)

	switch planet {
	case "Mercury":
		t *= 0.2408467
	case "Venus":
		t *= 0.61519726
	case "Mars":
		t *= 1.8808158
	case "Jupiter":
		t *= 11.862615
	case "Saturn":
		t *= 29.447498
	case "Uranus":
		t *= 84.016846
	case "Neptune":
		t *= 164.79132
	}

	return seconds / t
}

package twofer

// Iteration 1 with fmt.Sprintf:
//
// 1967450 ~629 ns/op 6 allocs/op
//
// Iteration 2 without fmt.Sprintf:
//
// 8778876 ~138 ns/op 0 allocs/op
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}

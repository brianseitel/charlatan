package charlatan

import "math/rand"

func randomElement(list []string) string {
	idx := rand.Intn(len(list))

	return list[idx]
}

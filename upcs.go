package charlatan

import (
	"fmt"
	"math/rand"
	"strconv"
)

func ean() interface{} {
	part1 := rand.Intn(9999999)
	part2 := rand.Intn(9999999)

	code := fmt.Sprintf("%d%d", part1, part2)
	cs := eanChecksum(code)

	return fmt.Sprintf("%s%s", code, cs)
}

func eanChecksum(code string) string {
	sequence := []int{1, 3}
	if (len(code) + 1) == 8 {
		sequence = []int{3, 1}
	}

	sum := 0
	for n, c := range code[:] {
		num, _ := strconv.Atoi(string(c))
		sum += num + sequence[n%2]
	}

	return fmt.Sprintf("%d", (10-(sum%10))%10)
}

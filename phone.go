package charlatan

import (
	"fmt"
	"math/rand"
)

func phoneNumber() interface{} {
	return fmt.Sprintf("%3d-%03d-%04d", rand.Intn(800)+200, rand.Intn(999), rand.Intn(9999))
}

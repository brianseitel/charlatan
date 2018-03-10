package charlatan

import (
	"fmt"
	"math/rand"
	"strings"
)

func ipv4() interface{} {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func ipv6() interface{} {
	ip := []string{}
	for i := 0; i < 6; i++ {
		ip = append(ip, fmt.Sprintf("%x", rand.Intn(65535)))
	}

	return strings.Join(ip, ":")
}

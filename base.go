package charlatan

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	goTime "time"

	netURL "net/url"
)

func digit() interface{} {
	return int64(rand.Intn(9))
}

func number() interface{} {
	return int64(rand.Int())
}

func boolean() interface{} {
	return goTime.Now().UnixNano()%2 == 0
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) interface{} {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func float() interface{} {
	return rand.Float64() * float64(rand.Int())
}

func price() interface{} {
	f := float64(rand.Intn(100)) + rand.Float64()
	s := fmt.Sprintf("%.2f", f)
	o, _ := strconv.ParseFloat(s, 64)
	return o
}

func age() interface{} {
	return int64(rand.Intn(90) + 1)
}

func id() interface{} {
	return int64(rand.Int())
}

func letter() interface{} {
	charset := "abcdefghijklmnopqrstuvwxyz"

	return string(charset[rand.Intn(len(charset))])
}

func url() interface{} {
	u, _ := netURL.Parse(fmt.Sprintf("https://www.example.com/%s", word()))

	return u.String()
}

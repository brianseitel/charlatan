package charlatan

import (
	"math/rand"
	t "time"
)

var daysOfWeek = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
var months = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

func unixTimestamp() interface{} {
	return t.Now().Unix()
}

func datetime() interface{} {
	return t.Now().Format(t.RFC3339)
}

func date() interface{} {
	return t.Now().Format("2006-01-02")
}

func time() interface{} {
	return t.Now().Format("15:04:05")
}

func dayOfWeek() interface{} {
	return randomElement(daysOfWeek)
}

func dayOfMonth() interface{} {
	return int64(rand.Intn(31))
}

func month() interface{} {
	return int64(rand.Intn(12))
}

func monthName() interface{} {
	return randomElement(months)
}

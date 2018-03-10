package main

import (
	"github.com/brianseitel/charlatan"
	"github.com/sanity-io/litter"
)

// Everything ...
type Everything struct {
	Address    string  `charlatan:"address"`
	City       string  `charlatan:"city"`
	Latitude   float64 `charlatan:"latitude"`
	Longitude  float64 `charlatan:"longitude"`
	StreetName string  `charlatan:"streetName"`
	ZIP        string  `charlatan:"zip"`
	State      string  `charlatan:"state"`
	StateAbbr  string  `charlatan:"stateAbbr"`

	Digit   int64   `charlatan:"digit"`
	Number  int64   `charlatan:"number"`
	Boolean bool    `charlatan:"boolean"`
	Float   float64 `charlatan:"float"`
	Price   float64 `charlatan:"price"`
	Age     int64   `charlatan:"age"`
	ID      int64   `charlatan:"id"`
	Letter  string  `charlatan:"letter"`
	URL     string  `charlatan:"url"`

	Sha1   string `charlatan:"sha1"`
	Sha256 string `charlatan:"sha256"`
	MD5    string `charlatan:"md5"`

	UnixTimestamp int64  `charlatan:"unixTimestamp"`
	DateTime      string `charlatan:"dateTime"`
	Date          string `charlatan:"date"`
	Time          string `charlatan:"time"`
	DayOfWeek     string `charlatan:"dayOfWeek"`
	DayOfMonth    int64  `charlatan:"dayOfMonth"`
	Month         int64  `charlatan:"month"`
	MonthName     string `charlatan:"monthName"`

	Email string `charlatan:"email"`
	IPV4  string `charlatan:"ipv4"`
	IPV6  string `charlatan:"ipv6"`

	FirstName string `charlatan:"firstName"`
	LastName  string `charlatan:"lastName"`
	FullName  string `charlatan:"fullName"`

	Phone string `charlatan:"phone"`
	Ean   string `charlatan:"ean"`
	UUID  string `charlatan:"uuid"`

	Word  string   `charlatan:"word"`
	Words []string `charlatan:"words"`
}

func main() {
	c := charlatan.New()

	foo, err := c.Generate(&Everything{})
	if err != nil {
		panic(err)
	}

	litter.Dump(foo)
}

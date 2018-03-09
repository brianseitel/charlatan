package main

import (
	"fmt"
	"math/rand"

	"github.com/brianseitel/charlatan"
)

// User ...
type User struct {
	ID            string `charlatan:"uuid"`
	Name          string `charlatan:"name"`
	FavoriteColor string `charlatan:"hexcolor"`
	UpdatedAt     string `charlatan:"datetime"`
}

func main() {
	c := charlatan.New()
	c.AddCustomTag("hexcolor", generateColor)

	foo, err := c.Generate(&User{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%# v\n", foo)
}

func generateColor() interface{} {
	charset := "1234567890ABCDEF"

	var letters []byte
	for i := 0; i < 6; i++ {
		idx := rand.Intn(len(charset))
		letters = append(letters, charset[idx])
	}

	return fmt.Sprintf("#%s", string(letters))
}

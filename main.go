package main

import (
	"github.com/brianseitel/faker/faker"
	"github.com/sanity-io/litter"
)

// User
type User struct {
	ID    string  `faker:"uuid"`
	Name  string  `faker:"name"`
	Email string  `faker:"email"`
	Age   int     `faker:"age"`
	Cost  float64 `faker:"price"`
}

// Product
type Product struct {
	Tags         []string `faker:"words"`
	Bogo         bool     `faker:"boolean"`
	BulletPoints []string `faker:"words"`
	Categories   struct {
		ID            int    `faker:"id"`
		Name          string `faker:"word"`
		TermsRequired string `faker:"boolean"`
		URL           string `faker:"url"`
	} `faker:"struct"`
	CategoryRank           int      `faker:"number"`
	CustomLabel            bool     `faker:"boolean"`
	DescriptionLabel       []string `faker:"words"`
	HasCustomLabel         bool     `faker:"boolean"`
	HierarchicalCategories struct {
		Level1 []string `faker:"words"`
		Level2 []string `faker:"words"`
		Level3 []string `faker:"words"`
		Level4 []string `faker:"words"`
	} `faker:"struct"`
	Image struct {
		Original string `faker:"url"`
		URL      string `faker:"url"`
	} `faker:"struct"`
	MetroID            int     `faker:"number"`
	Name               string  `faker:"word"`
	NeighborhoodVolume float64 `faker:"float64"`
	ObjectID           string  `faker:"word"`
	OnSale             bool    `faker:"boolean"`
	Price              float64 `faker:"price"`
	ProductID          int     `faker:"id"`
	ProductType        string  `faker:"word"`
	SaleRank           int     `faker:"number"`
	UpdatedAt          string  `faker:"datetime"`
}

func main() {
	foo := faker.New().Generate(&Product{})

	litter.Dump(foo)
}

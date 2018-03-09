package faker

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	netURL "net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/tjarratt/babble"

	uuidv4 "github.com/satori/go.uuid"
)

const tagName = "faker"

var babbler babble.Babbler

func init() {
	babbler = babble.NewBabbler()
	babbler.Separator = " "
}

// Faker interface ...
type Faker interface {
	Generate(thing interface{}) interface{}
}

// SimpleFaker ...
type SimpleFaker struct {
}

// New ...
func New() Faker {
	return &SimpleFaker{}
}

// Generate ...
func (f SimpleFaker) Generate(thing interface{}) interface{} {
	rand.Seed(time.Now().UnixNano())

	v := reflect.ValueOf(thing).Elem()

	vType := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := vType.Field(i)

		switch f.Tag.Get(tagName) {
		case "email":
			v.FieldByName(f.Name).SetString(email())
		case "name":
			v.FieldByName(f.Name).SetString(name())
		case "id":
			v.FieldByName(f.Name).SetInt(id())
		case "number":
			v.FieldByName(f.Name).SetInt(number())
		case "age":
			v.FieldByName(f.Name).SetInt(age())
		case "uuid":
			v.FieldByName(f.Name).SetString(uuid())
		case "float64":
			v.FieldByName(f.Name).SetFloat(float())
		case "price":
			v.FieldByName(f.Name).SetFloat(price())
		case "word":
			v.FieldByName(f.Name).SetString(word())
		case "words":
			v.FieldByName(f.Name).Set(reflect.ValueOf(words()))
		case "url":
			v.FieldByName(f.Name).SetString(url())
		case "datetime":
			v.FieldByName(f.Name).SetString(datetime())
		case "date":
			v.FieldByName(f.Name).SetString(date())
		case "struct":
			s := v.FieldByName(f.Name)
			st := s.Type()
			newitem := reflect.New(st)
			t := New().Generate((&newitem).Interface())
			v.FieldByName(f.Name).Set(reflect.ValueOf(t).Elem())
		}
	}

	return thing
}

func number() int64 {
	return int64(rand.Intn(100))
}

func boolean() bool {
	return time.Now().UnixNano()%2 == 0
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func float() float64 {
	return rand.Float64() * 10
}

func price() float64 {
	f := float64(rand.Intn(100)) + rand.Float64()
	s := fmt.Sprintf("%.2f", f)
	o, _ := strconv.ParseFloat(s, 64)
	return o
}

func age() int64 {
	return int64(rand.Intn(90) + 1)
}

func uuid() string {
	id, _ := uuidv4.NewV4()
	return id.String()
}

func name() string {
	idx := rand.Intn(len(names))

	name := names[idx]
	return fmt.Sprintf("%s%s", bytes.ToUpper([]byte{name[0]}), strings.ToLower(name[1:]))
}

func id() int64 {
	return int64(rand.Intn(99999999))
}

func email() string {
	emails := []string{
		"serendipity@theatre.org",
		"sardelle@iskandia.com",
		"jaxi@snarkyswords.net",
		"rysha@nerdychicks.com",
		"gal@wonderwoman.com",
		"shuri@leiawannabe.com",
	}

	idx := rand.Intn(len(emails))

	return emails[idx]
}

func word() string {
	return babbler.Babble()
}

func words() []string {
	n := rand.Intn(10)

	words := []string{}
	for i := 0; i < n; i++ {
		words = append(words, word())
	}

	return words
}

func url() string {
	babbler.Separator = "-"
	u, _ := netURL.Parse(fmt.Sprintf("https://www.example.com/%s", word()))
	babbler.Separator = " "

	return u.String()
}

func datetime() string {
	return time.Now().Format(time.RFC3339)
}

func date() string {
	return time.Now().Format("2006-01-02")
}

package charlatan

import (
	"fmt"
	"math"
	"math/rand"
	netURL "net/url"
	"reflect"
	"strconv"
	"time"
)

const tagName = "charlatan"

// Charlatan interface ...
type Charlatan interface {
	Generate(thing interface{}) (interface{}, error)
	AddCustomTag(name string, callback func() interface{})
}

// SimpleCharlatan ...
type SimpleCharlatan struct {
	CustomFuncs map[string]func() interface{}
}

// New ...
func New() Charlatan {
	return &SimpleCharlatan{
		CustomFuncs: make(map[string]func() interface{}),
	}
}

// Generate random values for the `thing` struct. If there are any errors, it
// collects them all and returns them in one shot.
func (c SimpleCharlatan) Generate(thing interface{}) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())

	v := reflect.ValueOf(thing).Elem()

	errs := charlatanErrors{}

	vType := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := vType.Field(i)

		tag := f.Tag.Get(tagName)
		switch tag {
		case "email":
			if errs.CheckType(f.Type, "string") {
				v.FieldByName(f.Name).SetString(email())
			}
		case "name":
			if errs.CheckType(f.Type, "string") {
				v.FieldByName(f.Name).SetString(name())
			}
		case "id", "int":
			if errs.CheckType(f.Type, "int") {
				v.FieldByName(f.Name).SetInt(id())
			}
		case "number":
			if errs.CheckType(f.Type, "int") {
				v.FieldByName(f.Name).SetInt(number())
			}
		case "age":
			if errs.CheckType(f.Type, "int") {
				v.FieldByName(f.Name).SetInt(age())
			}
		case "uuid":
			if errs.CheckType(f.Type, "string") {
				v.FieldByName(f.Name).SetString(uuid())
			}
		case "float64":
			if errs.CheckType(f.Type, "float64") {
				v.FieldByName(f.Name).SetFloat(float())
			}
		case "price":
			if errs.CheckType(f.Type, "float64") {
				v.FieldByName(f.Name).SetFloat(price())
			}
		case "word":
			if errs.CheckType(f.Type, "string") {
				v.FieldByName(f.Name).SetString(word())
			}
		case "words":
			if errs.CheckType(f.Type, "[]string") {
				v.FieldByName(f.Name).Set(reflect.ValueOf(words()))
			}
		case "url":
			if errs.CheckType(f.Type, "string") {
				v.FieldByName(f.Name).SetString(url())
			}
		case "datetime":
			if errs.CheckType(f.Type, "string") {
				v.FieldByName(f.Name).SetString(datetime())
			}
		case "date":
			if errs.CheckType(f.Type, "string") {
				v.FieldByName(f.Name).SetString(date())
			}
		case "struct":
			s := v.FieldByName(f.Name)
			st := s.Type()
			newitem := reflect.New(st)
			t, _ := New().Generate((&newitem).Interface())
			v.FieldByName(f.Name).Set(reflect.ValueOf(t).Elem())
		case "boolean":
			if errs.CheckType(f.Type, "bool") {
				v.FieldByName(f.Name).SetBool(boolean())
			}
		default:
			if cb, ok := c.CustomFuncs[tag]; ok {
				v.FieldByName(f.Name).Set(reflect.ValueOf(cb()))
			}
		}
	}

	return thing, nil
}

// AddCustomTag ...
func (c *SimpleCharlatan) AddCustomTag(name string, callback func() interface{}) {
	c.CustomFuncs[name] = callback
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

func id() int64 {
	return int64(rand.Intn(99999999))
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
	u, _ := netURL.Parse(fmt.Sprintf("https://www.example.com/%s", word()))

	return u.String()
}

func datetime() string {
	return time.Now().Format(time.RFC3339)
}

func date() string {
	return time.Now().Format("2006-01-02")
}

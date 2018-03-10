package charlatan

import (
	"fmt"
	"math/rand"
	"reflect"
	goTime "time"
)

const tagName = "charlatan"

var tags map[string]tag

type tag struct {
	Name         string
	Callback     func() interface{}
	ExpectedType string
}

func NewTag(name string, callback func() interface{}, expectedType string) tag {
	return tag{Name: name, Callback: callback, ExpectedType: expectedType}
}

func init() {
	tags = make(map[string]tag)

	// Address
	tags["address"] = NewTag("address", address, "string")
	tags["city"] = NewTag("city", city, "string")
	tags["latitude"] = NewTag("latitude", latitude, "float64")
	tags["longitude"] = NewTag("longitude", longitude, "float64")
	tags["streetName"] = NewTag("streetName", streetName, "string")
	tags["zip"] = NewTag("zip", zip, "string")
	tags["state"] = NewTag("state", state, "string")
	tags["stateAbbr"] = NewTag("stateAbbr", stateAbbr, "string")

	// Base
	tags["digit"] = NewTag("digit", digit, "int64")
	tags["number"] = NewTag("number", number, "int64")
	tags["boolean"] = NewTag("boolean", boolean, "bool")
	tags["float"] = NewTag("float", float, "float64")
	tags["price"] = NewTag("price", price, "float64")
	tags["age"] = NewTag("age", age, "int64")
	tags["id"] = NewTag("id", id, "int64")
	tags["letter"] = NewTag("letter", letter, "string")
	tags["url"] = NewTag("url", url, "string")

	// Crypto
	tags["sha1"] = NewTag("sha1", sha1, "string")
	tags["sha256"] = NewTag("sha256", sha256, "string")
	tags["md5"] = NewTag("md5", md5, "string")

	// DateTime
	tags["unixTimestamp"] = NewTag("unixTimestamp", unixTimestamp, "int64")
	tags["datetime"] = NewTag("datetime", datetime, "string")
	tags["date"] = NewTag("date", date, "string")
	tags["time"] = NewTag("time", time, "string")
	tags["dayOfWeek"] = NewTag("dayOfWeek", dayOfWeek, "string")
	tags["dayOfMonth"] = NewTag("dayOfMonth", dayOfMonth, "int64")
	tags["month"] = NewTag("month", month, "int64")
	tags["monthName"] = NewTag("monthName", monthName, "string")

	// Emails
	tags["email"] = NewTag("email", email, "string")

	// Internet
	tags["ipv4"] = NewTag("ipv4", ipv4, "string")
	tags["ipv6"] = NewTag("ipv6", ipv6, "string")

	// Person
	tags["firstName"] = NewTag("firstName", firstName, "string")
	tags["lastName"] = NewTag("lastName", lastName, "string")
	tags["fullName"] = NewTag("fullName", fullName, "string")

	// Phone
	tags["phone"] = NewTag("phone", phoneNumber, "string")

	// UPCs
	tags["ean"] = NewTag("ean", ean, "string")

	// UUID
	tags["uuid"] = NewTag("uuid", uuid, "string")

	// Words
	tags["word"] = NewTag("word", word, "string")
	tags["words"] = NewTag("words", words, "[]string")
}

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
	rand.Seed(goTime.Now().UnixNano())

	v := reflect.ValueOf(thing).Elem()

	errs := charlatanErrors{}

	vType := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := vType.Field(i)

		tag := f.Tag.Get(tagName)

		switch tag {
		case "struct":
			s := v.FieldByName(f.Name)
			st := s.Type()
			newitem := reflect.New(st)
			t, _ := New().Generate((&newitem).Interface())
			v.FieldByName(f.Name).Set(reflect.ValueOf(t).Elem())
		default:
			fmt.Println(tag)
			if t, ok := tags[tag]; ok {
				if errs.CheckType(tag, f.Type, t.ExpectedType) {
					val := t.Callback()
					switch t.ExpectedType {
					case "string":
						v.FieldByName(f.Name).SetString(val.(string))
					case "int":
						v.FieldByName(f.Name).SetInt(val.(int64))
					case "bool":
						v.FieldByName(f.Name).SetBool(val.(bool))
					default:
						v.FieldByName(f.Name).Set(reflect.ValueOf(val))
					}
				}
			} else if cb, ok := c.CustomFuncs[tag]; ok {
				v.FieldByName(f.Name).Set(reflect.ValueOf(cb()))
			}
		}
	}

	if errs.Errors() != nil {
		panic(errs.Errors())
	}

	return thing, nil
}

// AddCustomTag ...
func (c *SimpleCharlatan) AddCustomTag(name string, callback func() interface{}) {
	c.CustomFuncs[name] = callback
}

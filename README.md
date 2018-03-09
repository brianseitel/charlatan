# Charlatan

Charlatan is a fake data generator for Go that leans on struct tags to help build dummy data.

## Getting Started

Getting started is simple:

```
go get github.com/brianseitel/charlatan
```

To integrate into your app, define a struct and use the appropriate struct tags. Here's quick example:

```golang
type User struct {
    ID int `charlatan:"id"`
    Name string `charlatan:"name"`
    Email string `charlatan:"email"`
}

func main() {
	u := charlatan.New().Generate(&User{})
	fmt.Printf("%# v\n", u)
}
```

## Fake Data Types

| Type | Tag | Description | Example |
|---|---|---|---|
| Email | `email` | A fake email address | `chatty@gmail.com` |
| Name | `name` | Full name of a person | `Serendipity Jones` |
| ID | `id` | A number between 0 and 1,000,000 | `6697` |
| UUID | `uuid` | A standards-compliant UUID | `47036b2f-26af-42cb-92f9-2c6124f90f1347036b2f-26af-42cb-92f9-2c6124f90f13` |
| Number | `number` | A simple number | `25` |
| Age | `age` | A number between 0 and 90 | `37` |
| Float | `float64` | A 64-bit floating point number | `3.36642` |
| Price | `price` | A dollar amount with two decimals | `2.99` |
| Word | `word` | A random word | `chicken` |
| Words | `words` | A random number of random words | `["chicken", "apple", "tesla"]` |
| URL | `url` | A random URL | `http://www.example.com/foo-bar` |
| DateTime | `datetime` | An RFC 3339 compliant date time string | `2018-03-08T01:02:05Z` |
| Date | `date` | A date in the format YYYY-MM-DD | `2018-03-09` |
| Struct | `struct` | Allows you to nest structs | ~ |

## Custom Tags

If the built-in rules above aren't good enough for you, you can add your own! It's as easy as 1, 2, 3:

### One: Define a Function

The signature for the function is `func() interface{}`:

```golang
func generateColor() interface{} {
	charset := "1234567890ABCDEF"

	var letters []byte
	for i := 0; i < 6; i++ {
		idx := rand.Intn(len(charset))
		letters = append(letters, charset[idx])
	}

	return fmt.Sprintf("#%s", string(letters))
}
```

### Two: Add it to your Charlatan instance

The signature for adding a custom tag is `AddCustomTag(name string, callback func() interface{})`

```golang
c := charlatan.New()
c.AddCustomTag("hexcolor", generateColor)
```

### Three: Use it in your struct!

Now you can add it as a struct tag in your struct:

```golang
type App struct {
	BackgroundColor string `charlatan:"hexcolor"`
	ForegroundColor string `charlatan:"hexcolor"`
}
```

## Examples

You can run any of the examples by running the following:

```
go run --tags=examples /examples/{example}/main.go
```

## Contributing

Fork it and submit a PR!

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/brianseitel/charlatan/tags). 

## Authors

* **Brian Seitel** - *Initial work* - [Brian Seitel](https://github.com/PurpleBooth)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
